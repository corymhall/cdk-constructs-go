package cdkmonitoringconstructs


// Experimental.
type MetricFactoryProps struct {
	// Allows you to specify the global defaults, which can be overridden in the individual metrics or alarms.
	// Experimental.
	GlobalDefaults *MetricFactoryDefaults `field:"optional" json:"globalDefaults" yaml:"globalDefaults"`
}

