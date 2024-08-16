package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template
	if app.config.useCache {
		if templateFromMap, ok := app.templateMap[t]; ok {
			tmpl = templateFromMap
		}
	}

	if tmpl == nil {
		newTemplate, err := app.buildTemplateFromDist(t)
		if err != nil {
			log.Println("error building template from dist:", err)
			return
		}
		log.Println("building template from disk")
		tmpl = newTemplate
	}
	if td == nil {
		td = &templateData{}
	}
	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("error executing template:", t, " ", err)
		return

	}

}

func (app *application) buildTemplateFromDist(t string) (*template.Template, error) {
	templateSlice := []string{
		"./templates/base.layout.html",
		"./templates/partials/header.partial.html",
		"./templates/partials/footer.partial.html",
		fmt.Sprintf("./templates/%s", t),
	}
	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		return nil, err
	}
	app.templateMap[t] = tmpl
	return tmpl, nil
}
