package crud

import (
	"database/sql"
	// "github.com/joho/godotenv"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB_TYPE     string
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_SSLMODE  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		slog.Info("【DB】This is a Production Environment")
	}

	DB_TYPE = os.Getenv("DB_TYPE")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_HOST = os.Getenv("DB_HOST")
	DB_SSLMODE = os.Getenv("DB_SSLMODE")
}

func Connect() *sql.DB {
	dsn := DB_TYPE + "://" + DB_USERNAME + ":" + DB_PASSWORD + "@" + DB_HOST + "/" + DB_NAME + "?sslmode=" + DB_SSLMODE

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		slog.Error("failed to connect database:" + err.Error())
	}

	return db
}
