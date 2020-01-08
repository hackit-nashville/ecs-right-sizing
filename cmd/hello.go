package cmd

import (
	lib "hackit-ecs-right-sizing/lib"

	"github.com/spf13/cobra"
)

// Run Command ./pentaho-cli run
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Run: func(cmd *cobra.Command, args []string) {
		lib.EstimateReservation("file-monitor", "prd")
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
}
