package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
)

// Props to create BaseServiceMetricFactory.
// Experimental.
type BaseServiceMetricFactoryProps struct {
	// Experimental.
	Service awsecs.BaseService `field:"required" json:"service" yaml:"service"`
}

