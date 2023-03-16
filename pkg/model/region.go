package model

type RegionDB struct {
	Id   uint   `gorm:"primaryKey;autoIncrement;column:id"`
	Name string `json:"name" gorm:"unique;column:name"`
}

func (RegionDB) TableName() string {
	return "regions"
}
