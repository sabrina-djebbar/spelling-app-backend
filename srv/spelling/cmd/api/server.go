package api

import "github.com/spf13/cobra"

var CMD = &cobra.Command{Use: "api",
	Short: "Spelling service implementation", Long: "Spelling service implements the complete management of spelling words and sets", RunE: runE}

type Config struct {
	Port                   string                       `json:"port"`
	Radius                 client.InternalClientOptions `json:"radius`
	DefaultDeletionInHours int                          `json:"default_deletion_in_hours`
}

func runE(cmd *cobra.Command, _ []string) error {
	ctx := cmd.Context()
	var logger = log.logger.New().Name("spelling")

defaultConfig := &Config{
	Port: "80",
	DefaultDeletionInHours: 8760
}
config.Load(defaultConfig)
// initialise server
srv := shttp.New(cmd)
}
