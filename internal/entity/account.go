package entity

type Account struct {
	Token    string `json:"token"`
	Id       string `json:"id"`
	Password string `json:"password"`
	Login    string `josn:"login"`
	Email    string `json:"email"`
}
