package main

import (
"github.com/aws/aws-sdk-go/service/cloudwatch"
"github.com/aws/aws-lambda-go/events"
"encoding/json"
"github.com/aws/aws-sdk-go/aws"
)


const heStatusCodeMetrics = "HES_Metrics"

type HeCloudwatch struct {
	sess *cloudwatch.CloudWatch
}
type MetricClient interface {
	writeMetric(events.KinesisRecord)
}

type PutMetricService struct {
	mc MetricClient
}

type MetricRecord struct {
	Metric_Name string `json:Metric_name`
	Provider_Id int    `json:Provider_id`
	Dimensions string  `json:Dimensions`
	Instance_Count int `json:Instance_count`
}


func (cw HeCloudwatch) writeMetric(record events.KinesisRecord) {
	var metricRecord MetricRecord
	err := json.Unmarshal(record.Data, &metricRecord)
	if err != nil {
		println(err)
	}

	var dim map[string]string

	if metricRecord.Dimensions != "" {
		er := json.Unmarshal([]byte(metricRecord.Dimensions), &dim)
		if er != nil {
			println("error unmarshalling dimensions map: ", er)
		}
	}

	dimensions := getDimensions(dim)
	metricDatum := getMetricDataInput("", metricRecord.Metric_Name, dimensions)

	metricData := &cloudwatch.PutMetricDataInput{
		Namespace: aws.String(heStatusCodeMetrics),
		MetricData: metricDatum,
	}


	pmo, e := cw.sess.PutMetricData(metricData)

	if e != nil {
		println("error writing cloudwatch: ", e)
	} else {
		println("putting metric: ", pmo)
	}

}

func getDimensions(dimensionMap map[string]string) []*cloudwatch.Dimension{
	dimensions := []*cloudwatch.Dimension(nil)

	if len(dimensionMap) != 0 {
		for k, v := range dimensionMap {
			d := new(cloudwatch.Dimension)
			d.Name = aws.String(k)
			d.Value = aws.String(v)

			dimensions = append(dimensions, d)
		}
	}

	return dimensions
}

func getMetricDataInput(metricUnit string, metricName string, dimensions []*cloudwatch.Dimension) []*cloudwatch.MetricDatum {
	var metrics []*cloudwatch.MetricDatum
	metricDatum := &cloudwatch.MetricDatum{
		Value:      aws.Float64(1),
		Unit:       aws.String(metricUnit),
		MetricName: aws.String(metricName),
		Dimensions: dimensions,
	}

	metrics = append(metrics, metricDatum)

	return metrics
}

