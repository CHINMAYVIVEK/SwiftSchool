package helper

import (
	"context"
	"fmt"
	"maps"
	"strings"
	"sync"

	"swiftschool/config"
	"swiftschool/domain"

	"github.com/fatih/color"
	"github.com/hashicorp/vault/api"
)

func VaultDiagnostic(vaultAddr, vaultToken string) error {
	logger := GetLogger()

	client, err := api.NewClient(&api.Config{Address: vaultAddr})
	if err != nil {
		logger.Errorf("vault client creation failed: %v", err)
		return err
	}

	client.SetToken(vaultToken)
	color.Cyan("✅ Connected to Vault at %s\n", vaultAddr)

	// --- List all secret engines ---
	mounts, err := client.Sys().ListMounts()
	if err != nil {
		logger.Errorf("failed to list secret engines: %v", err)
		return err
	}

	kvMounts := make(map[string]string)
	fmt.Println("\n📁 Secret Engines:")

	for path, mount := range mounts {
		ver := mount.Options[domain.KvVersionKey]
		if ver == "" {
			ver = "1"
		}

		color.Green(" - %-20s : %-10s (version: %v)", path, mount.Type, ver)

		if mount.Type == domain.KvType {
			kvMounts[path] = ver
		}
	}

	// --- Traverse KV Engines ---
	fmt.Println("\n🔍 Traversing inside KV secret engines:")
	var wg sync.WaitGroup

	for path, ver := range kvMounts {
		wg.Add(1)
		go func(p, v string) {
			defer wg.Done()

			color.Cyan("\nEngine: %s (version: %s)", p, v)

			if err := traverseSecrets(client, p, v, ""); err != nil {
				logger.Errorf("failed traversing %s: %v", p, err)
				color.Red("   ❌ Error traversing %s: %v", p, err)
			}
		}(path, ver)
	}

	wg.Wait()
	color.Cyan("\nVault diagnostic completed successfully ✅\n")

	return nil
}

func traverseSecrets(client *api.Client, mountPath, version, prefix string) error {
	logger := GetLogger()

	listPath := mountPath
	if version == "2" {
		listPath += domain.KvV2Metadata + prefix
	} else {
		listPath += prefix
	}

	list, err := client.Logical().List(listPath)
	if err != nil {
		logger.Errorf("failed to list %s: %v", listPath, err)
		return err
	}

	if list == nil || list.Data == nil {
		color.Yellow("   ⚠️ No secrets at %s%s", mountPath, prefix)
		return nil
	}

	// type assertion using any
	keys, ok := list.Data["keys"].([]any)
	if !ok {
		color.Yellow("   ⚠️ Unexpected keys type at %s%s", mountPath, prefix)
		return nil
	}

	for _, key := range keys {
		k := key.(string)

		if strings.HasSuffix(k, "/") {
			// recursive folder traversal
			_ = traverseSecrets(client, mountPath, version, prefix+k)
			continue
		}

		fullPath := prefix + k
		color.White("   - %s", fullPath)

		if k == domain.InstituteSecretKey {
			secretPath := mountPath + fullPath
			if version == "2" {
				secretPath = mountPath + domain.KvV2Data + fullPath
			}

			secret, err := client.Logical().Read(secretPath)
			if err != nil {
				logger.Errorf("failed to read secret %s: %v", secretPath, err)
				color.Red("     ❌ Failed to read %s", fullPath)
				continue
			}

			data := extractSecretData(secret.Data, version)
			if data == nil {
				color.Red("     ❌ Invalid secret format for %s", fullPath)
				continue
			}

			color.Green("     🔑 Contents of %s:", fullPath)
			for key, val := range data {
				color.White("       %s: %v", key, val)
			}
		}
	}

	return nil
}

func extractSecretData(secretData map[string]any, version string) map[string]any {
	if version == "2" {
		if d, ok := secretData["data"].(map[string]any); ok {
			dst := map[string]any{}
			maps.Copy(dst, d)
			return dst
		}
	}

	dst := map[string]any{}
	maps.Copy(dst, secretData)
	return dst
}

// -------------------------
// Demo Function
// -------------------------
func EncryptAndDecryptDemo(ctx context.Context, cfg config.AppConfig, instituteID, plaintext string) (string, string, error) {

	cipherText, err := EncryptText(ctx, cfg, instituteID, plaintext)
	if err != nil {
		return "", "", err
	}

	decrypted, err := DecryptText(ctx, cfg, instituteID, cipherText)
	if err != nil {
		return "", "", err
	}

	return cipherText, decrypted, nil
}
