package form

type CategoryAddForm struct {
	Name			string		`validate:"required"`
	Pid          	uint		`json:"pid"`
	Status          uint		`json:"status"`
}

type CategoryUpdateForm struct {
	Id				uint		`validate:"required";json:"id"`
	Name			string		`validate:"required"`
	Pid          	uint		`json:"pid"`
	Status          uint		`json:"status"`
}

type CategoryUpdateStatusForm struct {
	Id				uint		`validate:"required";json:"id"`
	Status          uint		`json:"status"`
}

type CategoryFindForm struct {
	Name	string	`form:"name"`
	Pid		uint	`form:"pid"`
}