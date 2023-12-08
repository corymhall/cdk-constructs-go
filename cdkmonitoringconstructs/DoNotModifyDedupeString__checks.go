//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (d *jsiiProxy_DoNotModifyDedupeString) validateProcessDedupeStringParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DoNotModifyDedupeString) validateProcessDedupeStringOverrideParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

