package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
)

// Experimental.
type AutoScalingGroupMetricFactoryProps struct {
	// Experimental.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `field:"required" json:"autoScalingGroup" yaml:"autoScalingGroup"`
}

