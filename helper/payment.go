package helper

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// ------------------------ Gateway Types ------------------------
type PaymentGatewayType string

const (
	Razorpay  PaymentGatewayType = "razorpay"
	PayUMoney PaymentGatewayType = "payumoney"
	CCAvenue  PaymentGatewayType = "ccavenue"
	Paytm     PaymentGatewayType = "paytm"
	PhonePe   PaymentGatewayType = "phonepe"
	BharatPe  PaymentGatewayType = "bharatpe"
)

// ------------------------ Global Configs ------------------------
var (
	DefaultCurrency   = getEnv("PAYMENT_CURRENCY", "INR")
	DefaultTimeoutSec = getEnvAsInt("PAYMENT_TIMEOUT_SEC", 15)

	RazorpayBaseURL  = getEnv("RAZORPAY_BASE_URL", "https://api.razorpay.com/v1")
	PayUMoneyBaseURL = getEnv("PAYUMONEY_BASE_URL", "https://secure.payu.in/_payment")
	CCAvenueBaseURL  = getEnv("CCAVENUE_BASE_URL", "https://secure.ccavenue.com/transaction/initiate")
	PaytmBaseURL     = getEnv("PAYTM_BASE_URL", "https://securegw.paytm.in/theia/processTransaction")
	PhonePeBaseURL   = getEnv("PHONEPE_BASE_URL", "https://merchants.phonepe.com/upi/pay")
	BharatPeBaseURL  = getEnv("BHARATPE_BASE_URL", "https://www.bharatpe.com/payment")
)

// ------------------------ Helpers ------------------------
func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	if val := os.Getenv(key); val != "" {
		var i int
		fmt.Sscanf(val, "%d", &i)
		if i > 0 {
			return i
		}
	}
	return defaultVal
}

// ------------------------ Config & Request ------------------------
type PaymentConfig struct {
	Type        PaymentGatewayType
	Key         string
	Secret      string
	MerchantID  string
	MerchantKey string
	CallbackURL string
	Environment string
	Currency    string
	Timeout     time.Duration
}

type PaymentRequest struct {
	OrderID    string
	Amount     float64
	CustomerID string
	Email      string
	Phone      string
	Product    string
}

// ------------------------ Gateway Interface ------------------------
type PaymentGateway interface {
	CreatePayment(ctx context.Context, req PaymentRequest) (string, error)
	VerifyPayment(ctx context.Context, payload map[string]string) (bool, error)
}

// ------------------------ Razorpay Gateway ------------------------
type RazorpayGateway struct{ Config PaymentConfig }

func (r *RazorpayGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	url := fmt.Sprintf("%s/orders", RazorpayBaseURL)
	payload := map[string]interface{}{
		"amount":          int(req.Amount * 100),
		"currency":        r.Config.Currency,
		"receipt":         req.OrderID,
		"payment_capture": 1,
		"notes": map[string]string{
			"customer_name": req.CustomerID,
			"email":         req.Email,
			"phone":         req.Phone,
		},
	}

	body, _ := json.Marshal(payload)
	request, _ := http.NewRequestWithContext(ctx, "POST", url, strings.NewReader(string(body)))
	request.SetBasicAuth(r.Config.Key, r.Config.Secret)
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: r.Config.Timeout}
	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("razorpay request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("razorpay error: %s", string(respBody))
	}

	var data map[string]interface{}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return "", fmt.Errorf("razorpay response parse error: %w", err)
	}

	id, ok := data["id"].(string)
	if !ok {
		return "", fmt.Errorf("razorpay order id missing")
	}
	return id, nil
}

func (r *RazorpayGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	orderID := payload["razorpay_order_id"]
	paymentID := payload["razorpay_payment_id"]
	signature := payload["razorpay_signature"]

	h := hmac.New(sha256.New, []byte(r.Config.Secret))
	h.Write([]byte(orderID + "|" + paymentID))
	expected := hex.EncodeToString(h.Sum(nil))

	if expected != signature {
		return false, fmt.Errorf("razorpay signature mismatch")
	}
	return true, nil
}

// ------------------------ PayUMoney Gateway ------------------------
type PayUMoneyGateway struct{ Config PaymentConfig }

func (p *PayUMoneyGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	form := url.Values{}
	form.Add("key", p.Config.Key)
	form.Add("txnid", req.OrderID)
	form.Add("amount", fmt.Sprintf("%.2f", req.Amount))
	form.Add("productinfo", req.Product)
	form.Add("firstname", req.CustomerID)
	form.Add("email", req.Email)
	form.Add("phone", req.Phone)
	form.Add("surl", p.Config.CallbackURL)
	form.Add("furl", p.Config.CallbackURL)

	hashStr := fmt.Sprintf("%s|%s|%.2f|%s|%s|||||||||||%s",
		p.Config.Key, req.OrderID, req.Amount, req.Product, req.CustomerID, p.Config.MerchantKey)
	h := sha512.Sum512([]byte(hashStr))
	form.Add("hash", hex.EncodeToString(h[:]))

	return fmt.Sprintf("%s?%s", PayUMoneyBaseURL, form.Encode()), nil
}

func (p *PayUMoneyGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	status := payload["status"]
	hash := payload["hash"]
	return status == "success" && hash != "", nil
}

// ------------------------ Placeholder Gateways ------------------------
type CCAvenueGateway struct{ Config PaymentConfig }

func (c *CCAvenueGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	return CCAvenueBaseURL, nil
}
func (c *CCAvenueGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	return true, nil
}

type PaytmGateway struct{ Config PaymentConfig }

func (p *PaytmGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	return PaytmBaseURL, nil
}
func (p *PaytmGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	return true, nil
}

type PhonePeGateway struct{ Config PaymentConfig }

func (p *PhonePeGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	return PhonePeBaseURL, nil
}
func (p *PhonePeGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	return true, nil
}

type BharatPeGateway struct{ Config PaymentConfig }

func (b *BharatPeGateway) CreatePayment(ctx context.Context, req PaymentRequest) (string, error) {
	return BharatPeBaseURL, nil
}
func (b *BharatPeGateway) VerifyPayment(ctx context.Context, payload map[string]string) (bool, error) {
	return true, nil
}

// ------------------------ Factory ------------------------
func NewGateway(cfg PaymentConfig) PaymentGateway {
	switch cfg.Type {
	case Razorpay:
		return &RazorpayGateway{Config: cfg}
	case PayUMoney:
		return &PayUMoneyGateway{Config: cfg}
	case CCAvenue:
		return &CCAvenueGateway{Config: cfg}
	case Paytm:
		return &PaytmGateway{Config: cfg}
	case PhonePe:
		return &PhonePeGateway{Config: cfg}
	case BharatPe:
		return &BharatPeGateway{Config: cfg}
	default:
		return nil
	}
}

// ------------------------ Helper Function ------------------------
func CreatePaymentOrder(
	gatewayType PaymentGatewayType,
	key, secret, callbackURL, currency string,
	orderID string,
	amount float64,
	customerID, email, phone, product string,
	timeout time.Duration,
) (string, error) {

	cfg := PaymentConfig{
		Type:        gatewayType,
		Key:         key,
		Secret:      secret,
		Currency:    currency,
		CallbackURL: callbackURL,
		Timeout:     timeout,
	}

	gateway := NewGateway(cfg)
	if gateway == nil {
		return "", fmt.Errorf("invalid payment gateway: %s", gatewayType)
	}

	order := PaymentRequest{
		OrderID:    orderID,
		Amount:     amount,
		CustomerID: customerID,
		Email:      email,
		Phone:      phone,
		Product:    product,
	}

	return gateway.CreatePayment(context.Background(), order)
}

// =========
// Uses
// ==========
// paymentID, err := helper.CreatePaymentOrder(
// 	helper.Razorpay,
// 	"RAZORPAY_KEY",
// 	"RAZORPAY_SECRET",
// 	"https://schoolerp.com/payment/callback",
// 	"INR",
// 	"ORDER123",
// 	1000.0,
// 	"John Doe",
// 	"john@example.com",
// 	"9876543210",
// 	"School Fee",
// 	15*time.Second,
// )

// if err != nil {
// 	fmt.Println("Payment creation failed:", err)
// 	return
// }

// fmt.Println("Redirect user to payment URL/order ID:", paymentID)
