package templates

import (
	"bytes"
	"html/template"
	"path/filepath"
)

func Render(templateName string, data any) (string, error) {
	base := "templates/base.html"
	page := filepath.Join("templates", templateName+".html")

	tpl, err := template.ParseFiles(base, page)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	return buf.String(), err
}