package lib

import "fmt"

// HelloWorld says hello
func HelloWorld() {
	fmt.Println("Hello world")
}

// EstimateReservation looks at your ECS service's historical memory utilization and recommends a memory reservation
func EstimateReservation(serviceName, clusterName string) (reservation int64) {
	return 1
}
