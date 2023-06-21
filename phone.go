package Sarkor_test

type Phone struct {
	Id          int    `json:"id"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
	IsFax       bool   `json:"isFax"`
	UserId      int    `json:"userId"`
}
