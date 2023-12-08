//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (h *jsiiProxy_HeaderWidget) validatePositionParameters(x *float64, y *float64) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	if y == nil {
		return fmt.Errorf("parameter y is required, but nil was provided")
	}

	return nil
}

func validateNewHeaderWidgetParameters(text *string) error {
	if text == nil {
		return fmt.Errorf("parameter text is required, but nil was provided")
	}

	return nil
}

