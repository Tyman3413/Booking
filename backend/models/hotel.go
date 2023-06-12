package models

// Структура Hotel представляет модель отеля с различными связями с другими таблицами в базе данных
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
	Room        Room   `json:"room";gorm:"foreignkey:RoomID"`
	CityID      string `json:"cityid"`
	City        City   `json:"city";gorm:"foreignkey:CityID"`

	// ! Для процесса добавления пользователем нового объекта
	UserID string `json:"userid"`
	User   User   `json:"user";gorm:"foreignkey:UserID"`

	// TODO Связать с таблицей Удобства
}
