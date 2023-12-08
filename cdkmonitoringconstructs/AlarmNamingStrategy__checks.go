//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_AlarmNamingStrategy) validateGetDedupeStringParameters(props *AlarmNamingInput) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlarmNamingStrategy) validateGetNameParameters(props *AlarmNamingInput) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlarmNamingStrategy) validateGetWidgetLabelParameters(props *AlarmNamingInput) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func (a *jsiiProxy_AlarmNamingStrategy) validateJoinDistinctParameters(parts *[]*string, separator *string) error {
	if parts == nil {
		return fmt.Errorf("parameter parts is required, but nil was provided")
	}

	if separator == nil {
		return fmt.Errorf("parameter separator is required, but nil was provided")
	}

	return nil
}

func validateNewAlarmNamingStrategyParameters(globalPrefix *string, localPrefix *string) error {
	if globalPrefix == nil {
		return fmt.Errorf("parameter globalPrefix is required, but nil was provided")
	}

	if localPrefix == nil {
		return fmt.Errorf("parameter localPrefix is required, but nil was provided")
	}

	return nil
}

