package cmd

import (
	"fmt"
	lib "hackit-ecs-right-sizing/lib"

	"github.com/spf13/cobra"
)

var (
	clusterName, serviceName string
	clean                    bool
)

// Run Command ./pentaho-cli run
var recommendCmd = &cobra.Command{
	Use:   "recommend-memory-reservation",
	Short: "Recommend accepts a service name and cluster then produces a recommended memory reservation",
	Run: func(cmd *cobra.Command, args []string) {
		reservation := lib.EstimateReservation(serviceName, clusterName)
		if clean {
			fmt.Println(reservation)
		} else {
			fmt.Printf("Service: %s, Cluster: %s, Recommended Memory Reservation: %d\n", serviceName, clusterName, reservation)
		}
	},
}

func init() {
	rootCmd.AddCommand(recommendCmd)
	recommendCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "service name")
	recommendCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
	recommendCmd.PersistentFlags().BoolVar(&clean, "clean", false, "print only the recommended reservation")
}
