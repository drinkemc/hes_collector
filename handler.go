package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func (app *api) lambdaHandler(event events.KinesisEvent) {
	records := event.Records

	for _, r := range records {
		app.pms.mc.writeMetric(r.Kinesis)
	}

}

