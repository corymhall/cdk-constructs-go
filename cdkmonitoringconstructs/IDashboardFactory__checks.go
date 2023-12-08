//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (i *jsiiProxy_IDashboardFactory) validateAddSegmentParameters(props IDashboardFactoryProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}

	return nil
}

