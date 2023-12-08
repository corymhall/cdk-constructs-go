//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CustomAlarmFactory) validateAddCustomAlarmParameters(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *CustomThreshold) error {
	return nil
}

func validateNewCustomAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

