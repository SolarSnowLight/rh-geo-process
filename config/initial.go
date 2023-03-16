package config

import (
	"geo-process/pkg/model"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
	"github.com/spf13/viper"
)

/* Инициализация элементов логгера */
func InitLogrus() ([]*os.File, error) {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	fileError, err := os.OpenFile(viper.GetString("paths.logs.error"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.AddHook(&writer.Hook{
			Writer: fileError,
			LogLevels: []logrus.Level{
				logrus.ErrorLevel,
			},
		})
	} else {
		logrus.SetOutput(os.Stderr)
		return nil, err
	}

	fileInfo, err := os.OpenFile(viper.GetString("paths.logs.info"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.AddHook(&writer.Hook{
			Writer: fileInfo,
			LogLevels: []logrus.Level{
				logrus.InfoLevel,
				logrus.DebugLevel,
			},
		})
	} else {
		logrus.SetOutput(os.Stderr)
		return nil, err
	}

	fileWarn, err := os.OpenFile(viper.GetString("paths.logs.warn"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.AddHook(&writer.Hook{
			Writer: fileWarn,
			LogLevels: []logrus.Level{
				logrus.WarnLevel,
			},
		})
	} else {
		logrus.SetOutput(os.Stderr)
		return nil, err
	}

	fileFatal, err := os.OpenFile(viper.GetString("paths.logs.fatal"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.AddHook(&writer.Hook{
			Writer: fileFatal,
			LogLevels: []logrus.Level{
				logrus.FatalLevel,
			},
		})
	} else {
		logrus.SetOutput(os.Stderr)
		return nil, err
	}

	return []*os.File{
		fileError,
		fileInfo,
		fileWarn,
		fileFatal,
	}, nil
}

/* Функция парсинга страницы википедии о регионах РФ */
func ParseSubjects() []model.RegionCityModel {
	var result []model.RegionCityModel
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{viper.GetString("rf_subject")},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			r.HTMLDoc.Find("tr").Each(func(i int, s *goquery.Selection) {
				lenChildren := len(s.Children().Find("span.nowrap").Nodes)
				if (len(s.Children().Nodes) == 6) && (s.Children().Get(0).Data == "td") && (lenChildren >= 1) {
					var text []string
					s.Find("td > span.nowrap").Each(func(i int, s *goquery.Selection) {
						text = append(text, strings.TrimSpace(s.Text()))
					})
					s.Find("td > a").Each(func(i int, s *goquery.Selection) {
						text = append(text, strings.TrimSpace(s.Text()))
					})
					result = append(result, model.RegionCityModel{
						Region: text[0],
						City:   text[len(text)-1],
					})

					/*g.Exports <- map[string]interface{}{ <- пример экспорта данных через карту интерфейсов
						"region_name": text[0],
						"city_name":   text[1],
					}*/
				}
			})
		},
		// Exporters: []export.Exporter{&export.JSON{}}, <- Экспорт данных в файл после парсинга
	}).Start()

	return result
}
