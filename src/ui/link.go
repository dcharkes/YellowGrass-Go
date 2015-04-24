package ui

import (
	"bytes"
	"html/template"
)

type Link struct {
	Name     string
	UiConfig UiConfig
}

type LinkHtml struct {
	Name string
	Url  string
}

func (link Link) LinkToLinkHtml() LinkHtml {
	return LinkHtml{link.Name, link.UiConfig.ToUrl()}
}

func LinksToHtmlLinks(links []Link) (htmlLinks []LinkHtml) {
	for _, l := range links {
		htmlLinks = append(htmlLinks, l.LinkToLinkHtml())
	}
	return
}

func HtmlLinksToHtml(htmlLinks []LinkHtml) template.HTML {
	linkTemplate, _ := template.ParseFiles("templates/link.html")
	var b bytes.Buffer
	for _, h := range htmlLinks {
		linkTemplate.Execute(&b, h)
	}
	return template.HTML(b.String())
}
