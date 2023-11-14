package dto

type CreateEmployeeDTO struct {
	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`
	Position   string `json:"position" bson:"position"`
}
