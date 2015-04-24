package ui

import (
	"data"
	"fmt"
	"model"
	"strings"
)

func (uiConfig UiConfig) ToUrl() (url string) {
	switch page := uiConfig.Page.(type) {
	default:
		url = "/error"
	case Page_Home:
		url = "/"
	case Page_About:
		url = "/about"
	case Page_Projects:
		url = "/projects"
	case Page_Project:
		url = "/projects/" + page.toUrl()
	}
	return
}

func (p Page_Project) toUrl() (url string) {
	var purl string
	switch page := p.ProjectPage.(type) {
	default:
		purl = "/error"
	case ProjectPage_Home:
		purl = ""
	case ProjectPage_Issues:
		purl = "/issues"
	case ProjectPage_Issue:
		purl = "/issues/" + page.toUrl()
	}
	url = p.Project.Name + purl
	return
}

func (p ProjectPage_Issue) toUrl() (url string) {
	var purl string
	switch p.IssuePage.(type) {
	default:
		purl = "/error"
	case IssuePage_Home:
		purl = ""
	case IssuePage_Settings:
		purl = "/settings"
	}
	url = p.Issue.Title + purl
	return
}

func Route(url string, data *data.Data) (uiConfig UiConfig) {
	urlTrimmed := strings.Trim(url, "/")
	components := strings.Split(urlTrimmed, "/")
	var page Page
	switch components[0] {
	case "":
		page = Page_Home{}
	case "about":
		page = Page_About{}
	case "projects":
		if len(components) == 1 {
			page = Page_Projects{}
		} else {
			page = routeProjectPage(components[1:], data)
		}
	}
	if page == nil {
		page = Page_Error{url}
		fmt.Printf("ui.Route: %#v\n", url)
	}
	uiConfig = UiConfig{page}
	return
}

func routeProjectPage(path []string, data *data.Data) (projectPage Page_Project) {
	projectName := path[0]
	var project *model.Project
	for _, p := range data.Projects {
		if p.Name == projectName {
			project = p
		}
	}
	if len(path) == 1 {
		projectPage = Page_Project{project, ProjectPage_Home{}}
	} else if path[1] == "issues" {
		if len(path) == 2 {
			projectPage = Page_Project{project, ProjectPage_Issues{}}
		} else {
			projectPage = Page_Project{project, routeIssuePage(path[2:], project)}
		}
	}
	return
}

func routeIssuePage(path []string, project *model.Project) (issuePage ProjectPage_Issue) {
	issueName := path[0]
	var issue *model.Issue
	for _, i := range project.Issues {
		if i.Title == issueName {
			issue = i
		}
	}
	var x IssuePage
	if len(path) == 1 {
		x = IssuePage_Home{}
	} else {
		x = IssuePage_Settings{}
	}
	issuePage = ProjectPage_Issue{issue, x}
	return
}
