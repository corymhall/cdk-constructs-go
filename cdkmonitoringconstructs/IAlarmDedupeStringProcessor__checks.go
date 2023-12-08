//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (i *jsiiProxy_IAlarmDedupeStringProcessor) validateProcessDedupeStringParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_IAlarmDedupeStringProcessor) validateProcessDedupeStringOverrideParameters(dedupeString *string) error {
	if dedupeString == nil {
		return fmt.Errorf("parameter dedupeString is required, but nil was provided")
	}

	return nil
}

