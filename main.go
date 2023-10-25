package main

import (
	"fmt"
	"pkg/projects"
)

func main() {
	project1 := projects.NewProject(1, "Project 1")
	fmt.Println(project1)
}
