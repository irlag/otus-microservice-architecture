package config

import (
	"fmt"
	"time"
)

type DBConfig struct {
	User            string
	Password        string
	Host            string
	Port            string
	Database        string
	MaxIdleConns    int           `envconfig:"DB_MAX_IDLE_CONNS" default:"2"`
	MaxIdleConnTime time.Duration `envconfig:"DB_MAX_IDLE_CONN_TIME" default:"5m"`
	MaxConns        int           `envconfig:"DB_MAX_CONNS" default:"20"`
	ConnMaxLifetime time.Duration `envconfig:"DB_CONN_MAX_LIFETIME" default:"10m"`
	Driver          string        `envconfig:"DB_DRIVER"`
}

func (db *DBConfig) GetDSN() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable&binary_parameters=yes",
		db.Driver,
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Database,
	)
}
