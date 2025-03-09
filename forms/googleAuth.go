package forms

type GetQRForm struct {
	Username string `form:"username"`
}

type VerifyForm struct {
	Username string `form:"username"`
	Code     string `form:"code"`
}
