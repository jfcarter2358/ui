package sidebar

import (
	"bytes"
	_ "embed"
	"github.com/jfcarter2358/ui"
	"text/template"
)

//go:embed template.html
var TemplateString string

type Sidebar struct {
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
}

var Template *template.Template

func (s Sidebar) Render() (string, error) {
	if Template == nil {
		var err error
		Template, err = template.New("sidebar_html_template").Parse(TemplateString)
		if err != nil {
			return "", err
		}
	}
	var doc bytes.Buffer
	err := Template.Execute(&doc, s)
	return doc.String(), err
}
