package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	constant "geo-process/pkg/constant"
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

/* Метод получения всех регионов */
func (r *GeoPostgres) GetRegionAll() ([]model.RegionModel, error) {
	var middleResult []string
	var result []model.RegionModel
	/*if err := r.db.Raw(
		fmt.Sprintf(`
			SELECT c.*, r.name AS region_name, r.id AS region_id FROM %s cr
			INNER JOIN %s rc ON rc.cities_regions_id = cr.id
			INNER JOIN %s r ON r.id = cr.regions_id
			INNER JOIN %s c ON c.id = cr.cities_id
			ORDER BY r.id
		`, constant.CITIES_REGIONS, constant.REGIONS_CENTERS, constant.REGIONS, constant.CITIES),
	).Scan(&result).Error; err != nil {
		return nil, err
	}*/

	if err := r.db.Raw(
		fmt.Sprintf(`
			SELECT row_to_json(nested_data, true) FROM
			(
				SELECT row_to_json(c.*, true) AS main_city, r.name, r.id FROM %s cr
				INNER JOIN %s rc ON rc.cities_regions_id = cr.id
				INNER JOIN %s r ON r.id = cr.regions_id
				INNER JOIN %s c ON c.id = cr.cities_id
				ORDER BY r.id
			) AS nested_data
		`, constant.CITIES_REGIONS, constant.REGIONS_CENTERS, constant.REGIONS, constant.CITIES),
	).Scan(&middleResult).Error; err != nil {
		return nil, err
	}

	for _, item := range middleResult {
		var body model.RegionModel
		if err := json.Unmarshal([]byte(item), &body); err != nil {
			return nil, err
		}

		result = append(result, body)
	}

	return result, nil
}

/* Метод получения всех городов по принадлежности к определённому региону */
func (r *GeoPostgres) GetCitiesByRegion(regionId int) ([]model.CityDataModel, error) {
	var result []model.CityDataModel
	if err := r.db.Raw(
		fmt.Sprintf(`
			SELECT c.* FROM %s c
			INNER JOIN %s cr ON cr.cities_id = c.id
			INNER JOIN %s r ON cr.regions_id = r.id
			WHERE r.id = ?
		`, constant.CITIES, constant.CITIES_REGIONS, constant.REGIONS),
		regionId,
	).Scan(&result).Error; err != nil {
		return nil, err
	}

	if result == nil {
		return nil, errors.New("Ошибка: данный регион не обнаружен в БД")
	}

	return result, nil
}

/* Получение списка всех городов РФ */
func (r *GeoPostgres) GetCities() ([]model.CityDataModel, error) {
	var result []model.CityDataModel
	if err := r.db.Raw(
		fmt.Sprintf(`
			SELECT c.* FROM %s c
		`, constant.CITIES),
	).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

/* Добавление регионов в БД */
func (r *GeoPostgres) AddRegionList(list []model.RegionDB) ([]model.RegionDB, error) {
	if err := r.db.Create(&list).Error; err != nil {
		return nil, err
	}

	r.db.Model(&model.RegionDB{}).Where("name = ?", "Владикавказ").Update("name", "Северная Осетия")

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

/*
SELECT c.name AS city_name, r.name AS region_name FROM cities_regions cr
INNER JOIN regions_centers rc ON rc.cities_regions_id = cr.id
INNER JOIN regions r ON r.id = cr.regions_id
INNER JOIN cities c ON c.id = cr.cities_id
ORDER BY r.id;
*/
