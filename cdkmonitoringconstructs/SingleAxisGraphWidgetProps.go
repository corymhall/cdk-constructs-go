package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type SingleAxisGraphWidgetProps struct {
	// Experimental.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// Experimental.
	LeftAxis *awscloudwatch.YAxisProps `field:"required" json:"leftAxis" yaml:"leftAxis"`
	// Experimental.
	LeftMetrics *[]awscloudwatch.IMetric `field:"required" json:"leftMetrics" yaml:"leftMetrics"`
	// Experimental.
	Width *float64 `field:"required" json:"width" yaml:"width"`
	// Experimental.
	LeftAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"optional" json:"leftAnnotations" yaml:"leftAnnotations"`
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

