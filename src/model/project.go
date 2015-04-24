package model

import ()

type Project struct {
	Name        string
	Description string
	Issues      []*Issue
}

func CreateProject(projName string, desc string) (proj *Project) {
	proj = &Project{projName, desc, nil}
	return
}
