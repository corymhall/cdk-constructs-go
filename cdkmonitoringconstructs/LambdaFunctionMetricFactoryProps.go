package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Experimental.
type LambdaFunctionMetricFactoryProps struct {
	// Experimental.
	LambdaFunction awslambda.IFunction `field:"required" json:"lambdaFunction" yaml:"lambdaFunction"`
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// Generate dashboard charts for Lambda Insights metrics.
	//
	// To enable Lambda Insights on your Lambda function, see
	// https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Lambda-Insights-Getting-Started-clouddevelopmentkit.html
	// Default: - false.
	//
	// Experimental.
	LambdaInsightsEnabled *bool `field:"optional" json:"lambdaInsightsEnabled" yaml:"lambdaInsightsEnabled"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

