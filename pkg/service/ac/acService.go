package service

import (
	"github.com/pklimuk-eng-thesis/data-service/pkg/db"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

//go:generate --name ACService --output mock_acService.go
type ACService interface {
	GetLatestACRecordsLimitN(limit int) ([]domain.ACData, error)
	AddNewRecordToACTable(isEnabled bool, temperature float32, humidity float32) error
}

type acService struct {
	ac *domain.AC
	DB db.DBService
}

func NewACService(ac *domain.AC, db db.DBService) ACService {
	return &acService{ac: ac, DB: db}
}

func (s *acService) GetLatestACRecordsLimitN(limit int) ([]domain.ACData, error) {
	return s.DB.GetLatestACDataByTableNameLimitN(s.ac.TableName, limit)
}

func (s *acService) AddNewRecordToACTable(isEnabled bool, temperature float32, humidity float32) error {
	return s.DB.AddNewRecordToACTable(s.ac.TableName, isEnabled, temperature, humidity)
}
