package main

import (
	"os"
	"strings"
	"html/template"
)

type Course struct {
	Name string
	Hour int
}

type Courses []Course

func main() {
	htmlTemplates := []string{
		"important-packages/templates/header.html",
		"important-packages/templates/content.html",
		"important-packages/templates/footer.html",
	}

	tmp := template.Must(
		// template.New("MyTemplate").Parse("Course: {{ .Name }} - Hour: {{ .Hour }}"),
		template.New("content.html").Funcs(template.FuncMap{"ToUpper": strings.ToUpper}).ParseFiles(htmlTemplates...),
	)

	//If you working with http, replace os.Stdout for http.ResponseWriter
	err := tmp.Execute(os.Stdout, Courses{
		{ "Go", 40 },
		{ "Python", 30 },
		{ "PHP", 20 },
	})

	if err != nil {
		panic(err)
	}
}