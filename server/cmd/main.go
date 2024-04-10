package main

import (
	"CarCatalog/configuration"
	"CarCatalog/handler"
	"CarCatalog/model"
	"CarCatalog/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

// Путь к файлу конфигурации.
const configPath = "configuration.json"

func main() {
	// Получаем конфиг.
	config := configuration.GetConfig(configPath)

	log.Printf("Config: %v", config)

	// Подключаемся к базе данных.
	dataBase, err := sqlx.Open("postgres", getDSN(config.DBConf))
	if err != nil {
		panic(err)
	}

	log.Print("Database connected")

	server := handler.Server{
		Config: config,
		Storage: &storage.Storage{
			DB: dataBase,
		},
	}

	runServer(server)
}

func runServer(server handler.Server) {
	engine := gin.Default()

	engine.POST("/add_cars", server.AddCarsHandler)
	engine.POST("/modify_car", server.ModifyCarsHandler)
	engine.GET("/get_cars", server.GetCarsHandler)
	engine.DELETE("/delete_car", server.DeleteCarHandler)

	port := ":" + strconv.Itoa(server.Config.Port)

	err := engine.Run(port)
	if err != nil {
		panic(err)
	}
}

// getDSN функция создаёт строку Data Source Name.
func getDSN(cfg model.DBConf) string {
	return fmt.Sprintf("host=localhost port=5432 user=%s password=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.DBName, cfg.Sslmode)
}
