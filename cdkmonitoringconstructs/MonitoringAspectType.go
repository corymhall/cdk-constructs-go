package cdkmonitoringconstructs


// Experimental.
type MonitoringAspectType struct {
	// If the monitoring aspect is enabled for this resource.
	// Default: - true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
	// The monitoring props for this resource.
	// Default: - none.
	//
	// Experimental.
	Props interface{} `field:"optional" json:"props" yaml:"props"`
}

