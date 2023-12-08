//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BillingMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (b *jsiiProxy_BillingMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (b *jsiiProxy_BillingMonitoring) validateCreateChargesByServiceWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (b *jsiiProxy_BillingMonitoring) validateCreateTotalChargesWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (b *jsiiProxy_BillingMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewBillingMonitoringParameters(scope MonitoringScope, props *BillingMonitoringProps) error {
	return nil
}

