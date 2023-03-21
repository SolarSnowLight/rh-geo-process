package model

type CityDataModel struct {
	Id         uint    `json:"id" gorm:"column:id"`
	Lat        float64 `json:"lat" gorm:"column:lat"`
	Lng        float64 `json:"lng" gorm:"column:lng"`
	District   string  `json:"district" gorm:"column:district"`
	Name       string  `json:"name" gorm:"column:name"`
	Population float64 `json:"population" gorm:"column:population"`
}
