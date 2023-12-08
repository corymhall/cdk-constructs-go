package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Props to create NetworkLoadBalancerMetricFactory.
// Experimental.
type NetworkLoadBalancerMetricFactoryProps struct {
	// Invert the statistics of `HealthyHostCount` and `UnHealthyHostCount`.
	//
	// When `invertStatisticsOfTaskCountEnabled` is set to false, the minimum of `HealthyHostCount` and the maximum of `UnHealthyHostCount` are monitored.
	// When `invertStatisticsOfTaskCountEnabled` is set to true, the maximum of `HealthyHostCount` and the minimum of `UnHealthyHostCount` are monitored.
	//
	// `invertStatisticsOfTaskCountEnabled` is recommended to set to true as per the guidelines at
	// https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-cloudwatch-metrics.html#metric-statistics
	// Default: false.
	//
	// Experimental.
	InvertStatisticsOfTaskCountEnabled *bool `field:"optional" json:"invertStatisticsOfTaskCountEnabled" yaml:"invertStatisticsOfTaskCountEnabled"`
	// Experimental.
	NetworkLoadBalancer awselasticloadbalancingv2.INetworkLoadBalancer `field:"required" json:"networkLoadBalancer" yaml:"networkLoadBalancer"`
	// Experimental.
	NetworkTargetGroup awselasticloadbalancingv2.INetworkTargetGroup `field:"required" json:"networkTargetGroup" yaml:"networkTargetGroup"`
}

