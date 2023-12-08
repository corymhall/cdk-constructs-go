//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisDataAnalyticsAlarmFactory) validateAddDowntimeAlarmParameters(metric interface{}, props *MaxDowntimeThreshold) error {
	return nil
}

func (k *jsiiProxy_KinesisDataAnalyticsAlarmFactory) validateAddFullRestartAlarmParameters(metric interface{}, props *FullRestartCountThreshold) error {
	return nil
}

func validateNewKinesisDataAnalyticsAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

