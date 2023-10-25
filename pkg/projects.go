package projects

import (
	"time"
)

type Project struct {
	id     int
	name   string
	stages []Stage
}

type Stage struct {
	name  string
	begin time.Time
	end   time.Time
}

func NewProject(id int, name string) Project {
	var s []Stage
	p := Project{id, name, s}
	return p
}
