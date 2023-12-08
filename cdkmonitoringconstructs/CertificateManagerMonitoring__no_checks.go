//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CertificateManagerMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (c *jsiiProxy_CertificateManagerMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (c *jsiiProxy_CertificateManagerMonitoring) validateCreateDaysToExpiryWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (c *jsiiProxy_CertificateManagerMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewCertificateManagerMonitoringParameters(scope MonitoringScope, props *CertificateManagerMonitoringProps) error {
	return nil
}

