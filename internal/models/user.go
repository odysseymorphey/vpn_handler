package models

// NOTE: Эту структуру отдаем по GET запросу
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IPAdderess string `json:"ip_address"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}
