package main

import (
	_ "github.com/lib/pq"
	user "github.com/sabrina-djebbar/spelling-app-backend/srv/user"
	"github.com/spf13/cobra"
	"log"
	"math/rand"
	"time"
)

// https://www.youtube.com/watch?v=x_N2VjGQKr4
var rootCMD = &cobra.Command{
	Use:          "spelling",
	Short:        "Spelling app is a monobinary for all our go services",
	Long:         "",
	SilenceUsage: true,
	RunE:         func(cmd *cobra.Command, args []string) error { return cmd.Help() },
}

func init() {
	// Seed random number generators for entire program
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Add the config flag
	rootCMD.PersistentFlags().String("config", "", "Config variable json")

	// Load the services
	rootCMD.AddCommand(


		user.CMD,

	)
}

func main() {
	// Listen to the kill command in the background to not cause issues whilst executing commands
	//go killable.ListenToKill()

	// Run command and exit with error code upon error
	if err := rootCMD.Execute(); err != nil {
		log.Fatalw("an error occurred executing that command", "err_message", err.Error(), "err", err)
		return
	}
}
