package form

type UserLoginForm struct {
	Name			string		`validate:"required"`
	Password			string		`validate:"required"`
}

