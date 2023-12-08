//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMonitoringNamingStrategy_IsAlarmFriendlyParameters(str *string) error {
	if str == nil {
		return fmt.Errorf("parameter str is required, but nil was provided")
	}

	return nil
}

func validateNewMonitoringNamingStrategyParameters(input *NameResolutionInput) error {
	if input == nil {
		return fmt.Errorf("parameter input is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(input, func() string { return "parameter input" }); err != nil {
		return err
	}

	return nil
}

