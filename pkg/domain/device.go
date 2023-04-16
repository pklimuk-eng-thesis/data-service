package domain

import "time"

type Device struct {
	TableName string `json:"table_name"`
}

type DeviceData struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsEnabled bool      `json:"is_enabled" db:"is_enabled"`
}

type DeviceInfo struct {
	Enabled bool `json:"enabled"`
}
