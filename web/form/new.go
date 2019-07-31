package form

type NewAddForm struct {
	Cid			uint		`validate:"required";json:"cid"`
	Title		string		`validate:"required";json:"title"`
	Author		string		`validate:"required";json:"author"`
	Content		string		`validate:"required";json:"content"`
	Keywords	string		`validate:"required";json:"keywords"`
	Description	string		`validate:"required";json:"description"`
	Status		uint		`json:"status"`
}

type NewUpdateForm struct {
	Id			uint		`validate:"required";json:"id"`
	Cid			uint		`validate:"required";json:"cid"`
	Title		string		`validate:"required";json:"title"`
	Author		string		`validate:"required";json:"author"`
	Content		string		`validate:"required";json:"content"`
	Keywords	string		`validate:"required";json:"keywords"`
	Description	string		`validate:"required";json:"description"`
	Status		uint		`json:"status"`
}

type NewUpdateStatusForm struct {
	Id				uint		`validate:"required";json:"id"`
	Status          uint		`json:"status"`
}

type NewSearchForm struct {
	Page			uint		`form:"page"`
	PageSize		uint		`form:"pageSize"`
	Cid				uint		`form:"cid"`
}