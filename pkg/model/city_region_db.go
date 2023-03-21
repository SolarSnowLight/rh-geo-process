package model

type CityRegionDB struct {
	Id        uint `gorm:"primaryKey;autoIncrement;column:id"`
	CitiesId  uint `gorm:"column:cities_id"`
	RegionsId uint `gorm:"column:regions_id"`
}

func (CityRegionDB) TableName() string {
	return "cities_regions"
}
