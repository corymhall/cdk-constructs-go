//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (e *jsiiProxy_ExtendDedupeString) validateProcessDedupeStringParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_ExtendDedupeString) validateProcessDedupeStringOverrideParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

