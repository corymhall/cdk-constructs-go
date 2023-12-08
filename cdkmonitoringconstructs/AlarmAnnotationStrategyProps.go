package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type AlarmAnnotationStrategyProps struct {
	// Experimental.
	Action IAlarmActionStrategy `field:"required" json:"action" yaml:"action"`
	// Experimental.
	CustomParams *map[string]interface{} `field:"optional" json:"customParams" yaml:"customParams"`
	// Experimental.
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Experimental.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Experimental.
	Disambiguator *string `field:"optional" json:"disambiguator" yaml:"disambiguator"`
	// Experimental.
	Alarm awscloudwatch.Alarm `field:"required" json:"alarm" yaml:"alarm"`
	// Experimental.
	ComparisonOperator awscloudwatch.ComparisonOperator `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Experimental.
	DatapointsToAlarm *float64 `field:"required" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Experimental.
	EvaluationPeriods *float64 `field:"required" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Experimental.
	FillAlarmRange *bool `field:"required" json:"fillAlarmRange" yaml:"fillAlarmRange"`
	// Experimental.
	Metric interface{} `field:"required" json:"metric" yaml:"metric"`
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Experimental.
	MinMetricSamplesToAlarm *float64 `field:"optional" json:"minMetricSamplesToAlarm" yaml:"minMetricSamplesToAlarm"`
	// Experimental.
	MinSampleCountToEvaluateDatapoint *float64 `field:"optional" json:"minSampleCountToEvaluateDatapoint" yaml:"minSampleCountToEvaluateDatapoint"`
	// Experimental.
	OverrideAnnotationColor *string `field:"optional" json:"overrideAnnotationColor" yaml:"overrideAnnotationColor"`
	// Experimental.
	OverrideAnnotationLabel *string `field:"optional" json:"overrideAnnotationLabel" yaml:"overrideAnnotationLabel"`
	// Experimental.
	OverrideAnnotationVisibility *bool `field:"optional" json:"overrideAnnotationVisibility" yaml:"overrideAnnotationVisibility"`
}

