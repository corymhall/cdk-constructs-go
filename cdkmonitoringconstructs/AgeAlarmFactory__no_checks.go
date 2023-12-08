//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AgeAlarmFactory) validateAddDaysSinceUpdateAlarmParameters(metric interface{}, props *DaysSinceUpdateThreshold) error {
	return nil
}

func (a *jsiiProxy_AgeAlarmFactory) validateAddDaysToExpiryAlarmParameters(metric interface{}, props *DaysToExpiryThreshold) error {
	return nil
}

func (a *jsiiProxy_AgeAlarmFactory) validateAddIteratorMaxAgeAlarmParameters(metric interface{}, props *MaxAgeThreshold) error {
	return nil
}

func validateNewAgeAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

