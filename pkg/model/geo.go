package model

type RegionCityModel struct {
	Region string `json:"region" binding:"required"`
	City   string `json:"city" binding:"required"`
}

func CityIsCenter(city CityModel, regionList []RegionCityModel) bool {
	for _, item := range regionList {
		if item.City == city.Name {
			return true
		}
	}

	return false
}

type CityModel struct {
	Coords struct {
		Lat string `json:"lat"`
		Lon string `json:"lon"`
	} `json:"coords"`
	District   string `json:"district"`
	Name       string `json:"name"`
	Population int    `json:"population"`
	Subject    string `json:"subject"`
}

func (cm *CityModel) ConvertToDB() *CityDB {
	return &CityDB{
		Lat:        cm.Coords.Lat,
		Lng:        cm.Coords.Lon,
		Name:       cm.Name,
		Population: cm.Population,
		District:   cm.District,
	}
}
