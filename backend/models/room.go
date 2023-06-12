package models

// Структура Room представляет модель номера отеля
type Room struct {
	ID          uint   `json:"id"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	BedPlaces   string `json:"bedplaces"`
	Places      string `json:"places"`
	AddPlaces   string `json:"addplaces"`
	Square      string `json:"square"`
	Usability   string `json:"usability"`
	Price       string `json:"price"`
}
