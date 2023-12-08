package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
)

// Experimental.
type AppSyncMetricFactoryProps struct {
	// the GraphQL API to monitor.
	// Experimental.
	Api awsappsync.IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// whether the TPS should be filled with zeroes.
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// method to compute TPS.
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

