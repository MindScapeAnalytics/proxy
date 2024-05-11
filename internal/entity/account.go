package entity

type Account struct {
	Id       string `json:"id"`
	Password string `json:"password"`
	Login    string `josn:"login"`
	Email    string `json:"email"`
}
