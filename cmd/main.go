package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	db "github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
	"github.com/pklimuk-eng-thesis/data-service/pkg/http"
	deviceHttp "github.com/pklimuk-eng-thesis/data-service/pkg/http/device"
	sensorHttp "github.com/pklimuk-eng-thesis/data-service/pkg/http/sensor"
	deviceService "github.com/pklimuk-eng-thesis/data-service/pkg/service/device"
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
	initializeDevice("smart_bulb", dbService, r, "/smartBulb")
	initializeDevice("smart_plug", dbService, r, "/smartPlug")

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

func initializeDevice(tableName string, db db.DBService, r *gin.Engine, groupName string) {
	device := domain.Device{TableName: tableName}
	deviceService := deviceService.NewDeviceService(&device, db)
	deviceHandler := deviceHttp.NewDeviceHandler(deviceService)
	http.SetupDeviceRouter(r, deviceHandler, groupName)
}

func setupServiceAddress(identifier string, defaultAddress string) string {
	address := os.Getenv(identifier)
	if address == "" {
		address = defaultAddress
	}
	return address
}
