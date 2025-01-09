package accordion

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/jfcarter2358/ui"
)

//go:embed template.html
var TemplateString string

type AccordionItem struct {
	Components []ui.Component
	Title      string
	Classes    string
	Style      string
	HXGet      string
	HXPost     string
	HXPut      string
	HXTarget   string
	HXSwap     string
	HXTrigger  string
}

type Accordion struct {
	ID             string
	Items          []AccordionItem
	Classes        string
	ContentClasses string
	ContentStyle   string
	Style          string
	HXGet          string
	HXPost         string
	HXPut          string
	HXTarget       string
	HXSwap         string
	HXTrigger      string
}

var Template *template.Template

func (a Accordion) Render() (string, error) {
	if Template == nil {
		var err error
		Template, err = template.New("accordion_html_template").Parse(TemplateString)
		if err != nil {
			return "", err
		}
	}
	var doc bytes.Buffer
	err := Template.Execute(&doc, a)
	return doc.String(), err
}
