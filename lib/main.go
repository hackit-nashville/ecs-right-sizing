package lib

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

var (
	sess = session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
)

// HelloWorld says hello
func HelloWorld() {
	fmt.Println("Hello world")
}

// EstimateReservation looks at your ECS service's historical memory utilization and recommends a memory reservation
func EstimateReservation(serviceName, clusterName string) (reservation int64) {
	// var ecs = ecs.New(sess)
	var cloudwatchService = cloudwatch.New(sess)

	output, err := cloudwatchService.GetMetricData(&cloudwatch.GetMetricDataInput{
		EndTime: aws.Time(time.Now()),
		// MetricDataQueries []*MetricDataQuery `type:"list" required:"true"`
		// NextToken *string `type:"string"`
		// ScanBy *string `type:"string" enum:"ScanBy"`
		// StartTime *time.Time `type:"timestamp" required:"true"`
	})

	if err != nil {
		log.Println(err)
	}

	log.Println(output)

	return 1
}
