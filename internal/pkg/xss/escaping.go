package xss

import (
	"html/template"
	"net/http"
)

func RenderUserInput(w http.ResponseWriter, userInput string) {
	tmpl, _ := template.New("example").Parse("<html><body>{{.}}</body></html>")
	tmpl.Execute(w, template.HTMLEscapeString(userInput))
}
