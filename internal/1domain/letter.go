package domain

import "time"

type Letter struct {
	Id          int
	Creator     Employee
	CreatedAt   time.Time
	Date        time.Time
	Number      string
	From        Company
	To          Company
	Type        LetterType
	Projects    []Project
	Description string
}

type LetterType struct {
	Id       int
	TypeName string
}
