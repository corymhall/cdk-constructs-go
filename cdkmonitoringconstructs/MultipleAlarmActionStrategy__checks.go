//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (m *jsiiProxy_MultipleAlarmActionStrategy) validateAddAlarmActionsParameters(props *AlarmActionStrategyProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewMultipleAlarmActionStrategyParameters(actions *[]IAlarmActionStrategy) error {
	if actions == nil {
		return fmt.Errorf("parameter actions is required, but nil was provided")
	}

	return nil
}

