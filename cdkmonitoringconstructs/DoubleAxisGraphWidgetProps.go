package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type DoubleAxisGraphWidgetProps struct {
	// Experimental.
	Height *float64 `field:"required" json:"height" yaml:"height"`
	// Experimental.
	LeftAxis *awscloudwatch.YAxisProps `field:"required" json:"leftAxis" yaml:"leftAxis"`
	// Experimental.
	LeftMetrics *[]awscloudwatch.IMetric `field:"required" json:"leftMetrics" yaml:"leftMetrics"`
	// Experimental.
	RightAxis *awscloudwatch.YAxisProps `field:"required" json:"rightAxis" yaml:"rightAxis"`
	// Experimental.
	RightMetrics *[]awscloudwatch.IMetric `field:"required" json:"rightMetrics" yaml:"rightMetrics"`
	// Experimental.
	Width *float64 `field:"required" json:"width" yaml:"width"`
	// Experimental.
	LeftAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"optional" json:"leftAnnotations" yaml:"leftAnnotations"`
	// Experimental.
	RightAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"optional" json:"rightAnnotations" yaml:"rightAnnotations"`
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

