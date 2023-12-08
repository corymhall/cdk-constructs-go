package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
)

// Experimental.
type ApiGatewayMonitoringProps struct {
	// API to monitor.
	// Experimental.
	Api awsapigateway.IRestApi `field:"required" json:"api" yaml:"api"`
	// On undefined value is not set in dimensions.
	// Experimental.
	ApiMethod *string `field:"optional" json:"apiMethod" yaml:"apiMethod"`
	// On undefined value is not set in dimensions.
	// Experimental.
	ApiResource *string `field:"optional" json:"apiResource" yaml:"apiResource"`
	// Default: - prod.
	//
	// Experimental.
	ApiStage *string `field:"optional" json:"apiStage" yaml:"apiStage"`
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
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
	Add4XXErrorCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"add4XXErrorCountAlarm" yaml:"add4XXErrorCountAlarm"`
	// Experimental.
	Add4XXErrorRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"add4XXErrorRateAlarm" yaml:"add4XXErrorRateAlarm"`
	// Experimental.
	Add5XXFaultCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"add5XXFaultCountAlarm" yaml:"add5XXFaultCountAlarm"`
	// Experimental.
	Add5XXFaultRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"add5XXFaultRateAlarm" yaml:"add5XXFaultRateAlarm"`
	// Experimental.
	AddHighTpsAlarm *map[string]*HighTpsThreshold `field:"optional" json:"addHighTpsAlarm" yaml:"addHighTpsAlarm"`
	// Experimental.
	AddLatencyAverageAlarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyAverageAlarm" yaml:"addLatencyAverageAlarm"`
	// Experimental.
	AddLatencyP100Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP100Alarm" yaml:"addLatencyP100Alarm"`
	// Experimental.
	AddLatencyP50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP50Alarm" yaml:"addLatencyP50Alarm"`
	// Experimental.
	AddLatencyP70Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP70Alarm" yaml:"addLatencyP70Alarm"`
	// Experimental.
	AddLatencyP90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP90Alarm" yaml:"addLatencyP90Alarm"`
	// Experimental.
	AddLatencyP95Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP95Alarm" yaml:"addLatencyP95Alarm"`
	// Experimental.
	AddLatencyP9999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP9999Alarm" yaml:"addLatencyP9999Alarm"`
	// Experimental.
	AddLatencyP999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP999Alarm" yaml:"addLatencyP999Alarm"`
	// Experimental.
	AddLatencyP99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP99Alarm" yaml:"addLatencyP99Alarm"`
	// Experimental.
	AddLatencyTM50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM50Alarm" yaml:"addLatencyTM50Alarm"`
	// Experimental.
	AddLatencyTM70Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM70Alarm" yaml:"addLatencyTM70Alarm"`
	// Experimental.
	AddLatencyTM90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM90Alarm" yaml:"addLatencyTM90Alarm"`
	// Experimental.
	AddLatencyTM95Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM95Alarm" yaml:"addLatencyTM95Alarm"`
	// Experimental.
	AddLatencyTM95OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM95OutlierAlarm" yaml:"addLatencyTM95OutlierAlarm"`
	// Experimental.
	AddLatencyTM9999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM9999Alarm" yaml:"addLatencyTM9999Alarm"`
	// Experimental.
	AddLatencyTM9999OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM9999OutlierAlarm" yaml:"addLatencyTM9999OutlierAlarm"`
	// Experimental.
	AddLatencyTM999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM999Alarm" yaml:"addLatencyTM999Alarm"`
	// Experimental.
	AddLatencyTM999OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM999OutlierAlarm" yaml:"addLatencyTM999OutlierAlarm"`
	// Experimental.
	AddLatencyTM99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM99Alarm" yaml:"addLatencyTM99Alarm"`
	// Experimental.
	AddLatencyTM99OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyTM99OutlierAlarm" yaml:"addLatencyTM99OutlierAlarm"`
	// Experimental.
	AddLowTpsAlarm *map[string]*LowTpsThreshold `field:"optional" json:"addLowTpsAlarm" yaml:"addLowTpsAlarm"`
	// You can specify what latency types you want to be rendered in the dashboards.
	//
	// Note: any latency type with an alarm will be also added automatically.
	// If the list is undefined, default values will be shown.
	// If the list is empty, only the latency types with an alarm will be shown (if any).
	// See: DefaultLatencyTypesToRender).
	//
	// Default: - p50, p90, p99 (.
	//
	// Experimental.
	LatencyTypesToRender *[]LatencyType `field:"optional" json:"latencyTypesToRender" yaml:"latencyTypesToRender"`
}

