package domain

import "time"

type Task struct {
	Id         int
	ProjectId  int
	Creator    Employee
	Executor   Employee
	Created_at time.Time
	Done_at    time.Time
	Body       string
	Done       bool
}
