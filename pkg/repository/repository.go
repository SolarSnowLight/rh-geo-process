package repository

import (
	"geo-process/pkg/model"

	"gorm.io/gorm"
)

type Geo interface {
	GetRegionAll() ([]model.RegionModel, error)
	GetCitiesByRegion(regionid int) ([]model.CityDataModel, error)
	GetCities() ([]model.CityDataModel, error)
	AddRegionList(list []model.RegionDB) ([]model.RegionDB, error)
	AddCityList(filepath string, regionCity []model.RegionCityModel) error
}

type Repository struct {
	Geo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Geo: NewGeoPostgres(db),
	}
}
