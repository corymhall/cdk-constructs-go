package cdkmonitoringconstructs


// Custom metric with an alarm defined.
// Experimental.
type CustomMetricWithAlarm struct {
	// alarm definitions.
	// Experimental.
	AddAlarm *map[string]*CustomThreshold `field:"required" json:"addAlarm" yaml:"addAlarm"`
	// alarm friendly name.
	// Experimental.
	AlarmFriendlyName *string `field:"required" json:"alarmFriendlyName" yaml:"alarmFriendlyName"`
	// metric to alarm on.
	// Experimental.
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// axis (right or left) on which to graph metric default: AxisPosition.LEFT.
	// Experimental.
	Position AxisPosition `field:"optional" json:"position" yaml:"position"`
}

