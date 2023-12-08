//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsTopicMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_SnsTopicMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_SnsTopicMonitoring) validateCreateMessageCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicMonitoring) validateCreateMessageFailedWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicMonitoring) validateCreateMessageSizeWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SnsTopicMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSnsTopicMonitoringParameters(scope MonitoringScope, props *SnsTopicMonitoringProps) error {
	return nil
}

