//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (e *jsiiProxy_ErrorAlarmFactory) validateAddErrorCountAlarmParameters(metric interface{}, errorType ErrorType, props *ErrorCountThreshold) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	if errorType == "" {
		return fmt.Errorf("parameter errorType is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (e *jsiiProxy_ErrorAlarmFactory) validateAddErrorRateAlarmParameters(metric interface{}, errorType ErrorType, props *ErrorRateThreshold) error {
	if metric == nil {
		return fmt.Errorf("parameter metric is required, but nil was provided")
	}
	switch metric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(metric) {
			return fmt.Errorf("parameter metric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", metric, metric)
		}
	}

	if errorType == "" {
		return fmt.Errorf("parameter errorType is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewErrorAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	if alarmFactory == nil {
		return fmt.Errorf("parameter alarmFactory is required, but nil was provided")
	}

	return nil
}

