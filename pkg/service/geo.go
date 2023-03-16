package service

import (
	geoModel "geo-process/pkg/model"
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

func (s *GeoService) GetRegionAll() ([]geoModel.RegionDB, error) {
	return s.GetRegionAll()
}
