package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsautoscaling"
)

// Experimental.
type EC2MetricFactoryProps struct {
	// Auto-Scaling Group to monitor.
	// Default: - no Auto-Scaling Group filter.
	//
	// Experimental.
	AutoScalingGroup awsautoscaling.IAutoScalingGroup `field:"optional" json:"autoScalingGroup" yaml:"autoScalingGroup"`
	// Selected IDs of EC2 instances to monitor.
	// Default: - no instance filter.
	//
	// Experimental.
	InstanceIds *[]*string `field:"optional" json:"instanceIds" yaml:"instanceIds"`
}

