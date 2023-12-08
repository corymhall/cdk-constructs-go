package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type DynamoTableGlobalSecondaryIndexMetricFactoryProps struct {
	// Experimental.
	GlobalSecondaryIndexName *string `field:"required" json:"globalSecondaryIndexName" yaml:"globalSecondaryIndexName"`
	// Experimental.
	Table awsdynamodb.ITable `field:"required" json:"table" yaml:"table"`
}

