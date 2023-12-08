//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SqsQueueMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateMessageAgeWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateMessageCountWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateMessageSizeWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateProducerAndConsumerRateWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateCreateTimeToDrainWidgetParameters(width *float64, height *float64) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateResolveQueueNameParameters(queue awssqs.IQueue) error {
	return nil
}

func (s *jsiiProxy_SqsQueueMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSqsQueueMonitoringParameters(scope MonitoringScope, props *SqsQueueMonitoringProps) error {
	return nil
}

