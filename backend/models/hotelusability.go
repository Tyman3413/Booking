package models

// Структура HotelUsability представляет модель удобств отеля
type HotelUsability struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Icon    string `json:"icon"`
	HotelID int    `json:"hotelid"`
	Hotel   Hotel  `json:"hotel";gorm:"foreignkey:HotelID"`
}
