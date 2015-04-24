package ui

import (
	"data"
	"fmt"
	"model"
	"strings"
)

func (c UiConfig) ToUrl() (url string) { return c.Page.toUrl() }

func (p Page_Home) toUrl() string     { return "/" }
func (p Page_About) toUrl() string    { return "/about" }
func (p Page_Projects) toUrl() string { return "/projects" }
func (p Page_Project) toUrl() string  { return "/projects/" + p.Project.Name + p.ProjectPage.toUrl() }
func (p Page_Error) toUrl() string    { return "/error" }

func (p ProjectPage_Home) toUrl() string   { return "" }
func (p ProjectPage_Issues) toUrl() string { return "/issues" }
func (p ProjectPage_Issue) toUrl() string  { return "/issues/" + p.Issue.Title + p.IssuePage.toUrl() }

func (i IssuePage_Home) toUrl() string     { return "" }
func (i IssuePage_Settings) toUrl() string { return "/settings" }

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
