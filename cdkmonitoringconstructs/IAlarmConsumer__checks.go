//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (i *jsiiProxy_IAlarmConsumer) validateConsumeParameters(alarms *[]*AlarmWithAnnotation) error {
	if alarms == nil {
		return fmt.Errorf("parameter alarms is required, but nil was provided")
	}
	for idx_423c13, v := range *alarms {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter alarms[%#v]", idx_423c13) }); err != nil {
			return err
		}
	}

	return nil
}

