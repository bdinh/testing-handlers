package testing_handlers

type User struct {
	ID 					int64  `json:"id"`
	Email     	string 	`json:"-"`
	Password 		string 	`json:"-"`
	FirstName 	string 	`json:"firstName"`
	LastName 		string 	`json:"lastName"`
	Description string 	`json:"description"`
}

