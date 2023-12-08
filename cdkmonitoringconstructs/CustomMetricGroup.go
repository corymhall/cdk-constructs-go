package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Custom metric group represents a single widget.
// Experimental.
type CustomMetricGroup struct {
	// list of metrics in the group (can be defined in different ways, see the type documentation).
	// Experimental.
	Metrics *[]interface{} `field:"required" json:"metrics" yaml:"metrics"`
	// title of the whole group.
	// Experimental.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Flag indicating this metric group should be included in the summary as well.
	// Default: - addToSummaryDashboard from CustomMonitoringProps, defaulting to false.
	//
	// Experimental.
	AddToSummaryDashboard *bool `field:"optional" json:"addToSummaryDashboard" yaml:"addToSummaryDashboard"`
	// optional axis.
	// Default: undefined.
	//
	// Experimental.
	GraphWidgetAxis *awscloudwatch.YAxisProps `field:"optional" json:"graphWidgetAxis" yaml:"graphWidgetAxis"`
	// graph widget legend.
	// Default: BOTTOM.
	//
	// Experimental.
	GraphWidgetLegend awscloudwatch.LegendPosition `field:"optional" json:"graphWidgetLegend" yaml:"graphWidgetLegend"`
	// optional right axis.
	// Default: undefined.
	//
	// Experimental.
	GraphWidgetRightAxis *awscloudwatch.YAxisProps `field:"optional" json:"graphWidgetRightAxis" yaml:"graphWidgetRightAxis"`
	// type of the widget.
	// Default: line.
	//
	// Experimental.
	GraphWidgetType GraphWidgetType `field:"optional" json:"graphWidgetType" yaml:"graphWidgetType"`
	// optional custom horizontal annotations which will be displayed over the metrics on the left axis (if there are any alarms, any existing annotations will be merged together).
	// Experimental.
	HorizontalAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"optional" json:"horizontalAnnotations" yaml:"horizontalAnnotations"`
	// optional custom horizontal annotations which will be displayed over the metrics on the right axis (if there are any alarms, any existing annotations will be merged together).
	// Experimental.
	HorizontalRightAnnotations *[]*awscloudwatch.HorizontalAnnotation `field:"optional" json:"horizontalRightAnnotations" yaml:"horizontalRightAnnotations"`
	// See: addToSummaryDashboard.
	//
	// Deprecated: use addToSummaryDashboard. addToSummaryDashboard will take precedence over important.
	Important *bool `field:"optional" json:"important" yaml:"important"`
}

