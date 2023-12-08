package cdkmonitoringconstructs


// Experimental.
type LambdaFunctionMonitoringOptions struct {
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
	AddConcurrentExecutionsCountAlarm *map[string]*RunningTaskCountThreshold `field:"optional" json:"addConcurrentExecutionsCountAlarm" yaml:"addConcurrentExecutionsCountAlarm"`
	// Experimental.
	AddEnhancedMonitoringAvgCpuTotalTimeAlarm *map[string]*DurationThreshold `field:"optional" json:"addEnhancedMonitoringAvgCpuTotalTimeAlarm" yaml:"addEnhancedMonitoringAvgCpuTotalTimeAlarm"`
	// Experimental.
	AddEnhancedMonitoringAvgMemoryUtilizationAlarm *map[string]*UsageThreshold `field:"optional" json:"addEnhancedMonitoringAvgMemoryUtilizationAlarm" yaml:"addEnhancedMonitoringAvgMemoryUtilizationAlarm"`
	// Experimental.
	AddEnhancedMonitoringMaxCpuTotalTimeAlarm *map[string]*DurationThreshold `field:"optional" json:"addEnhancedMonitoringMaxCpuTotalTimeAlarm" yaml:"addEnhancedMonitoringMaxCpuTotalTimeAlarm"`
	// Experimental.
	AddEnhancedMonitoringMaxMemoryUtilizationAlarm *map[string]*UsageThreshold `field:"optional" json:"addEnhancedMonitoringMaxMemoryUtilizationAlarm" yaml:"addEnhancedMonitoringMaxMemoryUtilizationAlarm"`
	// Experimental.
	AddEnhancedMonitoringP90CpuTotalTimeAlarm *map[string]*DurationThreshold `field:"optional" json:"addEnhancedMonitoringP90CpuTotalTimeAlarm" yaml:"addEnhancedMonitoringP90CpuTotalTimeAlarm"`
	// Experimental.
	AddEnhancedMonitoringP90MemoryUtilizationAlarm *map[string]*UsageThreshold `field:"optional" json:"addEnhancedMonitoringP90MemoryUtilizationAlarm" yaml:"addEnhancedMonitoringP90MemoryUtilizationAlarm"`
	// Experimental.
	AddFaultCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addFaultCountAlarm" yaml:"addFaultCountAlarm"`
	// Experimental.
	AddFaultRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"addFaultRateAlarm" yaml:"addFaultRateAlarm"`
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
	// Experimental.
	AddMaxIteratorAgeAlarm *map[string]*MaxAgeThreshold `field:"optional" json:"addMaxIteratorAgeAlarm" yaml:"addMaxIteratorAgeAlarm"`
	// Experimental.
	AddMinInvocationsCountAlarm *map[string]*MinUsageCountThreshold `field:"optional" json:"addMinInvocationsCountAlarm" yaml:"addMinInvocationsCountAlarm"`
	// Experimental.
	AddProvisionedConcurrencySpilloverInvocationsCountAlarm *map[string]*RunningTaskCountThreshold `field:"optional" json:"addProvisionedConcurrencySpilloverInvocationsCountAlarm" yaml:"addProvisionedConcurrencySpilloverInvocationsCountAlarm"`
	// Experimental.
	AddProvisionedConcurrencySpilloverInvocationsRateAlarm *map[string]*RunningTaskRateThreshold `field:"optional" json:"addProvisionedConcurrencySpilloverInvocationsRateAlarm" yaml:"addProvisionedConcurrencySpilloverInvocationsRateAlarm"`
	// Experimental.
	AddThrottlesCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"addThrottlesCountAlarm" yaml:"addThrottlesCountAlarm"`
	// Experimental.
	AddThrottlesRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"addThrottlesRateAlarm" yaml:"addThrottlesRateAlarm"`
}

