package cdkmonitoringconstructs


// Experimental.
type ApiGatewayV2MonitoringOptions struct {
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
	Add4xxCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"add4xxCountAlarm" yaml:"add4xxCountAlarm"`
	// Experimental.
	Add4xxRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"add4xxRateAlarm" yaml:"add4xxRateAlarm"`
	// Experimental.
	Add5xxCountAlarm *map[string]*ErrorCountThreshold `field:"optional" json:"add5xxCountAlarm" yaml:"add5xxCountAlarm"`
	// Experimental.
	Add5xxRateAlarm *map[string]*ErrorRateThreshold `field:"optional" json:"add5xxRateAlarm" yaml:"add5xxRateAlarm"`
	// Experimental.
	AddHighTpsAlarm *map[string]*HighTpsThreshold `field:"optional" json:"addHighTpsAlarm" yaml:"addHighTpsAlarm"`
	// Experimental.
	AddIntegrationLatencyAverageAlarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyAverageAlarm" yaml:"addIntegrationLatencyAverageAlarm"`
	// Experimental.
	AddIntegrationLatencyP100Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP100Alarm" yaml:"addIntegrationLatencyP100Alarm"`
	// Experimental.
	AddIntegrationLatencyP50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP50Alarm" yaml:"addIntegrationLatencyP50Alarm"`
	// Experimental.
	AddIntegrationLatencyP70Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP70Alarm" yaml:"addIntegrationLatencyP70Alarm"`
	// Experimental.
	AddIntegrationLatencyP90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP90Alarm" yaml:"addIntegrationLatencyP90Alarm"`
	// Experimental.
	AddIntegrationLatencyP95Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP95Alarm" yaml:"addIntegrationLatencyP95Alarm"`
	// Experimental.
	AddIntegrationLatencyP9999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP9999Alarm" yaml:"addIntegrationLatencyP9999Alarm"`
	// Experimental.
	AddIntegrationLatencyP999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP999Alarm" yaml:"addIntegrationLatencyP999Alarm"`
	// Experimental.
	AddIntegrationLatencyP99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyP99Alarm" yaml:"addIntegrationLatencyP99Alarm"`
	// Experimental.
	AddIntegrationLatencyTM50Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM50Alarm" yaml:"addIntegrationLatencyTM50Alarm"`
	// Experimental.
	AddIntegrationLatencyTM70Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM70Alarm" yaml:"addIntegrationLatencyTM70Alarm"`
	// Experimental.
	AddIntegrationLatencyTM90Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM90Alarm" yaml:"addIntegrationLatencyTM90Alarm"`
	// Experimental.
	AddIntegrationLatencyTM95Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM95Alarm" yaml:"addIntegrationLatencyTM95Alarm"`
	// Experimental.
	AddIntegrationLatencyTM95OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM95OutlierAlarm" yaml:"addIntegrationLatencyTM95OutlierAlarm"`
	// Experimental.
	AddIntegrationLatencyTM9999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM9999Alarm" yaml:"addIntegrationLatencyTM9999Alarm"`
	// Experimental.
	AddIntegrationLatencyTM9999OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM9999OutlierAlarm" yaml:"addIntegrationLatencyTM9999OutlierAlarm"`
	// Experimental.
	AddIntegrationLatencyTM999Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM999Alarm" yaml:"addIntegrationLatencyTM999Alarm"`
	// Experimental.
	AddIntegrationLatencyTM999OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM999OutlierAlarm" yaml:"addIntegrationLatencyTM999OutlierAlarm"`
	// Experimental.
	AddIntegrationLatencyTM99Alarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM99Alarm" yaml:"addIntegrationLatencyTM99Alarm"`
	// Experimental.
	AddIntegrationLatencyTM99OutlierAlarm *map[string]*LatencyThreshold `field:"optional" json:"addIntegrationLatencyTM99OutlierAlarm" yaml:"addIntegrationLatencyTM99OutlierAlarm"`
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
	// See: DefaultLatencyTypesShown).
	//
	// Default: - p50, p90, p99 (.
	//
	// Experimental.
	LatencyTypesToRender *[]LatencyType `field:"optional" json:"latencyTypesToRender" yaml:"latencyTypesToRender"`
}

