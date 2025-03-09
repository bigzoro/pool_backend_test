package models

type GoogleAuth struct {
	BaseModel
	Username string `json:"username"`
	Secret   string `json:"secret"`
	QrUrl    string `json:"url"`
}
