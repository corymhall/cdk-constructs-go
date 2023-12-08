//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/constructs-go/constructs/v10"
)

func (a *jsiiProxy_AlarmFactory) validateAddAlarmParameters(metric interface{}, props *AddAlarmProps) error {
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

func (a *jsiiProxy_AlarmFactory) validateAddCompositeAlarmParameters(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) error {
	if alarms == nil {
		return fmt.Errorf("parameter alarms is required, but nil was provided")
	}
	for idx_423c13, v := range *alarms {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter alarms[%#v]", idx_423c13) }); err != nil {
			return err
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

func (a *jsiiProxy_AlarmFactory) validateCreateAnnotationParameters(props *AlarmAnnotationStrategyProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlarmFactory) validateDetermineCompositeAlarmRuleParameters(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) error {
	if alarms == nil {
		return fmt.Errorf("parameter alarms is required, but nil was provided")
	}
	for idx_423c13, v := range *alarms {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter alarms[%#v]", idx_423c13) }); err != nil {
			return err
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

func (a *jsiiProxy_AlarmFactory) validateGenerateDescriptionParameters(alarmDescription *string) error {
	if alarmDescription == nil {
		return fmt.Errorf("parameter alarmDescription is required, but nil was provided")
	}

	return nil
}

func validateNewAlarmFactoryParameters(alarmScope constructs.Construct, props *AlarmFactoryProps) error {
	if alarmScope == nil {
		return fmt.Errorf("parameter alarmScope is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

