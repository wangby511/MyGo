package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"

	"fmt"
)

func main() {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	fmt.Println("Helloworld!")
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)
	fmt.Println("svc", svc)

	_, err := svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueCalls"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(8.0),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("SiteName"),
						Value: aws.String("example.com"),
					},
				},
			},
			//&cloudwatch.MetricDatum{
			//	MetricName: aws.String("UniqueVisitors"),
			//	Unit:       aws.String("Count"),
			//	Value:      aws.Float64(5885.0),
			//	Dimensions: []*cloudwatch.Dimension{
			//		&cloudwatch.Dimension{
			//			Name:  aws.String("SiteName"),
			//			Value: aws.String("example.com"),
			//		},
			//	},
			//},
			//&cloudwatch.MetricDatum{
			//	MetricName: aws.String("UniqueVisits"),
			//	Unit:       aws.String("Count"),
			//	Value:      aws.Float64(8628.0),
			//	Dimensions: []*cloudwatch.Dimension{
			//		&cloudwatch.Dimension{
			//			Name:  aws.String("SiteName"),
			//			Value: aws.String("example.com"),
			//		},
			//	},
			//},
			//&cloudwatch.MetricDatum{
			//	MetricName: aws.String("PageViews"),
			//	Unit:       aws.String("Count"),
			//	Value:      aws.Float64(18057.0),
			//	Dimensions: []*cloudwatch.Dimension{
			//		&cloudwatch.Dimension{
			//			Name:  aws.String("PageURL"),
			//			Value: aws.String("my-page.html"),
			//		},
			//	},
			//},
		},
	})
	if err != nil {
		fmt.Println("Error adding metrics:", err.Error())
		return
	} else {
		fmt.Println("Publishing to CW successfully!")
	}

	//// Get information about metrics
	//result, err := svc.ListMetrics(&cloudwatch.ListMetricsInput{
	//	Namespace: aws.String("Site/Traffic"),
	//})
	//if err != nil {
	//	fmt.Println("Error getting metrics:", err.Error())
	//	return
	//}
	//
	//for _, metric := range result.Metrics {
	//	fmt.Println(*metric.MetricName)
	//
	//	for _, dim := range metric.Dimensions {
	//		fmt.Println(*dim.Name + ":", *dim.Value)
	//		fmt.Println()
	//	}
	//}
}