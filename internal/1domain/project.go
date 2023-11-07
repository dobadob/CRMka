package domain

import (
	"time"
)

type Project struct {
	Id            int
	Type          ProjectType
	Name          string
	LastCondition Condition
	Stages        []Stage
	Employees     []EmployeeInProject
	Companies     []CompanyInProject
}

type ProjectType struct {
	Id         int
	NameOfType string
}

type Condition struct {
	Id          int
	ProjectId   int
	CreatedAt   time.Time
	Creator     Employee
	ConfirmedAt time.Time
	Confirmer   Employee
	Body        string
}

type Stage struct {
	Name  string
	Begin time.Time
	End   time.Time
	Done  bool
}

type EmployeeInProject struct {
	Employee Employee
	Role     RoleEmployeeInProject
}
type RoleEmployeeInProject struct {
	Id         int
	NameOfRole string
}

type CompanyInProject struct {
	Company Company
	Role    RoleCompanyInProject
}
type RoleCompanyInProject struct {
	Id         int
	NameOfRole string
}
