//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AnomalyDetectingAlarmFactory) validateAddAlarmWhenOutOfBandParameters(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *AnomalyDetectionThreshold) error {
	return nil
}

func validateNewAnomalyDetectingAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

