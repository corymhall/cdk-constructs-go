package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsstepfunctions"
)

// Experimental.
type StepFunctionMonitoringProps struct {
	// Experimental.
	StateMachine awsstepfunctions.IStateMachine `field:"required" json:"stateMachine" yaml:"stateMachine"`
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
	AddAbortedExecutionCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addAbortedExecutionCountAlarm" yaml:"addAbortedExecutionCountAlarm"`
	// Experimental.
	AddDurationP50Alarm *map[string]*DurationThreshold `field:"optional" json:"addDurationP50Alarm" yaml:"addDurationP50Alarm"`
	// Experimental.
	AddDurationP90Alarm *map[string]*DurationThreshold `field:"optional" json:"addDurationP90Alarm" yaml:"addDurationP90Alarm"`
	// Experimental.
	AddDurationP99Alarm *map[string]*DurationThreshold `field:"optional" json:"addDurationP99Alarm" yaml:"addDurationP99Alarm"`
	// Experimental.
	AddFailedExecutionCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addFailedExecutionCountAlarm" yaml:"addFailedExecutionCountAlarm"`
	// Experimental.
	AddFailedExecutionRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"addFailedExecutionRateAlarm" yaml:"addFailedExecutionRateAlarm"`
	// Add minimum started execution count alarm for the stepfunctions.
	// Experimental.
	AddMinStartedExecutionCountAlarm *map[string]*MinRunningTaskCountThreshold `field:"optional" json:"addMinStartedExecutionCountAlarm" yaml:"addMinStartedExecutionCountAlarm"`
	// Experimental.
	AddThrottledExecutionCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addThrottledExecutionCountAlarm" yaml:"addThrottledExecutionCountAlarm"`
	// Experimental.
	AddTimedOutExecutionCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addTimedOutExecutionCountAlarm" yaml:"addTimedOutExecutionCountAlarm"`
}

