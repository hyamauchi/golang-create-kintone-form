package common

import (
	"bytes"
	"text/template"
)

func EvalTemplate(t string, input map[string]string) string {
	tmpl, err := template.New("test").Parse(t)
	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	if err := tmpl.Execute(&doc, input); err != nil {
		panic(err)
	}

	return doc.String()
}
