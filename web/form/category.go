package form

type CategoryAddForm struct {
	Name			string		`validate:"required"`
	Pid          	uint		`json:"pid"`
	Status          uint		`json:"status"`
}

type CategoryFindForm struct {
	Name	string	`json:"name"`
	Pid		uint	`json:"pid"`
}