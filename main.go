package main

import (
"github.com/aws/aws-lambda-go/lambda"
"github.com/aws/aws-sdk-go/aws/session"
"github.com/aws/aws-sdk-go/service/cloudwatch"
)


type api struct {
	pms PutMetricService
}



func main() {

	cSess := session.Must(session.NewSession())
	cw := PutMetricService{HeCloudwatch{cloudwatch.New(cSess)}}

	app := &api{cw}

	lambda.Start(app.lambdaHandler)
}