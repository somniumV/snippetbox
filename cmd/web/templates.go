package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"snippetbox.somniumV.net/internal/models"
	"snippetbox.somniumV.net/ui"
	"time"
)

type templateData struct {
	Snippet       *models.Snippet
	Snippets      []*models.Snippet
	CurrentYear   int
	Form          any
	Authenticated bool
	Flash         string
	CSRFToken     string
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}
	return cache, nil
}
