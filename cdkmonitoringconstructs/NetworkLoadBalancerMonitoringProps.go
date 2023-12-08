package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Experimental.
type NetworkLoadBalancerMonitoringProps struct {
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
	// Plain name, used in naming alarms.
	//
	// This unique among other resources, and respect the AWS CDK restriction posed on alarm names.
	// The length must be 1 - 255 characters and although the validation rules are undocumented, we recommend using ASCII and hyphens.
	// Default: - derives name from the construct itself.
	//
	// Experimental.
	AlarmFriendlyName *string `field:"optional" json:"alarmFriendlyName" yaml:"alarmFriendlyName"`
	// Human-readable name is a freeform string, used as a caption or description.
	//
	// There are no limitations on what it can be.
	// Default: - use alarmFriendlyName.
	//
	// Experimental.
	HumanReadableName *string `field:"optional" json:"humanReadableName" yaml:"humanReadableName"`
	// If this is defined, the local alarm name prefix used in naming alarms for the construct will be set to this value.
	//
	// The length must be 1 - 255 characters and although the validation rules are undocumented, we recommend using ASCII and hyphens.
	// See: AlarmNamingStrategy for more details on alarm name prefixes.
	//
	// Experimental.
	LocalAlarmNamePrefixOverride *string `field:"optional" json:"localAlarmNamePrefixOverride" yaml:"localAlarmNamePrefixOverride"`
	// Flag indicating if the widgets should be added to alarm dashboard.
	// Default: - true.
	//
	// Experimental.
	AddToAlarmDashboard *bool `field:"optional" json:"addToAlarmDashboard" yaml:"addToAlarmDashboard"`
	// Flag indicating if the widgets should be added to detailed dashboard.
	// Default: - true.
	//
	// Experimental.
	AddToDetailDashboard *bool `field:"optional" json:"addToDetailDashboard" yaml:"addToDetailDashboard"`
	// Flag indicating if the widgets should be added to summary dashboard.
	// Default: - true.
	//
	// Experimental.
	AddToSummaryDashboard *bool `field:"optional" json:"addToSummaryDashboard" yaml:"addToSummaryDashboard"`
	// Calls provided function to process all alarms created.
	// Experimental.
	UseCreatedAlarms IAlarmConsumer `field:"optional" json:"useCreatedAlarms" yaml:"useCreatedAlarms"`
	// Experimental.
	AddHealthyTaskCountAlarm *map[string]*HealthyTaskCountThreshold `field:"optional" json:"addHealthyTaskCountAlarm" yaml:"addHealthyTaskCountAlarm"`
	// Experimental.
	AddHealthyTaskPercentAlarm *map[string]*HealthyTaskPercentThreshold `field:"optional" json:"addHealthyTaskPercentAlarm" yaml:"addHealthyTaskPercentAlarm"`
	// Experimental.
	AddMinProcessedBytesAlarm *map[string]*MinProcessedBytesThreshold `field:"optional" json:"addMinProcessedBytesAlarm" yaml:"addMinProcessedBytesAlarm"`
	// Experimental.
	AddUnhealthyTaskCountAlarm *map[string]*UnhealthyTaskCountThreshold `field:"optional" json:"addUnhealthyTaskCountAlarm" yaml:"addUnhealthyTaskCountAlarm"`
}

