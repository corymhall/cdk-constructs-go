//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DocumentDbMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	if alarm == nil {
		return fmt.Errorf("parameter alarm is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(alarm, func() string { return "parameter alarm" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	if alarmNamePrefix == nil {
		return fmt.Errorf("parameter alarmNamePrefix is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateConnectionsWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateResourceUsageWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateCreateTransactionsWidgetParameters(width *float64, height *float64) error {
	if width == nil {
		return fmt.Errorf("parameter width is required, but nil was provided")
	}

	if height == nil {
		return fmt.Errorf("parameter height is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateNewDocumentDbMonitoringParameters(scope MonitoringScope, props *DocumentDbMonitoringProps) error {
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

