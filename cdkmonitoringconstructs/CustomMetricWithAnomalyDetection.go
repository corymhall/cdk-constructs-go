package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Custom metric with anomaly detection.
// Experimental.
type CustomMetricWithAnomalyDetection struct {
	// alarm friendly name.
	// Experimental.
	AlarmFriendlyName *string `field:"required" json:"alarmFriendlyName" yaml:"alarmFriendlyName"`
	// standard deviation for the anomaly detection to be rendered on the graph widget.
	// Experimental.
	AnomalyDetectionStandardDeviationToRender *float64 `field:"required" json:"anomalyDetectionStandardDeviationToRender" yaml:"anomalyDetectionStandardDeviationToRender"`
	// metric to alarm on.
	// Experimental.
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// adds alarm on a detected anomaly.
	// Experimental.
	AddAlarmOnAnomaly *map[string]*AnomalyDetectionThreshold `field:"optional" json:"addAlarmOnAnomaly" yaml:"addAlarmOnAnomaly"`
	// anomaly detection period.
	// Default: - metric period (if defined) or global default.
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
}

