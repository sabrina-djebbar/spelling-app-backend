package user

import (
	"github.com/sabrina-djebbar/spelling-app-backend/srv/user/cmd"
	"github.com/spf13/cobra"
)

var CMD = &cobra.Command{
	Use:   "user",
	Short: "",
}

func init() {
	CMD.AddCommand(cmd.CMD)

}
