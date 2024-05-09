package cmd

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/spelling/cmd/api"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:   "spelling",
	Short: "",
}

func init() {
	CMD.AddCommand(api.CMD)
}
