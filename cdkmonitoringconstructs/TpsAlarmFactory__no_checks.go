//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TpsAlarmFactory) validateAddMaxTpsAlarmParameters(metric interface{}, props *HighTpsThreshold) error {
	return nil
}

func (t *jsiiProxy_TpsAlarmFactory) validateAddMinTpsAlarmParameters(metric interface{}, props *LowTpsThreshold) error {
	return nil
}

func validateNewTpsAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

