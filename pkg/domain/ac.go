package domain

import "time"

type AC struct {
	TableName string `json:"table_name"`
}

type ACData struct {
	ID          int       `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	IsEnabled   bool      `json:"is_enabled" db:"is_enabled"`
	Temperature float32   `json:"temperature" db:"temperature"`
	Humidity    float32   `json:"humidity" db:"humidity"`
}

type ACInfo struct {
	Enabled     bool    `json:"enabled"`
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
}
