package Sarkor_test

type Phone struct {
	Id          int    `json:"id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"isFax"`
	UserId      int    `json:"userId"`
}

type PhoneDto struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsFax       bool   `json:"isFax"`
	UserId      int    `json:"userId"`
}

type UpdatePhone struct {
	Id          *int    `json:"id"`
	Phone       *string `json:"phone"`
	Description *string `json:"description"`
	IsFax       *bool   `json:"isFax"`
}
