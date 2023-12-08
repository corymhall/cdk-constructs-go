package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
)

// Experimental.
type CloudFrontDistributionMetricFactoryProps struct {
	// Experimental.
	Distribution awscloudfront.IDistribution `field:"required" json:"distribution" yaml:"distribution"`
	// Generate dashboard charts for additional CloudFront distribution metrics.
	//
	// To enable additional metrics on your CloudFront distribution, see
	// https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewing-cloudfront-metrics.html#monitoring-console.distributions-additional
	// Default: - true.
	//
	// Experimental.
	AdditionalMetricsEnabled *bool `field:"optional" json:"additionalMetricsEnabled" yaml:"additionalMetricsEnabled"`
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

