package db

import (
	"fmt"

	dbUtils "github.com/pklimuk-eng-thesis/data-service/pkg/db/utils"
	"github.com/pklimuk-eng-thesis/data-service/pkg/domain"
)

func (s *dbService) GetLatestACDataByTableNameLimitN(tableName string, limit int) ([]domain.ACData, error) {
	query := fmt.Sprintf("SELECT * FROM smart_home.%s ORDER BY created_at DESC LIMIT $1", tableName)
	rows, err := s.DB.Queryx(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var acs []domain.ACData
	for rows.Next() {
		var id int
		var createdAt string
		var isEnabled bool
		var temperature float32
		var humidity float32
		err = rows.Scan(&id, &createdAt, &isEnabled, &temperature, &humidity)
		if err != nil {
			return nil, err
		}
		createdAtTime, err := dbUtils.ParseTime(createdAt)
		if err != nil {
			return nil, err
		}
		acs = append(acs, domain.ACData{ID: id, CreatedAt: createdAtTime, IsEnabled: isEnabled, Temperature: temperature, Humidity: humidity})
	}
	return acs, nil
}

func (s *dbService) AddNewRecordToACTable(tableName string, isEnabled bool, temperature float32, humidity float32) error {
	query := fmt.Sprintf("INSERT INTO smart_home.%s (is_enabled, temperature, humidity) VALUES ($1, $2, $3)", tableName)
	_, err := s.DB.Exec(query, isEnabled, temperature, humidity)
	if err != nil {
		return err
	}
	return nil
}
