package http

import (
	"github.com/gin-gonic/gin"
	sensor "github.com/pklimuk-eng-thesis/data-service/pkg/http/sensor"
)

func SetupSensorRouter(r *gin.Engine, sensorHandler *sensor.SensorHandler, groupName string) {
	route := r.Group(groupName)
	route.GET("/latest", sensorHandler.GetLatestSensorRecordsLimitN)
	route.POST("/add", sensorHandler.AddNewRecordToSensorTable)
}
