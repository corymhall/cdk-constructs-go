package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics"
)

// Experimental.
type SyntheticsCanaryMetricFactoryProps struct {
	// CloudWatch Canary to monitor.
	// Experimental.
	Canary awssynthetics.Canary `field:"required" json:"canary" yaml:"canary"`
	// Method used to calculate relative rates.
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

