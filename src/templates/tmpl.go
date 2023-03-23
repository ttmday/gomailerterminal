package tmpls

import (
	"bytes"
	"html/template"
)

func LoadTemplate(tmpl string) *template.Template {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		println("Error parsing template")
		panic(err)
	}

	return t
}

func LoadTemplateByString(name string, text string) *template.Template {
	t, err := template.New(name).Parse(text)

	if err != nil {
		println("Error parsing template")
		panic(err)
	}

	return t
}

func LoadTemplateWithFuncs(tmplName, tmplFolder string, funcsMap template.FuncMap) *template.Template {
	t, err := template.New(tmplName).Funcs(funcsMap).ParseFiles(tmplFolder + tmplName)

	if err != nil {
		println("Error parsing template")
		panic(err)
	}

	return t
}

func RenderTemplateByBuf(tmpl *template.Template, data interface{}) string {
	buf := new(bytes.Buffer)

	err := tmpl.Execute(buf, data)

	if err != nil {
		println("Error executing template")
		panic(err)
	}

	return buf.String()
}
