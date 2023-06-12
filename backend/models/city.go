package models

// Структура City представляет модель направлений (городов)
type City struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	// TODO Сделать привязку к отелям
}
