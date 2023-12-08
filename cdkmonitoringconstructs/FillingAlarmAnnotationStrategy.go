package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Annotation strategy that fills the annotation provided, using the input and user requirements.
// Experimental.
type FillingAlarmAnnotationStrategy interface {
	IAlarmAnnotationStrategy
	// Creates annotation based on the metric and alarm properties.
	// Experimental.
	CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
	// Experimental.
	CreateAnnotationToFill(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
	// Experimental.
	GetAlarmingRangeShade(props *AlarmAnnotationStrategyProps) awscloudwatch.Shading
}

// The jsii proxy struct for FillingAlarmAnnotationStrategy
type jsiiProxy_FillingAlarmAnnotationStrategy struct {
	jsiiProxy_IAlarmAnnotationStrategy
}

// Experimental.
func NewFillingAlarmAnnotationStrategy_Override(f FillingAlarmAnnotationStrategy) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.FillingAlarmAnnotationStrategy",
		nil, // no parameters
		f,
	)
}

func (f *jsiiProxy_FillingAlarmAnnotationStrategy) CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := f.validateCreateAnnotationParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		f,
		"createAnnotation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FillingAlarmAnnotationStrategy) CreateAnnotationToFill(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := f.validateCreateAnnotationToFillParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		f,
		"createAnnotationToFill",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FillingAlarmAnnotationStrategy) GetAlarmingRangeShade(props *AlarmAnnotationStrategyProps) awscloudwatch.Shading {
	if err := f.validateGetAlarmingRangeShadeParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Shading

	_jsii_.Invoke(
		f,
		"getAlarmingRangeShade",
		[]interface{}{props},
		&returns,
	)

	return returns
}

