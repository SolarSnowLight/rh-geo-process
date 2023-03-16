package repository

import (
	"encoding/json"
	"geo-process/pkg/model"
	"io/ioutil"

	"gorm.io/gorm"
)

type GeoPostgres struct {
	db *gorm.DB
}

func NewGeoPostgres(
	db *gorm.DB,
) *GeoPostgres {
	return &GeoPostgres{
		db: db,
	}
}

func (r *GeoPostgres) GetRegionAll() ([]model.RegionDB, error) {
	return nil, nil
}

/* Добавление регионов в БД */
func (r *GeoPostgres) AddRegionList(list []model.RegionDB) ([]model.RegionDB, error) {
	if err := r.db.Create(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

/* Добавление списка городов в БД */
func (r *GeoPostgres) AddCityList(filepath string, regionCity []model.RegionCityModel) error {
	// Чтение данных из файла
	data, _ := ioutil.ReadFile(filepath)
	var list []model.CityModel

	// Десереализация данных из фалйа
	err := json.Unmarshal(data, &list)
	if err != nil {
		return err
	}

	tx := r.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	for _, item := range list {
		// Определяем регион по его имени
		var region model.RegionDB
		if err := tx.First(&region, "name = ?", item.Subject).Error; err != nil {
			return err
		}

		// Конвертируем модель текущей итерации в удобную модель для БД и создаём новую запись (новый город)
		city := item.ConvertToDB()
		if err := tx.Create(city).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Создаём связь между регионом и городом
		cityRegion := &model.CityRegionDB{
			CitiesId:  city.Id,
			RegionsId: region.Id,
		}

		// Добавляем эту связь в БД
		if err := tx.Create(cityRegion).Error; err != nil {
			tx.Rollback()
			return err
		}

		// Если данный город является административным центром субъекта РФ, то добавляем его в БД
		if model.CityIsCenter(item, regionCity) {
			regionCenter := &model.RegionCenterDB{
				CitiesRegionsId: cityRegion.Id,
			}

			if err := tx.Create(regionCenter).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
