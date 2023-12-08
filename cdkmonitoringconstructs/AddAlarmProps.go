package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Properties necessary to create a single alarm and configure it.
// Experimental.
type AddAlarmProps struct {
	// Alarm description is included in the ticket and therefore should describe what happened, with as much context as possible.
	// Experimental.
	AlarmDescription *string `field:"required" json:"alarmDescription" yaml:"alarmDescription"`
	// Suffix added to base alarm name.
	//
	// Alarm names need to be unique.
	// Experimental.
	AlarmNameSuffix *string `field:"required" json:"alarmNameSuffix" yaml:"alarmNameSuffix"`
	// Comparison operator used to compare actual value against the threshold.
	// Experimental.
	ComparisonOperator awscloudwatch.ComparisonOperator `field:"required" json:"comparisonOperator" yaml:"comparisonOperator"`
	// Threshold to alarm on.
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// Behaviour in case the metric data is missing.
	// Experimental.
	TreatMissingData awscloudwatch.TreatMissingData `field:"required" json:"treatMissingData" yaml:"treatMissingData"`
	// Allows to override the default action strategy.
	// Default: - default action will be used.
	//
	// Experimental.
	ActionOverride IAlarmActionStrategy `field:"optional" json:"actionOverride" yaml:"actionOverride"`
	// Enables the configured CloudWatch alarm ticketing actions.
	// Default: - the same as monitoring facade default.
	//
	// Experimental.
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// If this is defined, the default resource-specific alarm dedupe string will be set and this will be added as a suffix.
	//
	// This allows you to specify the same dedupe string for a family of alarms.
	// Cannot be defined at the same time as alarmDedupeStringOverride.
	// Default: - undefined (no suffix).
	//
	// Experimental.
	AlarmDedupeStringSuffix *string `field:"optional" json:"alarmDedupeStringSuffix" yaml:"alarmDedupeStringSuffix"`
	// A text included in the generated ticket description body, which fully replaces the generated text.
	// Default: - default auto-generated content only.
	//
	// Experimental.
	AlarmDescriptionOverride *string `field:"optional" json:"alarmDescriptionOverride" yaml:"alarmDescriptionOverride"`
	// If this is defined, the alarm name is set to this exact value.
	//
	// Please be aware that you need to specify prefix for different stages (Beta, Prod...) and regions (EU, NA...) manually.
	// Experimental.
	AlarmNameOverride *string `field:"optional" json:"alarmNameOverride" yaml:"alarmNameOverride"`
	// This allows user to attach custom parameters to this alarm, which can later be accessed from the "useCreatedAlarms" method.
	// Default: - no parameters.
	//
	// Experimental.
	CustomParams *map[string]interface{} `field:"optional" json:"customParams" yaml:"customParams"`
	// This allows user to attach custom values to this alarm, which can later be accessed from the "useCreatedAlarms" method.
	// Default: - no tags.
	//
	// Experimental.
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Number of breaches required to transition into an ALARM state.
	// Experimental.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// If this is defined, the alarm dedupe string is set to this exact value.
	//
	// Please be aware that you need to handle deduping for different stages (Beta, Prod...) and regions (EU, NA...) manually.
	// Default: - undefined (no override).
	//
	// Experimental.
	DedupeStringOverride *string `field:"optional" json:"dedupeStringOverride" yaml:"dedupeStringOverride"`
	// Disambiguator is a string that differentiates this alarm from other similar ones.
	// Default: - undefined (no disambiguator).
	//
	// Experimental.
	Disambiguator *string `field:"optional" json:"disambiguator" yaml:"disambiguator"`
	// An optional link included in the generated ticket description body.
	// Default: - no additional link will be added.
	//
	// Experimental.
	DocumentationLink *string `field:"optional" json:"documentationLink" yaml:"documentationLink"`
	// Used only for alarms based on percentiles.
	//
	// If you specify <code>false</code>, the alarm state does not change during periods with too few data points to be statistically significant.
	// If you specify <code>true</code>, the alarm is always evaluated and possibly changes state no matter how many data points are available.
	// Default: - true.
	//
	// Experimental.
	EvaluateLowSampleCountPercentile *bool `field:"optional" json:"evaluateLowSampleCountPercentile" yaml:"evaluateLowSampleCountPercentile"`
	// Number of periods to consider when checking the number of breaching datapoints.
	// Default: - Same as datapointsToAlarm.
	//
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// Indicates whether the alarming range of values should be highlighted in the widget.
	// Default: - false.
	//
	// Experimental.
	FillAlarmRange *bool `field:"optional" json:"fillAlarmRange" yaml:"fillAlarmRange"`
	// If specified, adjusts the metric before creating an alarm from it.
	// Default: - no adjuster.
	//
	// Experimental.
	MetricAdjuster IMetricAdjuster `field:"optional" json:"metricAdjuster" yaml:"metricAdjuster"`
	// Specifies how many samples (N) of the metric is needed to trigger the alarm.
	//
	// If this property is specified, an artificial composite alarm is created of the following:
	// <ul>
	// <li>The original alarm, created without this property being used; this alarm will have no actions set.</li>
	// <li>A secondary alarm, which will monitor the same metric with the N (SampleCount) statistic, checking the sample count.</li>
	// </ul>
	// The newly created composite alarm will be returned as a result, and it will take the original alarm actions.
	// Default: - default behaviour - no condition on sample count will be added to the alarm.
	//
	// Deprecated: Use minSampleCountToEvaluateDatapoint instead. minMetricSamplesAlarm uses different evaluation
	// period for its child alarms, so it doesn't guarantee that each datapoint in the evaluation period has
	// sufficient number of samples.
	MinMetricSamplesToAlarm *float64 `field:"optional" json:"minMetricSamplesToAlarm" yaml:"minMetricSamplesToAlarm"`
	// Specifies how many samples (N) of the metric is needed in a datapoint to be evaluated for alarming.
	//
	// If this property is specified, your metric will be subject to MathExpression that will add an IF condition
	// to your metric to make sure that each datapoint is evaluated only if it has sufficient number of samples.
	// If the number of samples is not sufficient, the datapoint will be treated as missing data and will be evaluated
	// according to the treatMissingData parameter.
	// If specified, deprecated minMetricSamplesToAlarm has no effect.
	// Default: - default behaviour - no condition on sample count will be used.
	//
	// Experimental.
	MinSampleCountToEvaluateDatapoint *float64 `field:"optional" json:"minSampleCountToEvaluateDatapoint" yaml:"minSampleCountToEvaluateDatapoint"`
	// If specified, it modifies the final alarm annotation color.
	// Default: - no override (default color).
	//
	// Experimental.
	OverrideAnnotationColor *string `field:"optional" json:"overrideAnnotationColor" yaml:"overrideAnnotationColor"`
	// If specified, it modifies the final alarm annotation label.
	// Default: - no override (default label).
	//
	// Experimental.
	OverrideAnnotationLabel *string `field:"optional" json:"overrideAnnotationLabel" yaml:"overrideAnnotationLabel"`
	// If specified, it modifies the final alarm annotation visibility.
	// Default: - no override (default visibility).
	//
	// Experimental.
	OverrideAnnotationVisibility *bool `field:"optional" json:"overrideAnnotationVisibility" yaml:"overrideAnnotationVisibility"`
	// Period override for the metric to alarm on.
	// Default: - the default specified in MetricFactory.
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// An optional link included in the generated ticket description body.
	// Default: - no additional link will be added.
	//
	// Experimental.
	RunbookLink *string `field:"optional" json:"runbookLink" yaml:"runbookLink"`
	// This property is required in the following situation: <ol>      <li><code>minSampleCountToEvaluateDatapoint</code> is specified</li>      <li>the metric used for the alarm is a <code>MathExpression</code></li>      <li>the <code>MathExpression</code> is composed of more than one metric</li> </ol>.
	//
	// In this situation, this property indicates the metric Id in the MathExpression’s <code>usingMetrics</code>
	// property that should be used as the sampleCount metric for the new MathExpression as described in the documentation
	// for <code>minSampleCountToEvaluateDatapoint</code>.
	// Experimental.
	SampleCountMetricId *string `field:"optional" json:"sampleCountMetricId" yaml:"sampleCountMetricId"`
}

