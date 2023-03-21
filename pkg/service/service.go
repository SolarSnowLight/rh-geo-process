package service

import (
	"geo-process/pkg/model"
	"geo-process/pkg/repository"
)

type Geo interface {
	GetRegionAll() ([]model.RegionModel, error)
	GetCitiesByRegion(regionId int) ([]model.CityDataModel, error)
	GetCities() ([]model.CityDataModel, error)
}

type Service struct {
	Geo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Geo: NewGeoService(repos),
	}
}
