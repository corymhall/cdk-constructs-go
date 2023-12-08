package cdkmonitoringconstructs


// Experimental.
type CertificateManagerMonitoringOptions struct {
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
	AddDaysToExpiryAlarm *map[string]*DaysToExpiryThreshold `field:"optional" json:"addDaysToExpiryAlarm" yaml:"addDaysToExpiryAlarm"`
}

