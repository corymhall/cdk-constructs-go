package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync"
)

// Experimental.
type AppSyncMonitoringProps struct {
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
	AddLatencyP50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP50Alarm" yaml:"addLatencyP50Alarm"`
	// Experimental.
	AddLatencyP90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP90Alarm" yaml:"addLatencyP90Alarm"`
	// Experimental.
	AddLatencyP99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addLatencyP99Alarm" yaml:"addLatencyP99Alarm"`
	// Experimental.
	AddLowTpsAlarm *map[string]*LowTpsThreshold `field:"optional" json:"addLowTpsAlarm" yaml:"addLowTpsAlarm"`
	// the GraphQL API to monitor.
	// Experimental.
	Api awsappsync.IGraphqlApi `field:"required" json:"api" yaml:"api"`
	// whether the TPS should be filled with zeroes.
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// method to compute TPS.
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

