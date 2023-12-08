package cdkmonitoringconstructs


// Experimental.
type AlarmFactoryProps struct {
	// Experimental.
	GlobalAlarmDefaults *AlarmFactoryDefaults `field:"required" json:"globalAlarmDefaults" yaml:"globalAlarmDefaults"`
	// Experimental.
	GlobalMetricDefaults *MetricFactoryDefaults `field:"required" json:"globalMetricDefaults" yaml:"globalMetricDefaults"`
	// Experimental.
	LocalAlarmNamePrefix *string `field:"required" json:"localAlarmNamePrefix" yaml:"localAlarmNamePrefix"`
}

