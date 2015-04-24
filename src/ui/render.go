package ui

import (
	"data"
	"html/template"
	"model"
)

// render UiConfig
func (c UiConfig) RenderUiConfig(d data.Data) template.HTML {
	menuLinks := []SubUiConfigLink{
		SubUiConfigLink{"YellowGrass", Page_Home{}},
		SubUiConfigLink{"About", Page_About{}},
		SubUiConfigLink{"Projects", Page_Projects{}}}
	headerHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(menuLinks, c)))
	pageHtml := c.Page.renderPage(d, c)
	type templateData struct {
		HeaderHtml template.HTML
		PageHtml   template.HTML
	}
	return renderTemplate("main", templateData{headerHtml, pageHtml})
}

// render Page
func (p Page_Home) renderPage(d data.Data, c UiConfig) template.HTML {
	var projectList []SubUiConfigLink
	var issueList []SubUiConfigLink
	for _, project := range d.Projects {
		projectList = append(projectList, SubUiConfigLink{project.Name, Page_Project{project, ProjectPage_Home{}}})
		for _, issue := range project.Issues {
			issueList = append(issueList, SubUiConfigLink{project.Name + ": " + issue.Title, Page_Project{project, ProjectPage_Issue{issue, IssuePage_Home{}}}})
		}
	}
	projectListHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(projectList, c)))
	issueListHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(issueList, c)))
	type templateData struct {
		ProjectListHtml template.HTML
		IssueListHtml   template.HTML
	}
	return renderTemplate("page_home", templateData{projectListHtml, issueListHtml})
}
func (p Page_About) renderPage(d data.Data, c UiConfig) template.HTML {
	return renderTemplate("page_about", d)
}
func (p Page_Projects) renderPage(d data.Data, c UiConfig) template.HTML {
	var projectList []SubUiConfigLink
	for _, project := range d.Projects {
		projectList = append(projectList, SubUiConfigLink{project.Name, Page_Project{project, ProjectPage_Home{}}})
	}
	projectListHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(projectList, c)))
	type templateData struct {
		ProjectListHtml template.HTML
	}
	return renderTemplate("page_projects", templateData{projectListHtml})
}
func (p Page_Project) renderPage(d data.Data, c UiConfig) template.HTML {
	projectMenuLinks := []SubUiConfigLink{
		SubUiConfigLink{"Home", ProjectPage_Home{}},
		SubUiConfigLink{"Issues", ProjectPage_Issues{}}}
	projectMenuHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(projectMenuLinks, c)))
	projectPageHtml := p.ProjectPage.renderProjectPage(d, c, p)
	type templateData struct {
		Project         *model.Project
		ProjectMenuHtml template.HTML
		ProjectPageHtml template.HTML
	}
	return renderTemplate("page_project", templateData{p.Project, projectMenuHtml, projectPageHtml})
}
func (p Page_Error) renderPage(d data.Data, c UiConfig) template.HTML {
	return renderTemplate("page_error", p)
}

// render ProjectPage
func (pp ProjectPage_Home) renderProjectPage(d data.Data, c UiConfig, p Page_Project) template.HTML {
	return renderTemplate("projectpage_home", p.Project)
}
func (pp ProjectPage_Issues) renderProjectPage(d data.Data, c UiConfig, p Page_Project) template.HTML {
	var issueList []SubUiConfigLink
	for _, i := range p.Project.Issues {
		issueList = append(issueList, SubUiConfigLink{i.Title, ProjectPage_Issue{i, IssuePage_Home{}}})
	}
	issueListHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(issueList, c)))
	type templateData struct {
		IssueListHtml template.HTML
	}
	return renderTemplate("projectpage_issues", templateData{issueListHtml})
}
func (pp ProjectPage_Issue) renderProjectPage(d data.Data, c UiConfig, p Page_Project) template.HTML {
	issueMenuLinks := []SubUiConfigLink{
		SubUiConfigLink{"Home", IssuePage_Home{}},
		SubUiConfigLink{"Settings", IssuePage_Settings{}}}
	issueMenuHtml := HtmlLinksToHtml(LinksToHtmlLinks(toLink(issueMenuLinks, c)))
	issuePageHtml := pp.IssuePage.renderIssuePage(d, c, pp)
	type templateData struct {
		Issue         *model.Issue
		IssueMenuHtml template.HTML
		IssuePageHtml template.HTML
	}
	return renderTemplate("projectpage_issue", templateData{pp.Issue, issueMenuHtml, issuePageHtml})
}

// render IssuePage
func (ip IssuePage_Home) renderIssuePage(d data.Data, c UiConfig, i ProjectPage_Issue) template.HTML {
	return renderTemplate("issuepage_home", i.Issue)
}
func (ip IssuePage_Settings) renderIssuePage(d data.Data, c UiConfig, i ProjectPage_Issue) template.HTML {
	return renderTemplate("issuepage_settings", i.Issue)
}
