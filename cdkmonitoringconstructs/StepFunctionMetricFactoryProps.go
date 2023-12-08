package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Experimental.
type StepFunctionMetricFactoryProps struct {
	// Experimental.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

