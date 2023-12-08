//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (i *jsiiProxy_IDynamicDashboardSegment) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

