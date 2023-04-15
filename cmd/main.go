package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/pklimuk-eng-thesis/data-service/pkg/http"
	sensorHttp "github.com/pklimuk-eng-thesis/data-service/pkg/http/sensor"
	sensorService "github.com/pklimuk-eng-thesis/data-service/pkg/service/sensor"
)

func main() {
	postgresDB := db.NewPostgresDB()
	defer postgresDB.Close()
	dbService := db.NewDBService(postgresDB)

	r := gin.Default()

	initializeSensor("presence_sensor", dbService, r, "/presenceSensor")
	initializeSensor("gas_sensor", dbService, r, "/gasSensor")
	initializeSensor("doors_sensor", dbService, r, "/doorsSensor")

	serviceAddress := setupServiceAddress("ADDRESS", ":8087")
	err := r.Run(serviceAddress)
	if err != nil {
		panic(err)
	}
}

func initializeSensor(tableName string, db db.DBService, r *gin.Engine, groupName string) {
	sensor := domain.Sensor{TableName: tableName}
	sensorService := sensorService.NewSensorService(&sensor, db)
	sensorHandler := sensorHttp.NewSensorHandler(sensorService)
	http.SetupSensorRouter(r, sensorHandler, groupName)
}

func setupServiceAddress(identifier string, defaultAddress string) string {
	address := os.Getenv(identifier)
	if address == "" {
		address = defaultAddress
	}
	return address
}
