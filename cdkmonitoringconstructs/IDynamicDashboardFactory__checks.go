//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (i *jsiiProxy_IDynamicDashboardFactory) validateAddDynamicSegmentParameters(segment IDynamicDashboardSegment) error {
	if segment == nil {
		return fmt.Errorf("parameter segment is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IDynamicDashboardFactory) validateGetDashboardParameters(type_ *string) error {
	if type_ == nil {
		return fmt.Errorf("parameter type_ is required, but nil was provided")
	}

	return nil
}

