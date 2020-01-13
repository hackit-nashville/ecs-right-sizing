# Hackit ECS Right-Sizing
A simple package to give memory reservation recommendations for ECS Containers based on historical usage.

## Getting Started

Turn on [CloudWatch Container Insights](https://console.aws.amazon.com/ecs/home?region=us-east-1#/settings).
You can enable it with the command below. 
**Note:** The metrics from Container Insights will only begin tracking after you enable them. Historical data will not be present and may skew recomendations until Container Insights has been enable for a few weeks or months.
    # Turn on "Cloudwatch Container Insights"
    aws ecs update-cluster-settings \
        --cluster default \    
        --settings name=containerInsights,value=enabled

## Usage
[![asciicast](https://asciinema.org/a/293397.svg)](https://asciinema.org/a/293397)

## RoadMap

- Add verbose flag for a detailed report for each service in a cluster
- Accept a parameter to recommend by weekday or daily
