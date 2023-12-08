//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WafV2Monitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (w *jsiiProxy_WafV2Monitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (w *jsiiProxy_WafV2Monitoring) validateCreateAllowedRequestsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (w *jsiiProxy_WafV2Monitoring) validateCreateBlockedRequestsRateWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (w *jsiiProxy_WafV2Monitoring) validateCreateBlockedRequestsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (w *jsiiProxy_WafV2Monitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewWafV2MonitoringParameters(scope MonitoringScope, props *WafV2MonitoringProps) error {
	return nil
}

