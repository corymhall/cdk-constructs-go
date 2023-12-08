package cdkmonitoringconstructs


// Experimental.
type BaseFargateServiceAlarms struct {
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
}

