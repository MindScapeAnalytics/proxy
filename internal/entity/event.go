package entity

type Event struct {
	Id          string
	Name        string
	Description string
	Data        []byte
	UserId      string `db:"user_id"`
}