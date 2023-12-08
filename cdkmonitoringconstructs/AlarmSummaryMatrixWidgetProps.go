package cdkmonitoringconstructs


// Experimental.
type AlarmSummaryMatrixWidgetProps struct {
	// Height of the widget.
	// Default: - 6 for Alarm and Graph widgets.
	// 3 for single value widgets where most recent value of a metric is displayed.
	//
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// The region the metrics of this graph should be taken from.
	// Default: - Current region.
	//
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Title for the graph.
	// Default: - None.
	//
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: 6.
	//
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
	// Experimental.
	AlarmArns *[]*string `field:"required" json:"alarmArns" yaml:"alarmArns"`
}

