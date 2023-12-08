//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ApiGatewayMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateCreateErrorCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateCreateErrorRateWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateCreateLatencyWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateCreateTpsWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (a *jsiiProxy_ApiGatewayMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewApiGatewayMonitoringParameters(scope MonitoringScope, props *ApiGatewayMonitoringProps) error {
	return nil
}

