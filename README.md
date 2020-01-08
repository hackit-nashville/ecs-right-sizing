# Hackit ECS Right-Sizing

## Getting Started

You need to turn on [CloudWatch Container Insights](https://console.aws.amazon.com/ecs/home?region=us-east-1#/settings).

    aws ecs update-cluster-settings \
        --cluster ops \    
        --settings name=containerInsights,value=enabled

## Data

See data.json

    aws cloudwatch get-metric-data --cli-input-json file://metrics.json > data.json

## start.sh

sh start.sh data interval
Starts the python script. Prints values to console and creates a .csv file in relative directory.

## RoadMap

- Update `metricDataResultsToDailyView` to return recommendation: (MAX of []hours / .8)
- Accept a parameter to recommend by weekday or daily
- Write an option to iterate over all services in ECS and make a recommendation/report for each service
