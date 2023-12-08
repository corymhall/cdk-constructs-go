package cdkmonitoringconstructs


// Experimental.
type DynamoTableMonitoringOptions struct {
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
	AddAverageSuccessfulBatchGetItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulBatchGetItemLatencyAlarm" yaml:"addAverageSuccessfulBatchGetItemLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulBatchWriteItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulBatchWriteItemLatencyAlarm" yaml:"addAverageSuccessfulBatchWriteItemLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulDeleteItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulDeleteItemLatencyAlarm" yaml:"addAverageSuccessfulDeleteItemLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulGetItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulGetItemLatencyAlarm" yaml:"addAverageSuccessfulGetItemLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulGetRecordsLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulGetRecordsLatencyAlarm" yaml:"addAverageSuccessfulGetRecordsLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulPutItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulPutItemLatencyAlarm" yaml:"addAverageSuccessfulPutItemLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulQueryLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulQueryLatencyAlarm" yaml:"addAverageSuccessfulQueryLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulScanLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulScanLatencyAlarm" yaml:"addAverageSuccessfulScanLatencyAlarm"`
	// Experimental.
	AddAverageSuccessfulUpdateItemLatencyAlarm *map[string]*LatencyThreshold `field:"optional" json:"addAverageSuccessfulUpdateItemLatencyAlarm" yaml:"addAverageSuccessfulUpdateItemLatencyAlarm"`
	// Experimental.
	AddConsumedReadCapacityAlarm *map[string]*ConsumedCapacityThreshold `field:"optional" json:"addConsumedReadCapacityAlarm" yaml:"addConsumedReadCapacityAlarm"`
	// Experimental.
	AddConsumedWriteCapacityAlarm *map[string]*ConsumedCapacityThreshold `field:"optional" json:"addConsumedWriteCapacityAlarm" yaml:"addConsumedWriteCapacityAlarm"`
	// Experimental.
	AddReadThrottledEventsCountAlarm *map[string]*ThrottledEventsThreshold `field:"optional" json:"addReadThrottledEventsCountAlarm" yaml:"addReadThrottledEventsCountAlarm"`
	// Experimental.
	AddSystemErrorCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addSystemErrorCountAlarm" yaml:"addSystemErrorCountAlarm"`
	// Experimental.
	AddWriteThrottledEventsCountAlarm *map[string]*ThrottledEventsThreshold `field:"optional" json:"addWriteThrottledEventsCountAlarm" yaml:"addWriteThrottledEventsCountAlarm"`
}

