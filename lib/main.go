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

	input := &cloudwatch.GetMetricDataInput{
		EndTime: aws.Time(time.Now()),
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			&cloudwatch.MetricDataQuery{
				Id: aws.String("taskcount"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{

						// The name of the metric.
						MetricName: aws.String("RunningTaskCount"),

						// The namespace of the metric.
						Namespace: aws.String("ECS/ContainerInsights"),

						// The dimensions for the metric.
						Dimensions: []*cloudwatch.Dimension{
							&cloudwatch.Dimension{
								Name:  aws.String("ServiceName"),
								Value: &serviceName,
							},
							&cloudwatch.Dimension{
								Name:  aws.String("ClusterName"),
								Value: &clusterName,
							},
						},
					},
					Period: aws.Int64(3600),
					Stat:   aws.String("Maximum"),
					Unit:   aws.String("Count"),
				},
			},
			&cloudwatch.MetricDataQuery{
				Id: aws.String("reserved"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						// The name of the metric.

						// The namespace of the metric.
						Namespace:  aws.String("ECS/ContainerInsights"),
						MetricName: aws.String("MemoryReserved"),
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
		StartTime: aws.Time(time.Now().AddDate(0, 0, -14)),
	}

	for input.NextToken == nil {
		output, err := cloudwatchService.GetMetricData(input)

		if err != nil {
			log.Println(err)
		}

		log.Println(output)

		input.NextToken = output.NextToken
	}

	return 1
}
