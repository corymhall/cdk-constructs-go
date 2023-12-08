package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Experimental.
type S3BucketMetricFactoryProps struct {
	// Experimental.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// Experimental.
	StorageType StorageType `field:"optional" json:"storageType" yaml:"storageType"`
}

