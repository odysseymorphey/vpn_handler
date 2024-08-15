package models

// TODO: Болванка. В будущем сделать норм структуру
type User struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}