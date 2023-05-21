package models

type Property struct {
	MyModel
	Id          int    `json:"id" gorm:"primary_key"`
	UserId      int    `json:"userId"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	Photo       string `json:"photo" gorm:"nullable"`
	City        string `json:"city"`
	Country     string `json:"country"`
	County      string `json:"county"`
	Property_id int    `json:"property_id"`
	Status      string `json:"status"`
	Size        string `json:"size"`
	Bedrooms    string `json:"bedrooms"`
	Bathrooms   string `json:"bathrooms"`
}
