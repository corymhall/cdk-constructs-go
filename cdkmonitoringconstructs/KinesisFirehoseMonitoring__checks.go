//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(alarm, func() string { return "parameter alarm" }); err != nil {
		return err
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	if alarmNamePrefix == nil {
		return fmt.Errorf("parameter alarmNamePrefix is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateCreateConversionWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateCreateIncomingRecordWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateCreateLimitWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (k *jsiiProxy_KinesisFirehoseMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewKinesisFirehoseMonitoringParameters(scope MonitoringScope, props *KinesisFirehoseMonitoringProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}
