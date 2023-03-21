package service

import (
	"geo-process/pkg/model"
	"geo-process/pkg/repository"
)

type GeoService struct {
	repo repository.Geo
}

func NewGeoService(repo repository.Geo) *GeoService {
	return &GeoService{
		repo: repo,
	}
}

func (s *GeoService) GetRegionAll() ([]model.RegionModel, error) {
	return s.repo.GetRegionAll()
}

func (s *GeoService) GetCitiesByRegion(regionId int) ([]model.CityDataModel, error) {
	return s.repo.GetCitiesByRegion(regionId)
}

func (s *GeoService) GetCities() ([]model.CityDataModel, error) {
	return s.repo.GetCities()
}
