package modal

import (
	"bytes"
	_ "embed"
	"github.com/jfcarter2358/ui"
	"text/template"
)

//go:embed template.html
var TemplateString string

type Modal struct {
	ID         string
	Components []ui.Component
	Classes    string
	Style      string
	HXGet      string
	HXPost     string
	HXPut      string
	HXTarget   string
	HXSwap     string
	HXTrigger  string
	BoxClasses string
}

var Template *template.Template

func (m Modal) Render() (string, error) {
	if Template == nil {
		var err error
		Template, err = template.New("modal_html_template").Parse(TemplateString)
		if err != nil {
			return "", err
		}
	}
	var doc bytes.Buffer
	err := Template.Execute(&doc, m)
	return doc.String(), err
}
