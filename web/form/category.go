package form

type CategoryForm struct {
	Name			string   `validate:"required"`
	PantId          int64    `validate:"required"`
	Status          int64    `validate:"required"`
}