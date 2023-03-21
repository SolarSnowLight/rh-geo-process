package model

type RegionCenterDB struct {
	Id              uint `gorm:"primaryKey;autoIncrement;column:id"`
	CitiesRegionsId uint `gorm:"column:cities_regions_id"`
}

func (RegionCenterDB) TableName() string {
	return "regions_centers"
}
