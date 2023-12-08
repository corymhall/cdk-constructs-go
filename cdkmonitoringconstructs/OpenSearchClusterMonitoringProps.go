package cdkmonitoringconstructs


// Experimental.
type OpenSearchClusterMonitoringProps struct {
	// Experimental.
	Domain interface{} `field:"required" json:"domain" yaml:"domain"`
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
	AddClusterAutomatedSnapshotFailureAlarm *map[string]*OpenSearchClusterAutomatedSnapshotFailureThreshold `field:"optional" json:"addClusterAutomatedSnapshotFailureAlarm" yaml:"addClusterAutomatedSnapshotFailureAlarm"`
	// Experimental.
	AddClusterIndexWritesBlockedAlarm *map[string]*OpenSearchClusterIndexWritesBlockedThreshold `field:"optional" json:"addClusterIndexWritesBlockedAlarm" yaml:"addClusterIndexWritesBlockedAlarm"`
	// Experimental.
	AddClusterNodeCountAlarm *map[string]*OpenSearchClusterNodesThreshold `field:"optional" json:"addClusterNodeCountAlarm" yaml:"addClusterNodeCountAlarm"`
	// Experimental.
	AddClusterStatusAlarm *map[string]*OpenSearchClusterStatusCustomization `field:"optional" json:"addClusterStatusAlarm" yaml:"addClusterStatusAlarm"`
	// Experimental.
	AddCpuSpaceUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addCpuSpaceUsageAlarm" yaml:"addCpuSpaceUsageAlarm"`
	// Experimental.
	AddDiskSpaceUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addDiskSpaceUsageAlarm" yaml:"addDiskSpaceUsageAlarm"`
	// Experimental.
	AddIndexingLatencyP50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIndexingLatencyP50Alarm" yaml:"addIndexingLatencyP50Alarm"`
	// Experimental.
	AddIndexingLatencyP90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIndexingLatencyP90Alarm" yaml:"addIndexingLatencyP90Alarm"`
	// Experimental.
	AddIndexingLatencyP99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIndexingLatencyP99Alarm" yaml:"addIndexingLatencyP99Alarm"`
	// Experimental.
	AddJvmMemoryPressureAlarm *map[string]*UsageThreshold `field:"optional" json:"addJvmMemoryPressureAlarm" yaml:"addJvmMemoryPressureAlarm"`
	// Experimental.
	AddKmsKeyErrorAlarm *map[string]*OpenSearchKmsKeyErrorThreshold `field:"optional" json:"addKmsKeyErrorAlarm" yaml:"addKmsKeyErrorAlarm"`
	// Experimental.
	AddKmsKeyInaccessibleAlarm *map[string]*OpenSearchKmsKeyInaccessibleThreshold `field:"optional" json:"addKmsKeyInaccessibleAlarm" yaml:"addKmsKeyInaccessibleAlarm"`
	// Experimental.
	AddMasterCpuSpaceUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addMasterCpuSpaceUsageAlarm" yaml:"addMasterCpuSpaceUsageAlarm"`
	// Experimental.
	AddMasterJvmMemoryPressureAlarm *map[string]*UsageThreshold `field:"optional" json:"addMasterJvmMemoryPressureAlarm" yaml:"addMasterJvmMemoryPressureAlarm"`
	// Experimental.
	AddSearchLatencyP50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addSearchLatencyP50Alarm" yaml:"addSearchLatencyP50Alarm"`
	// Experimental.
	AddSearchLatencyP90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addSearchLatencyP90Alarm" yaml:"addSearchLatencyP90Alarm"`
	// Experimental.
	AddSearchLatencyP99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addSearchLatencyP99Alarm" yaml:"addSearchLatencyP99Alarm"`
}

