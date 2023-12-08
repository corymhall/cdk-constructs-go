//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LogLevelAlarmFactory) validateAddLogCountAlarmParameters(metric interface{}, logLevel LogLevel, props *LogLevelCountThreshold) error {
	return nil
}

func validateNewLogLevelAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

