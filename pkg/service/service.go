package service

import (
	geoModel "geo-process/pkg/model"
	"geo-process/pkg/repository"
)

type Geo interface {
	GetRegionAll() ([]geoModel.RegionDB, error)
}

type Service struct {
	Geo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Geo: NewGeoService(repos),
	}
}
