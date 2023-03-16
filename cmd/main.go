package main

import (
	"context"
	mainServer "geo-process"
	initConfigure "geo-process/config"
	"geo-process/pkg/handler"
	model "geo-process/pkg/model"
	"geo-process/pkg/repository"
	"geo-process/pkg/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"github.com/samber/lo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

// @title Rental Housing
// @version 1.0
// description Проект для аренды жилья

// @host localhost:5001
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	// Инициализация файлов конфигурации
	if err := initConfig(); err != nil {
		logrus.Fatal("Ошибка инициализации файлов конфигураций: %s", err.Error())
	}

	// Загрузка переменных окружения
	if err := gotenv.Load(); err != nil {
		logrus.Fatalf("Ошибка инициализации переменных окружения: %s", err.Error())
	}

	openLogFiles, err := initConfigure.InitLogrus()
	if err != nil {
		logrus.Error("Ошибка при настройке параметров логгера. Вывод всех ошибок будет осуществлён в консоль")
	}

	// Закрытие всех открытых файлов в результате настройки логгера
	defer func() {
		for _, item := range openLogFiles {
			item.Close()
		}
	}()

	// Создание нового подключения к БД
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Database connection error: %s", err.Error())
	}

	// Dependency Injection
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(mainServer.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Geo process service started")

	// Парсим данные со страницы википедии
	resParse := initConfigure.ParseSubjects()

	// Добавление регионов (при условии, что они не были добавлены ранее)
	repos.AddRegionList(lo.Map(resParse, func(x model.RegionCityModel, index int) model.RegionDB {
		return model.RegionDB{
			Name: x.Region,
		}
	}))

	// Добавление городов РФ (при условии, что они не были добавлены ранее)
	// repos.AddCityList(viper.GetString("paths.cities"), resParse)

	// Блокировка функции main
	quit := make(chan os.Signal, 1)

	// запись в канал
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	logrus.Print("Geo process service shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

/* Инициализация файлов конфигурации */
func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
