package model

import ()

type Issue struct {
	Title    string
	Body     string
	Project  *Project
	Comments []*Comment
}

func CreateIssue(title string, body string, project *Project) (i *Issue) {
	i = &Issue{title, body, project, nil}
	project.Issues = append(project.Issues, i)
	return
}
