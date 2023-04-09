package http

import "github.com/gin-gonic/gin"

func SetupSensorRouter(r *gin.Engine, sensorHandler *SensorHandler, groupName string) {
	route := r.Group(groupName)
	route.GET("/latest", sensorHandler.GetLatestSensorRecordsLimitN)
	route.POST("/add", sensorHandler.AddNewRecordToSensorTable)
}
