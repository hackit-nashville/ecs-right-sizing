# Hackit ECS Right-Sizing

## Getting Started

You need to turn on [CloudWatch Container Insights](https://console.aws.amazon.com/ecs/home?region=us-east-1#/settings)
## Data

See data.json

    aws cloudwatch get-metric-data --cli-input-json file://metrics.json > data.json

