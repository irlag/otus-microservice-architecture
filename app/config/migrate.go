package config

import "time"

type Migrate struct {
	LockTimeout        time.Duration `envconfig:"default=15s"`
	PrefetchMigrations uint          `envconfig:"default=10"`
}
