package cdkmonitoringconstructs


// Experimental.
type ElastiCacheClusterMonitoringProps struct {
	// Cluster to monitor.
	// Default: - monitor all clusters.
	//
	// Experimental.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
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
	// Cluster type (needed, since each type has their own specific metrics).
	// Experimental.
	ClusterType ElastiCacheClusterType `field:"required" json:"clusterType" yaml:"clusterType"`
	// Add CPU usage alarm (useful for all clusterTypes including Redis).
	// Experimental.
	AddCpuUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addCpuUsageAlarm" yaml:"addCpuUsageAlarm"`
	// Add alarm on number of evicted items.
	// Experimental.
	AddMaxEvictedItemsCountAlarm *map[string]*MaxItemsCountThreshold `field:"optional" json:"addMaxEvictedItemsCountAlarm" yaml:"addMaxEvictedItemsCountAlarm"`
	// Add alarm on total number of items.
	// Experimental.
	AddMaxItemsCountAlarm *map[string]*MaxItemsCountThreshold `field:"optional" json:"addMaxItemsCountAlarm" yaml:"addMaxItemsCountAlarm"`
	// Add alarm on amount of used swap memory.
	// Experimental.
	AddMaxUsedSwapMemoryAlarm *map[string]*MaxUsedSwapMemoryThreshold `field:"optional" json:"addMaxUsedSwapMemoryAlarm" yaml:"addMaxUsedSwapMemoryAlarm"`
	// Add alarm on amount of freeable memory.
	// Experimental.
	AddMinFreeableMemoryAlarm *map[string]*MinFreeableMemoryThreshold `field:"optional" json:"addMinFreeableMemoryAlarm" yaml:"addMinFreeableMemoryAlarm"`
	// Add Redis engine CPU usage alarm.
	//
	// It is recommended to monitor CPU utilization with `addCpuUsageAlarm`
	// as well for hosts with two vCPUs or less.
	// Experimental.
	AddRedisEngineCpuUsageAlarm *map[string]*UsageThreshold `field:"optional" json:"addRedisEngineCpuUsageAlarm" yaml:"addRedisEngineCpuUsageAlarm"`
}

