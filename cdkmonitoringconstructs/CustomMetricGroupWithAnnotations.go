package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type CustomMetricGroupWithAnnotations struct {
	// Experimental.
	Annotations *[]*awscloudwatch.HorizontalAnnotation `field:"required" json:"annotations" yaml:"annotations"`
	// Experimental.
	MetricGroup *CustomMetricGroup `field:"required" json:"metricGroup" yaml:"metricGroup"`
	// Experimental.
	RightAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"required" json:"rightAnnotations" yaml:"rightAnnotations"`
	// Experimental.
	TitleAddons *[]*string `field:"required" json:"titleAddons" yaml:"titleAddons"`
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
}

