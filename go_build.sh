#!/usr/bin/env bash

FILES="main.go handler.go cloudwatch.go"

if [[ -f ./go/bin/go ]]
then
    ./go/bin/go build -o hes_metrics_collector $FILES
else
    GOOS=linux go build -o hes_metrics_collector $FILES
fi