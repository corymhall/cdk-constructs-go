//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (n *jsiiProxy_NoopAlarmActionStrategy) validateAddAlarmActionsParameters(_props *AlarmActionStrategyProps) error {
	if _props == nil {
		return fmt.Errorf("parameter _props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(_props, func() string { return "parameter _props" }); err != nil {
		return err
	}

	return nil
}

