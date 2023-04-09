package service

import (
	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

//go:generate --name SensorService --output mock_sensorService.go
type SensorService interface {
	GetLatestSensorRecordsLimitN(limit int) ([]domain.SensorData, error)
	AddNewRecordToSensorTable(isEnabled bool, detected bool) error
}

type sensorService struct {
	sensor *domain.Sensor
	DB     db.DBService
}

func NewSensorService(sensor *domain.Sensor, db db.DBService) SensorService {
	return &sensorService{sensor: sensor, DB: db}
}

func (s *sensorService) GetLatestSensorRecordsLimitN(limit int) ([]domain.SensorData, error) {
	return s.DB.GetLatestSensorDataByTableNameLimitN(s.sensor.TableName, limit)
}

func (s *sensorService) AddNewRecordToSensorTable(isEnabled bool, detected bool) error {
	return s.DB.AddNewRecordToSensorTable(s.sensor.TableName, isEnabled, detected)
}
