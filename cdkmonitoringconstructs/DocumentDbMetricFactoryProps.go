package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdocdb"
)

// Experimental.
type DocumentDbMetricFactoryProps struct {
	// database cluster.
	// Experimental.
	Cluster awsdocdb.IDatabaseCluster `field:"required" json:"cluster" yaml:"cluster"`
}

