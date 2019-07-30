package form

type CategoryAddForm struct {
	Name			string		`validate:"required"`
	Pid          	uint		`validate:"required"`
	Status          uint		`validate:"required"`
}

type CategoryFindForm struct {
	Name	string	`json:"name"`
	Pid		uint	`json:"pid"`
}