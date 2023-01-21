package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"otus-microservice-architecture/app/config"
	"otus-microservice-architecture/app/server"
)

var apiServer = &cobra.Command{
	Use:   "application",
	Short: "Run api http server",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := config.NewConfig()

		if err != nil {
			log.Fatal(err)
		}

		s := server.New(config)

		if err := s.Start(); err != nil {
			log.Fatal(err)
		}
	},
}
