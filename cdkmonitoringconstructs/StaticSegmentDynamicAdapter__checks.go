//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (s *jsiiProxy_StaticSegmentDynamicAdapter) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewStaticSegmentDynamicAdapterParameters(props IDashboardFactoryProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

