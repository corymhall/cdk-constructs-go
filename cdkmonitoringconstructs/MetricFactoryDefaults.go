package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// These are the globals used for each metric, unless there is some kind of override.
// Experimental.
type MetricFactoryDefaults struct {
	// Account where the metrics exist.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Cross-Account-Cross-Region.html
	//
	// Default: The account configured by the construct holding the Monitoring construct.
	//
	// Experimental.
	Account *string `field:"optional" json:"account" yaml:"account"`
	// Each metric exists in a namespace.
	//
	// AWS Services have their own namespace, but here you can specify your custom one.
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Metric period.
	//
	// Default value is used if not defined.
	// Default: - DefaultMetricPeriod.
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// Region where the metrics exist.
	// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Cross-Account-Cross-Region.html
	//
	// Default: The region configured by the construct holding the Monitoring construct.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

