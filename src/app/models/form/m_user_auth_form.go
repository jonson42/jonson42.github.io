package form

type MUserAuthForm struct {
	UserName string `form:"userName"`
	Password string `form:"password"`
}

