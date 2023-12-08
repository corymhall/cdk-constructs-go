//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KinesisAlarmFactory) validateAddFirehoseStreamExceedSafetyThresholdAlarmParameters(metric interface{}, metricName *string, quotaName *string, props *FirehoseStreamLimitThreshold) error {
	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddIteratorMaxAgeAlarmParameters(metric interface{}, props *MaxIteratorAgeThreshold) error {
	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddProvisionedReadThroughputExceededAlarmParameters(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) error {
	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddProvisionedWriteThroughputExceededAlarmParameters(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) error {
	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddPutRecordsFailedAlarmParameters(metric interface{}, props *RecordsFailedThreshold) error {
	return nil
}

func (k *jsiiProxy_KinesisAlarmFactory) validateAddPutRecordsThrottledAlarmParameters(metric interface{}, props *RecordsThrottledThreshold) error {
	return nil
}

func validateNewKinesisAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

