package model

import (
	"fmt"
)

type Project struct {
	Name        string
	Description string
	Issues      []*Issue `json:"-"`
}

func CreateProject(projName string, desc string) (proj *Project) {
	proj = &Project{projName, desc, nil}
	return
}

func (self *Project) String() string {
	return fmt.Sprintf("%#v", self)
}
