package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type DynamoTableMetricFactoryProps struct {
	// table to monitor.
	// Experimental.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
	// table billing mode.
	// Default: - best effort auto-detection or PROVISIONED as a fallback.
	//
	// Experimental.
	BillingMode awsdynamodb.BillingMode `field:"optional" json:"billingMode" yaml:"billingMode"`
}

