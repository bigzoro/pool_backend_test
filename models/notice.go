package models

type Notice struct {
	BaseModel
	Content string `json:"content"`
	IsShow  bool   `json:"is_show"`
}
