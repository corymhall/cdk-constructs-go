package cdkmonitoringconstructs


// Experimental.
type KinesisFirehoseMonitoringProps struct {
	// Experimental.
	DeliveryStreamName *string `field:"required" json:"deliveryStreamName" yaml:"deliveryStreamName"`
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
	AddIncomingBytesExceedThresholdAlarm *map[string]*FirehoseStreamLimitThreshold `field:"optional" json:"addIncomingBytesExceedThresholdAlarm" yaml:"addIncomingBytesExceedThresholdAlarm"`
	// Experimental.
	AddIncomingPutRequestsExceedThresholdAlarm *map[string]*FirehoseStreamLimitThreshold `field:"optional" json:"addIncomingPutRequestsExceedThresholdAlarm" yaml:"addIncomingPutRequestsExceedThresholdAlarm"`
	// Experimental.
	AddIncomingRecordsExceedThresholdAlarm *map[string]*FirehoseStreamLimitThreshold `field:"optional" json:"addIncomingRecordsExceedThresholdAlarm" yaml:"addIncomingRecordsExceedThresholdAlarm"`
	// Experimental.
	AddRecordsThrottledAlarm *map[string]*RecordsThrottledThreshold `field:"optional" json:"addRecordsThrottledAlarm" yaml:"addRecordsThrottledAlarm"`
}

