package ui

import (
	"bytes"
	"html/template"
)

var templates = template.Must(template.ParseFiles(
	"templates/issuepage_home.html",
	"templates/issuepage_settings.html",
	"templates/link.html",
	"templates/main.html",
	"templates/page_about.html",
	"templates/page_error.html",
	"templates/page_home.html",
	"templates/page_project.html",
	"templates/page_projects.html",
	"templates/projectpage_home.html",
	"templates/projectpage_issue.html",
	"templates/projectpage_issues.html"))

func renderTemplate(tmpl string, param interface{}) template.HTML {
	var b bytes.Buffer
	_ = templates.ExecuteTemplate(&b, tmpl+".html", param)
	return template.HTML(b.String())
}
