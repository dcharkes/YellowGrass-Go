package ui

import (
	"data"
	"html/template"
	"model"
)

type SubUiConfig interface {
	toUiConfig(c UiConfig) UiConfig
	toUrl() string
}
type UiConfig struct {
	Page Page //case class Page
}

// case class Page
type Page interface {
	isPage()
	renderPage(d data.Data, c UiConfig) template.HTML
	SubUiConfig
}
type Page_Home struct{}
type Page_About struct{}
type Page_Projects struct{}
type Page_Project struct {
	Project     *model.Project
	ProjectPage ProjectPage //case class ProjectPage
}
type Page_Error struct {
	RequestUrl string
}

func (p Page_Home) isPage()     {}
func (p Page_About) isPage()    {}
func (p Page_Projects) isPage() {}
func (p Page_Project) isPage()  {}
func (p Page_Error) isPage()    {}

// case class ProjectPage
type ProjectPage interface {
	isProjectPage()
	renderProjectPage(d data.Data, c UiConfig, p Page_Project) template.HTML
	SubUiConfig
}
type ProjectPage_Home struct{}
type ProjectPage_Issues struct{}
type ProjectPage_Issue struct {
	Issue     *model.Issue
	IssuePage IssuePage //case class IssuePage
}

func (p ProjectPage_Home) isProjectPage()   {}
func (p ProjectPage_Issues) isProjectPage() {}
func (p ProjectPage_Issue) isProjectPage()  {}

// case class IssuePage
type IssuePage interface {
	isIssuePage()
	renderIssuePage(d data.Data, c UiConfig, i ProjectPage_Issue) template.HTML
	SubUiConfig
}
type IssuePage_Home struct{}
type IssuePage_Settings struct{}

func (p IssuePage_Home) isIssuePage()     {}
func (p IssuePage_Settings) isIssuePage() {}
