package models

// NOTE: Пока так, мб ваще будем через квери параметры
// данные юзера передавать
type PostReq struct {
	Name string `json:"name"`
}