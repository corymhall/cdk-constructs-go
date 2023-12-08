package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Experimental.
type SnsTopicMetricFactoryProps struct {
	// Experimental.
	Topic awssns.ITopic `field:"required" json:"topic" yaml:"topic"`
}

