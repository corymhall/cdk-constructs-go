//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddAvailabilityAlarmParameters(metric interface{}, props *AvailabilityThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddHealthyTaskCountAlarmParameters(metric interface{}, props *HealthyTaskCountThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddHealthyTaskPercentAlarmParameters(metric interface{}, props *HealthyTaskPercentThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddMinRunningTaskCountAlarmParameters(metric interface{}, props *MinRunningTaskCountThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddRunningTaskCountAlarmParameters(metric interface{}, props *RunningTaskCountThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddRunningTaskRateAlarmParameters(metric interface{}, props *RunningTaskRateThreshold) error {
	return nil
}

func (t *jsiiProxy_TaskHealthAlarmFactory) validateAddUnhealthyTaskCountAlarmParameters(metric interface{}, props *UnhealthyTaskCountThreshold) error {
	return nil
}

func validateNewTaskHealthAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

