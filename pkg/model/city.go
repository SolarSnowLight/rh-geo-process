package model

type CityDB struct {
	Id         uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Lat        string `json:"lat" gorm:"unique"`
	Lng        string `json:"lon" gorm:"unique"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func (CityDB) TableName() string {
	return "cities"
}
