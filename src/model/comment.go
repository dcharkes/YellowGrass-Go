package model

import ()

type Comment struct {
	Body  string
	Issue *Issue
}

func CreateComment(body string, i *Issue) (c *Comment) {
	c = &Comment{body, i}
	i.Comments = append(i.Comments, c)
	return
}
