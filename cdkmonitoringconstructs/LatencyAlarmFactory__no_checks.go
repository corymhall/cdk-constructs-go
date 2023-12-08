//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LatencyAlarmFactory) validateAddDurationAlarmParameters(metric interface{}, latencyType LatencyType, props *DurationThreshold) error {
	return nil
}

func (l *jsiiProxy_LatencyAlarmFactory) validateAddIntegrationLatencyAlarmParameters(metric interface{}, latencyType LatencyType, props *LatencyThreshold) error {
	return nil
}

func (l *jsiiProxy_LatencyAlarmFactory) validateAddJvmGarbageCollectionDurationAlarmParameters(metric interface{}, latencyType LatencyType, props *DurationThreshold) error {
	return nil
}

func (l *jsiiProxy_LatencyAlarmFactory) validateAddLatencyAlarmParameters(metric interface{}, latencyType LatencyType, props *LatencyThreshold) error {
	return nil
}

func validateNewLatencyAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

