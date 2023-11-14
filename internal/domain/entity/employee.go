package entity

type Employee struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	FirstName  string `json:"firstName" bson:"firstName"`
	MiddleName string `json:"middleName" bson:"middleName"`
	LastName   string `json:"lastName" bson:"lastName"`
	Position   string `json:"position" bson:"position"`
	//Username     string `json:"username" bson:"username"`
	//PasswordHash string `json:"-" bson:"password"`
}
