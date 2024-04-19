package cmd

var CMD = &cobra.Command{
	Use:   "user",
	Short: "",
}

func init() {
	CMD.AddCommand(api.CMD)
	CMD.AddCommand(deleteoldusersbyuserid.CMD)
}