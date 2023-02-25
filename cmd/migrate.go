package cmd

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/spf13/cobra"

	"otus-microservice-architecture/app/config"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run migrations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfig()
		if err != nil {
			panic(err)
		}

		migration, err := migrate.New("file://migrations", config.DB.GetDSN())
		if err != nil {
			panic(err)
		}

		migration.LockTimeout = config.Migrate.LockTimeout
		migration.PrefetchMigrations = config.Migrate.PrefetchMigrations

		err = migration.Up()
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				log.Println("migrations up done (no changes)")

				return
			}

			panic(err)
		}

		log.Println("all migrations up done")
	},
}
