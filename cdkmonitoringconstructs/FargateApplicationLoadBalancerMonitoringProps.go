package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awselasticloadbalancingv2"
)

// Monitoring props for Fargate service with application load balancer and plain service.
// Experimental.
type FargateApplicationLoadBalancerMonitoringProps struct {
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
	ApplicationLoadBalancer awselasticloadbalancingv2.IApplicationLoadBalancer `field:"required" json:"applicationLoadBalancer" yaml:"applicationLoadBalancer"`
	// Experimental.
	ApplicationTargetGroup awselasticloadbalancingv2.IApplicationTargetGroup `field:"required" json:"applicationTargetGroup" yaml:"applicationTargetGroup"`
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
	AddCpuUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addCpuUsageAlarm" yaml:"addCpuUsageAlarm"`
	// Experimental.
	AddMemoryUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addMemoryUsageAlarm" yaml:"addMemoryUsageAlarm"`
	// Container Insights needs to be enabled for the cluster for this alarm.
	// Experimental.
	AddRunningTaskCountAlarm *map[string]*RunningTaskCountThreshold `field:"optional" json:"addRunningTaskCountAlarm" yaml:"addRunningTaskCountAlarm"`
	// maximum number of tasks, as specified in your auto scaling config.
	// Experimental.
	MaxAutoScalingTaskCount *float64 `field:"optional" json:"maxAutoScalingTaskCount" yaml:"maxAutoScalingTaskCount"`
	// minimum number of tasks, as specified in your auto scaling config.
	// Experimental.
	MinAutoScalingTaskCount *float64 `field:"optional" json:"minAutoScalingTaskCount" yaml:"minAutoScalingTaskCount"`
	// Experimental.
	FargateService awsecs.FargateService `field:"required" json:"fargateService" yaml:"fargateService"`
	// Experimental.
	AddHealthyTaskCountAlarm *map[string]*HealthyTaskCountThreshold `field:"optional" json:"addHealthyTaskCountAlarm" yaml:"addHealthyTaskCountAlarm"`
	// Experimental.
	AddHealthyTaskPercentAlarm *map[string]*HealthyTaskPercentThreshold `field:"optional" json:"addHealthyTaskPercentAlarm" yaml:"addHealthyTaskPercentAlarm"`
	// Experimental.
	AddMinProcessedBytesAlarm *map[string]*MinProcessedBytesThreshold `field:"optional" json:"addMinProcessedBytesAlarm" yaml:"addMinProcessedBytesAlarm"`
	// Experimental.
	AddUnhealthyTaskCountAlarm *map[string]*UnhealthyTaskCountThreshold `field:"optional" json:"addUnhealthyTaskCountAlarm" yaml:"addUnhealthyTaskCountAlarm"`
	// Invert the statistics of `HealthyHostCount` and `UnHealthyHostCount`.
	//
	// When `invertLoadBalancerTaskCountMetricsStatistics` is set to false, the minimum of `HealthyHostCount` and the maximum of `UnHealthyHostCount` are monitored.
	// When `invertLoadBalancerTaskCountMetricsStatistics` is set to true, the maximum of `HealthyHostCount` and the minimum of `UnHealthyHostCount` are monitored.
	//
	// `invertLoadBalancerTaskCountMetricsStatistics` is recommended to set to true as per the guidelines at
	// https://docs.aws.amazon.com/elasticloadbalancing/latest/network/load-balancer-cloudwatch-metrics.html#metric-statistics
	// Default: false.
	//
	// Experimental.
	InvertLoadBalancerTaskCountMetricsStatistics *bool `field:"optional" json:"invertLoadBalancerTaskCountMetricsStatistics" yaml:"invertLoadBalancerTaskCountMetricsStatistics"`
}

