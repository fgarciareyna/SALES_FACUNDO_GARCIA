package model

type User struct {
	Name string `json:"name"`
	Role string
	Pass string `json:"pass"`
}
