//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxCpuUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxDiskUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxFileDescriptorPercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxHeapMemoryAfterGCUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxMasterCpuUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxMasterMemoryUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxMemoryUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxThreadCountUsageAlarmParameters(percentMetric interface{}, props *UsageCountThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMaxUsageCountAlarmParameters(metric interface{}, props *UsageCountThreshold) error {
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

func (u *jsiiProxy_UsageAlarmFactory) validateAddMemoryUsagePercentAlarmParameters(percentMetric interface{}, props *UsageThreshold, usageType UsageType) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
		}
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	if usageType == "" {
		return fmt.Errorf("parameter usageType is required, but nil was provided")
	}

	return nil
}

func (u *jsiiProxy_UsageAlarmFactory) validateAddMinUsageCountAlarmParameters(percentMetric interface{}, props *MinUsageCountThreshold) error {
	if percentMetric == nil {
		return fmt.Errorf("parameter percentMetric is required, but nil was provided")
	}
	switch percentMetric.(type) {
	case awscloudwatch.Metric:
		// ok
	case awscloudwatch.MathExpression:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(percentMetric) {
			return fmt.Errorf("parameter percentMetric must be one of the allowed types: awscloudwatch.Metric, awscloudwatch.MathExpression; received %#v (a %T)", percentMetric, percentMetric)
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

func validateNewUsageAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	if alarmFactory == nil {
		return fmt.Errorf("parameter alarmFactory is required, but nil was provided")
	}

	return nil
}

