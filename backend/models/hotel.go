package models

type Hotel struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Rating      string `json:"rating"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      string `json:"userid"`
	User        User   `json:"user";gorm:"foreignkey:UserID"`
}
