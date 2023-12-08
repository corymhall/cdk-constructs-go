package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// Experimental.
type CertificateManagerMetricFactoryProps struct {
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"required" json:"certificate" yaml:"certificate"`
}

