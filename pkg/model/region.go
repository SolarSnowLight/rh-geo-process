package model

type RegionModel struct {
	Id       uint          `json:"id"`
	Name     string        `json:"name"`
	MainCity CityDataModel `json:"main_city"`
}
