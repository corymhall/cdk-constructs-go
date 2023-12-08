package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Experimental.
type SqsQueueMetricFactoryProps struct {
	// Experimental.
	Queue awssqs.IQueue `field:"required" json:"queue" yaml:"queue"`
}

