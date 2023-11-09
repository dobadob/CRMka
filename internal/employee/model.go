package employee

type Employee struct {
	Id           string `json:"id" bson:"_id,omitempty"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"password"`
	FirstName    string `json:"firstName" bson:"firstName"`
	MiddleName   string `json:"middleName" bson:"middleName"`
	LastName     string `json:"lastName" bson:"lastName"`
	Position     string `json:"position" bson:"position"`
}

type CreateEmployeeDTO struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Position   string `json:"position"`
}
