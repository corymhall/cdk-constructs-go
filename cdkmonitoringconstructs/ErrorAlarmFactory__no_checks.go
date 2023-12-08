//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ErrorAlarmFactory) validateAddErrorCountAlarmParameters(metric interface{}, errorType ErrorType, props *ErrorCountThreshold) error {
	return nil
}

func (e *jsiiProxy_ErrorAlarmFactory) validateAddErrorRateAlarmParameters(metric interface{}, errorType ErrorType, props *ErrorRateThreshold) error {
	return nil
}

func validateNewErrorAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

