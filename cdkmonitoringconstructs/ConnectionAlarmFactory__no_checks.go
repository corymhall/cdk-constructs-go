//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConnectionAlarmFactory) validateAddMaxConnectionCountAlarmParameters(metric interface{}, props *HighConnectionCountThreshold) error {
	return nil
}

func (c *jsiiProxy_ConnectionAlarmFactory) validateAddMinConnectionCountAlarmParameters(metric interface{}, props *LowConnectionCountThreshold) error {
	return nil
}

func validateNewConnectionAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

