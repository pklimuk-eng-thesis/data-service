package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

//go:generate --name DBService --output mock_dbService.go
type DBService interface {
	GetLatestSensorDataByTableNameLimitN(tableName string, limit int) ([]domain.SensorData, error)
	AddNewRecordToSensorTable(tableName string, isEnabled bool, detected bool) error
	GetLatestDeviceDataByTableNameLimitN(tableName string, limit int) ([]domain.DeviceData, error)
	AddNewRecordToDeviceTable(tableName string, isEnabled bool) error
}

type dbService struct {
	DB *sqlx.DB
}

func NewDBService(db *sqlx.DB) DBService {
	return &dbService{DB: db}
}
