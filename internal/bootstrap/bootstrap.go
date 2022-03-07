package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"yummlog/internal/controller"
	"yummlog/internal/db"
	"yummlog/internal/service"
)

type AppConfig struct {
	DBReadConfig
	DBWriteConfig
	HTTPConfig
}

type DBReadConfig struct {
	Host     string `envconfig:"DB_READ_HOST" default:"localhost"`
	Port     string `envconfig:"DB_READ_PORT" default:"5432"`
	User     string `envconfig:"DB_READ_USER" default:"root"`
	Password string `envconfig:"DB_READ_PASSWORD" default:"root"`
	DB       string `envconfig:"DB_READ_NAME" default:"yummlog"`
	Schema   string `envconfig:"DB_READ_SCHEMA_NAME" default:"yummlog"`
	SSLMode  string `envconfig:"DB_READ_SSL_MODE" default:"disable"`
}

func (dbr DBReadConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=%s search_path=%s", dbr.Host, dbr.Port, dbr.DB, dbr.User, dbr.Password, dbr.SSLMode, dbr.Schema)
}

type DBWriteConfig struct {
	Host     string `envconfig:"DB_WRITE_HOST" default:"localhost"`
	Port     string `envconfig:"DB_WRITE_PORT" default:"5432"`
	User     string `envconfig:"DB_WRITE_USER" default:"root"`
	Password string `envconfig:"DB_WRITE_PASSWORD" default:"root"`
	DB       string `envconfig:"DB_WRITE_NAME" default:"yummlog"`
	Schema   string `envconfig:"DB_WRITE_SCHEMA_NAME" default:"yummlog"`
	SSLMode  string `envconfig:"DB_WRITE_SSL_MODE" default:"disable"`
}

func (dbw DBWriteConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s user=%s password='%s' sslmode=%s search_path=%s", dbw.Host, dbw.Port, dbw.DB, dbw.User, dbw.Password, dbw.SSLMode, dbw.Schema)
}

type HTTPConfig struct {
	Port    int    `envconfig:"HTTP_PORT" default:"3000"`
	BaseURL string `envconfig:"HTTP_BASE_URL" default:"/yummlog"`
}

type Application struct {
	Config         *AppConfig
	RESTController controller.RESTController
}

func NewApplication(ctx context.Context) (*Application, error) {
	cfg := AppConfig{}

	fps, err := newFoodPostsCRUD(ctx, &cfg)
	if err != nil {
		return nil, err
	}

	rc := controller.RESTController{
		FoodPostsService: fps,
	}

	return &Application{
		Config:         &cfg,
		RESTController: rc,
	}, nil
}

func newFoodPostsCRUD(ctx context.Context, cfg *AppConfig) (service.FoodPostsCRUD, error) {
	read, err := sql.Open("postgres", cfg.DBReadConfig.ConnectionString())
	if err != nil {
		return nil, err
	}
	readQuery := db.New(read)

	write, err := sql.Open("postgres", cfg.DBWriteConfig.ConnectionString())
	if err != nil {
		return nil, err
	}
	writeQuery := db.New(write)

	return service.NewFoodPostsCRUD(ctx, readQuery, writeQuery), nil
}
