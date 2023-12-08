package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Default annotation strategy that returns the built-in alarm annotation.
// Experimental.
type DefaultAlarmAnnotationStrategy interface {
	FillingAlarmAnnotationStrategy
	// Creates annotation based on the metric and alarm properties.
	// Experimental.
	CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
	// Experimental.
	CreateAnnotationToFill(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
	// Experimental.
	GetAlarmingRangeShade(props *AlarmAnnotationStrategyProps) awscloudwatch.Shading
}

// The jsii proxy struct for DefaultAlarmAnnotationStrategy
type jsiiProxy_DefaultAlarmAnnotationStrategy struct {
	jsiiProxy_FillingAlarmAnnotationStrategy
}

// Experimental.
func NewDefaultAlarmAnnotationStrategy() DefaultAlarmAnnotationStrategy {
	_init_.Initialize()

	j := jsiiProxy_DefaultAlarmAnnotationStrategy{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultAlarmAnnotationStrategy",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultAlarmAnnotationStrategy_Override(d DefaultAlarmAnnotationStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultAlarmAnnotationStrategy",
		nil, // no parameters
		d,
	)
}

func (d *jsiiProxy_DefaultAlarmAnnotationStrategy) CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := d.validateCreateAnnotationParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		d,
		"createAnnotation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultAlarmAnnotationStrategy) CreateAnnotationToFill(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := d.validateCreateAnnotationToFillParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		d,
		"createAnnotationToFill",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DefaultAlarmAnnotationStrategy) GetAlarmingRangeShade(props *AlarmAnnotationStrategyProps) awscloudwatch.Shading {
	if err := d.validateGetAlarmingRangeShadeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Shading

	_jsii_.Invoke(
		d,
		"getAlarmingRangeShade",
		[]interface{}{props},
		&returns,
	)

	return returns
}

