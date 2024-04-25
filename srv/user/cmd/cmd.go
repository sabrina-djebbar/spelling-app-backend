package cmd

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/cmd/api"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:   "user",
	Short: "",
}

func init() {
	CMD.AddCommand(api.CMD)
}
