package models

type City struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// TODO Сделать привязку к отелям
}
