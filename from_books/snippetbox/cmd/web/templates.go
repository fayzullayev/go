package main

import (
	"github.com/fayzullayev/snippetbox/internal/models"
	"github.com/fayzullayev/snippetbox/ui"
	"html/template"
	"io/fs"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear     int
	Snippet         *models.Snippet
	Snippets        []*models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (TemplateCacheType, error) {
	cache := TemplateCacheType{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		//fmt.Println(index, "page : ", page, name)
		//files := []string{
		//	"./ui/html/base.tmpl.html",
		//	"./ui/html/partials/nav.tmpl.html",
		//	page,
		//}

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err2 := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)

		if err2 != nil {
			return nil, err2
		}

		cache[name] = ts
	}

	return cache, nil
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}
