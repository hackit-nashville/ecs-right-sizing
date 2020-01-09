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

// EstimateReservation looks at your ECS service's historical memory utilization and recommends a memory reservation
func EstimateReservation(serviceName, clusterName string) (reservation int64) {
	// var ecs = ecs.New(sess)
	var cloudwatchService = cloudwatch.New(sess)

	input := &cloudwatch.GetMetricDataInput{
		EndTime: aws.Time(time.Now()),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			&cloudwatch.MetricDataQuery{
				Id: aws.String("utilized"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						// The name of the metric.

						// The namespace of the metric.
						Namespace:  aws.String("ECS/ContainerInsights"),
						MetricName: aws.String("MemoryUtilized"),
						// The dimensions for the metric.
						Dimensions: []*cloudwatch.Dimension{
							&cloudwatch.Dimension{
								Name:  aws.String("TaskDefinitionFamily"),
								Value: aws.String(clusterName + "-" + serviceName),
							},
							&cloudwatch.Dimension{
								Name:  aws.String("ClusterName"),
								Value: &clusterName,
							},
						},
					},
					Period: aws.Int64(3600),
					Stat:   aws.String("Maximum"),
					Unit:   aws.String("Megabytes"),
				},
			},
		},
		StartTime: aws.Time(time.Now().AddDate(0, 0, -90)),
	}

	results := []memoryUtilizationPoint{}

	for {
		output, err := cloudwatchService.GetMetricData(input)

		if err != nil {
			log.Println(err)
		}

		for _, metricDataResult := range output.MetricDataResults {
			if *metricDataResult.Id == "utilized" {
				for i, value := range metricDataResult.Values {
					results = append(results, memoryUtilizationPoint{metricDataResult.Timestamps[i], value})
				}
				// results = append(results, {metricDataResult.Values...)

			}
		}

		input.NextToken = output.NextToken
		if input.NextToken == nil {
			break
		}
	}

	for _, mu := range results {
		fmt.Println(mu.toString())
	}

	return 0.0
}
