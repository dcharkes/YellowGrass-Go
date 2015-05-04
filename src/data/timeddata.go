package data

import (
	"time"
	"model"
	"fmt"
)

func AddProjectsOverTime() {
	i := 0;
	for true {
		time.Sleep(4 * time.Second)
		i++
		p := model.CreateProject(fmt.Sprintf("Project %d", i), fmt.Sprintf("Description %d", i))
		D.addProject(p)
	}
}
