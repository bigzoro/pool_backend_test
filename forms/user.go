package forms

type EmailRegisterForm struct {
	RegisterType string `json:"register_way"`
	Email        string `json:"email"`
	NickName     string `json:"nick_name"`
	Password     string `json:"password"`
	Username     string `json:"username"`
	InviteCode   string `json:"invite_code"`
	Code         string `json:"code"`
}

type EmailLoginForm struct {
	LoginType string `json:"login_type"`
	NickName  string `json:"nick_name"`
	Password  string `json:"password"`
	Username  string `json:"username"`
}

type UpdatePasswordForm struct {
	UserId   int    `json:"user_id"`
	Password string `json:"password"`
}

type UpdateWithdrawPasswordForm struct {
	UserId           int    `json:"user_id"`
	WithdrawPassword string `json:"withdraw_password"`
}

type UpdateEmailForm struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
}

type UserInfoForm struct {
	UserId int `json:"user_id"`
}
