#!/bin/bash

data=$1
interval=$2
echo $data
echo $interval

#aws cloudwatch get-metric-data --cli-input-json file://metrics.json > data.json

# python3 memory_utilization.calculate(data, interval)
python3 memory_utilization.py $data $interval