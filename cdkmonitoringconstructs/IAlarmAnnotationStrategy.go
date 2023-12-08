package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Helper class for creating annotations for alarms.
// Experimental.
type IAlarmAnnotationStrategy interface {
	// Creates annotation based on the metric and alarm properties.
	// Experimental.
	CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation
}

// The jsii proxy for IAlarmAnnotationStrategy
type jsiiProxy_IAlarmAnnotationStrategy struct {
	_ byte // padding
}

func (i *jsiiProxy_IAlarmAnnotationStrategy) CreateAnnotation(props *AlarmAnnotationStrategyProps) *awscloudwatch.HorizontalAnnotation {
	if err := i.validateCreateAnnotationParameters(props); err != nil {
		panic(err)
	}
	var returns *awscloudwatch.HorizontalAnnotation

	_jsii_.Invoke(
		i,
		"createAnnotation",
		[]interface{}{props},
		&returns,
	)

	return returns
}

