package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	host         string
	port         int
	user         string
	password     string
	dbName       string
	maxOpenConns int
	maxIdleConns int
	sslMode      string
}

func (pDB *PostgresDB) initializeDbParams() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	pDB.host = os.Getenv("POSTGRES_HOST")
	pDB.port, err = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		log.Fatal("Could not convert POSTGRES_PORT to int")
	}
	pDB.user = os.Getenv("POSTGRES_USER")
	pDB.password = os.Getenv("POSTGRES_PASSWORD")
	pDB.dbName = os.Getenv("POSTGRES_DB")
	pDB.maxOpenConns, err = strconv.Atoi(os.Getenv("POSTGRES_MAX_OPEN_CONNS"))
	if err != nil {
		log.Fatal("Could not convert POSTGRES_MAX_OPEN_CONNS to int")
	}
	pDB.maxIdleConns, err = strconv.Atoi(os.Getenv("POSTGRES_MAX_IDLE_CONNS"))
	if err != nil {
		log.Fatal("Could not convert POSTGRES_MAX_IDLE_CONNS to int")
	}
	pDB.sslMode = os.Getenv("POSTGRES_SSL_MODE")
}

func NewPostgresDB() *sqlx.DB {
	postgresDB := PostgresDB{}
	postgresDB.initializeDbParams()
	postgresDBInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		postgresDB.host, postgresDB.port, postgresDB.user, postgresDB.password, postgresDB.dbName, postgresDB.sslMode)
	db, err := sqlx.Connect("postgres", postgresDBInfo)
	if err != nil {
		log.Fatal("Could not connect to DB\n", err)
	}
	db.SetMaxOpenConns(postgresDB.maxOpenConns)
	db.SetMaxIdleConns(postgresDB.maxIdleConns)
	return db
}
