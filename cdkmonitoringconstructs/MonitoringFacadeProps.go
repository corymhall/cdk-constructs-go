package cdkmonitoringconstructs


// Experimental.
type MonitoringFacadeProps struct {
	// Defaults for alarm factory.
	// Default: - actions enabled, facade logical ID used as default alarm name prefix.
	//
	// Experimental.
	AlarmFactoryDefaults *AlarmFactoryDefaults `field:"optional" json:"alarmFactoryDefaults" yaml:"alarmFactoryDefaults"`
	// Defaults for dashboard factory.
	// Default: - An instance of {@link DynamicDashboardFactory}; facade logical ID used as default name.
	//
	// Experimental.
	DashboardFactory IDynamicDashboardFactory `field:"optional" json:"dashboardFactory" yaml:"dashboardFactory"`
	// Defaults for metric factory.
	// Default: - empty (no preferences).
	//
	// Experimental.
	MetricFactoryDefaults *MetricFactoryDefaults `field:"optional" json:"metricFactoryDefaults" yaml:"metricFactoryDefaults"`
}

