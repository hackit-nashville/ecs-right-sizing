package cmd

import (
	"fmt"

	lib "github.com/hackit-nashville/ecs-right-sizing/lib"

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
		if len(serviceName) > 0 && len(clusterName) > 0 {
			reservation := lib.EstimateReservation(serviceName, clusterName)
			if clean {
				fmt.Println(reservation)
			} else {
				printRecommendation(serviceName, clusterName, reservation)
			}
		} else if len(clusterName) > 0 {
			for _, serviceName := range lib.GetServices(clusterName) {
				reservation := lib.EstimateReservation(serviceName, clusterName)
				printRecommendation(serviceName, clusterName, reservation)
			}
		} else {
			fmt.Println("missing required arg '--cluster'")
		}
	},
}

func printRecommendation(serviceName, clusterName string, reservation int64) {
	fmt.Printf("Service: %s, Cluster: %s, Recommended Memory Reservation: %d\n", serviceName, clusterName, reservation)
}

func init() {
	rootCmd.AddCommand(recommendCmd)
	recommendCmd.PersistentFlags().StringVarP(&serviceName, "service", "s", "", "service name")
	recommendCmd.PersistentFlags().StringVarP(&clusterName, "cluster", "c", "", "cluster name")
	recommendCmd.PersistentFlags().BoolVar(&clean, "clean", false, "print only the recommended reservation")
}
