package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func Parse(filepath string) (Template, error) {
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

type Template struct {
	htmlTpl *template.Template
}

// assume that template input has been parsed
func (t Template) Execute(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Executing template %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}
