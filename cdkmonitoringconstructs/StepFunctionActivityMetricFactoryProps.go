package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Experimental.
type StepFunctionActivityMetricFactoryProps struct {
	// Experimental.
	Activity awsstepfunctions.IActivity `field:"required" json:"activity" yaml:"activity"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

