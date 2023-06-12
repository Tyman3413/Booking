package models

type Hotel struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Reviews     string `json:"reviews"`
	Image       string `json:"image"`
	RoomID      string `json:"roomid"`
	Room        Room   `json:"room";gorm:"foreignkey:RoomID"` // Связка с таблицей Номеров отелей
	CityID      string `json:"cityid"`
	City        City   `json:"city";gorm:"foreignkey:CityID"` //

	// Скип
	UserID string `json:"userid"`
	User   User   `json:"user";gorm:"foreignkey:UserID"`

	// TODO Связать с таблицей Удобства
}
