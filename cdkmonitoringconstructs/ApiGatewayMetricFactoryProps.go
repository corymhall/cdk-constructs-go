package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Experimental.
type ApiGatewayMetricFactoryProps struct {
	// API to monitor.
	// Experimental.
	Api awsapigateway.IRestApi `field:"required" json:"api" yaml:"api"`
	// On undefined value is not set in dimensions.
	// Experimental.
	ApiMethod *string `field:"optional" json:"apiMethod" yaml:"apiMethod"`
	// On undefined value is not set in dimensions.
	// Experimental.
	ApiResource *string `field:"optional" json:"apiResource" yaml:"apiResource"`
	// Default: - prod.
	//
	// Experimental.
	ApiStage *string `field:"optional" json:"apiStage" yaml:"apiStage"`
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

