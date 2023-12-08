//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (k *jsiiProxy_KinesisAlarmFactory) validateAddFirehoseStreamExceedSafetyThresholdAlarmParameters(metric interface{}, metricName *string, quotaName *string, props *FirehoseStreamLimitThreshold) error {
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

	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if quotaName == nil {
		return fmt.Errorf("parameter quotaName is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddIteratorMaxAgeAlarmParameters(metric interface{}, props *MaxIteratorAgeThreshold) error {
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

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddProvisionedReadThroughputExceededAlarmParameters(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) error {
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

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if disambiguator == nil {
		return fmt.Errorf("parameter disambiguator is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddProvisionedWriteThroughputExceededAlarmParameters(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) error {
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

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if disambiguator == nil {
		return fmt.Errorf("parameter disambiguator is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddPutRecordsFailedAlarmParameters(metric interface{}, props *RecordsFailedThreshold) error {
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

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddPutRecordsThrottledAlarmParameters(metric interface{}, props *RecordsThrottledThreshold) error {
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

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewKinesisAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	if alarmFactory == nil {
		return fmt.Errorf("parameter alarmFactory is required, but nil was provided")
	}

	return nil
}

