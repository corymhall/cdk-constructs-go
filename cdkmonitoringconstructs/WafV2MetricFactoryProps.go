package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswafv2"
)

// Experimental.
type WafV2MetricFactoryProps struct {
	// Experimental.
	Acl awswafv2.CfnWebACL `field:"required" json:"acl" yaml:"acl"`
	// Required if acl has a "REGIONAL" scope.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

