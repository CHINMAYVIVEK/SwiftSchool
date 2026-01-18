package helper

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"sync"
)

var logger = GetLogger()

type ViewType string

const (
	StudentView ViewType = "student"
	SchoolView  ViewType = "school"
	WebsiteView ViewType = "website"
)

//////////////////////////////////////////////////////
//                     Layouts                     //
//////////////////////////////////////////////////////

var layoutFiles = map[ViewType][]string{
	StudentView: {"base.html", "_header.html", "_sidebar.html", "_footer.html"},
	SchoolView:  {"base.html", "_header.html", "_sidebar.html", "_footer.html"},
}

var viewBasePaths = map[ViewType]string{
	StudentView: "template/student",
	SchoolView:  "template/school",
	WebsiteView: "template",
}

//////////////////////////////////////////////////////
//                     Cache                       //
//////////////////////////////////////////////////////

type TemplateCache struct {
	mu    sync.RWMutex
	items map[string]*template.Template
}

var templateCache = &TemplateCache{items: make(map[string]*template.Template)}

//////////////////////////////////////////////////////
//                     Helpers                     //
//////////////////////////////////////////////////////

var funcMap = template.FuncMap{
	"marshal": func(v any) template.JS {
		b, _ := json.Marshal(v)
		return template.JS(b)
	},
}

func cacheKey(view ViewType, page string) string {
	return string(view) + ":" + page
}

// buildTemplateFiles builds absolute paths for templates
func buildTemplateFiles(view ViewType, page string) ([]string, error) {
	basePath, ok := viewBasePaths[view]
	if !ok {
		return nil, fmt.Errorf("unknown view type: %s", view)
	}

	files := []string{}

	if layout, ok := layoutFiles[view]; ok {
		for _, f := range layout {
			files = append(files, basePath+"/"+f)
		}
	}

	// Page file must exist
	pageFile := basePath + "/" + page + ".html"
	if _, err := os.Stat(pageFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("page template does not exist: %s", pageFile)
	}
	files = append(files, pageFile)

	logger.Info("Template files:", files)
	return files, nil
}

//////////////////////////////////////////////////////
//                     Public API                  //
//////////////////////////////////////////////////////

func Render(w http.ResponseWriter, view ViewType, page string, data any) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	key := cacheKey(view, page)

	// Load from cache
	templateCache.mu.RLock()
	tmpl, ok := templateCache.items[key]
	templateCache.mu.RUnlock()

	if !ok {
		files, err := buildTemplateFiles(view, page)
		if err != nil {
			logger.Error("Build template files error:", err)
			return err
		}

		tmpl, err = template.New(page).Funcs(funcMap).ParseFiles(files...)
		if err != nil {
			logger.Error("Parse template files error:", err)
			return fmt.Errorf("template parse error: %w", err)
		}

		templateCache.mu.Lock()
		templateCache.items[key] = tmpl
		templateCache.mu.Unlock()
	}

	// Execute template
	if view == WebsiteView {
		// WebsiteView is standalone
		if err := tmpl.Execute(w, data); err != nil {
			logger.Error("Execute WebsiteView template error:", err)
			return err
		}
	} else {
		// Student/School view uses base.html as root
		if err := tmpl.ExecuteTemplate(w, "base.html", data); err != nil {
			logger.Error("Execute Student/School template error:", err)
			return err
		}
	}

	logger.Info("Template rendered successfully:", page)
	return nil
}
