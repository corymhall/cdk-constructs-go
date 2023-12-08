//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatchactions"
)

func (o *jsiiProxy_OpsItemAlarmActionStrategy) validateAddAlarmActionsParameters(props *AlarmActionStrategyProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNewOpsItemAlarmActionStrategyParameters(severity awscloudwatchactions.OpsItemSeverity) error {
	if severity == "" {
		return fmt.Errorf("parameter severity is required, but nil was provided")
	}

	return nil
}

