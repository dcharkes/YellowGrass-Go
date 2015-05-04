package data

import (
	"model"
)

type Data struct {
	Projects []*model.Project
}

var D Data

func (data *Data) addProject(project *model.Project) {
	data.Projects = append(data.Projects, project)
}

func CreateData() (data Data) {
	data = Data{}
	p1 := model.CreateProject("Spoofax", "Some description about Spoofax.")
	p2 := model.CreateProject("WebDSL", "WebDSL is very nice.")
	data.addProject(p1)
	data.addProject(p2)
	i1 := model.CreateIssue("Tabs Vs Spaces", "Should we do tabs or spaces? Or maybe a combination?", p1)
	_ = model.CreateIssue("Occurence Heighlighting", "Bla bla bla", p1)
	_ = model.CreateComment("Well, maybe both, should be a config option.", i1)
	_ = model.CreateComment("We need tabs for indentation and spaces for layout!", i1)
	return
}
