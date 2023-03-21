package model

type CityDB struct {
	Id         uint   `gorm:"primaryKey;column:id"`
	Lat        string `json:"lat"`
	Lng        string `json:"lon"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Population int    `json:"population"`
}

func (CityDB) TableName() string {
	return "cities"
}
