package ui

import (
	"fmt"
)

type SubUiConfigLink struct {
	Name        string
	SubUiConfig SubUiConfig
}

func (s SubUiConfigLink) toLink(c UiConfig) Link {
	return Link{s.Name, s.SubUiConfig.toUiConfig(c)}
}
func toLink(ss []SubUiConfigLink, c UiConfig) (ls []Link) {
	for _, s := range ss {
		ls = append(ls, s.toLink(c))
	}
	return
}

func pageToUiConfig(p Page, c UiConfig) UiConfig {
	c.Page = p
	return c
}
func (p Page_Home) toUiConfig(c UiConfig) UiConfig     { return pageToUiConfig(p, c) }
func (p Page_About) toUiConfig(c UiConfig) UiConfig    { return pageToUiConfig(p, c) }
func (p Page_Projects) toUiConfig(c UiConfig) UiConfig { return pageToUiConfig(p, c) }
func (p Page_Project) toUiConfig(c UiConfig) UiConfig  { return pageToUiConfig(p, c) }
func (p Page_Error) toUiConfig(c UiConfig) UiConfig    { return pageToUiConfig(p, c) }

func projectPageToUiConfig(projectPage ProjectPage, c UiConfig) UiConfig {
	switch page := c.Page.(type) {
	default:
		fmt.Printf("ui.ProjectPageToUiConfig: %#v\n", c)
	case Page_Project:
		page.ProjectPage = projectPage
		c.Page = page
	}
	return c
}
func (p ProjectPage_Home) toUiConfig(c UiConfig) UiConfig   { return projectPageToUiConfig(p, c) }
func (p ProjectPage_Issues) toUiConfig(c UiConfig) UiConfig { return projectPageToUiConfig(p, c) }
func (p ProjectPage_Issue) toUiConfig(c UiConfig) UiConfig  { return projectPageToUiConfig(p, c) }

func issuePageToUiConfig(issuePage IssuePage, c UiConfig) UiConfig {
	switch page := c.Page.(type) {
	default:
		fmt.Printf("ui.IssuePageToUiConfig: %#v\n", c)
	case Page_Project:
		switch page2 := page.ProjectPage.(type) {
		default:
			fmt.Printf("ui.IssuePageToUiConfig: %#v\n", c)
		case ProjectPage_Issue:
			page2.IssuePage = issuePage
			page.ProjectPage = page2
			c.Page = page
		}
	}
	return c
}
func (p IssuePage_Home) toUiConfig(c UiConfig) UiConfig     { return issuePageToUiConfig(p, c) }
func (p IssuePage_Settings) toUiConfig(c UiConfig) UiConfig { return issuePageToUiConfig(p, c) }
