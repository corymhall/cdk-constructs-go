//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (q *jsiiProxy_QueueAlarmFactory) validateAddMaxQueueIncomingMessagesCountAlarmParameters(metric interface{}, props *MaxIncomingMessagesCountThreshold) error {
	return nil
}

func (q *jsiiProxy_QueueAlarmFactory) validateAddMaxQueueMessageAgeAlarmParameters(metric interface{}, props *MaxMessageAgeThreshold) error {
	return nil
}

func (q *jsiiProxy_QueueAlarmFactory) validateAddMaxQueueMessageCountAlarmParameters(metric interface{}, props *MaxMessageCountThreshold) error {
	return nil
}

func (q *jsiiProxy_QueueAlarmFactory) validateAddMaxQueueTimeToDrainMessagesAlarmParameters(metric interface{}, props *MaxTimeToDrainThreshold) error {
	return nil
}

func (q *jsiiProxy_QueueAlarmFactory) validateAddMinQueueIncomingMessagesCountAlarmParameters(metric interface{}, props *MinIncomingMessagesCountThreshold) error {
	return nil
}

func (q *jsiiProxy_QueueAlarmFactory) validateAddMinQueueMessageCountAlarmParameters(metric interface{}, props *MinMessageCountThreshold) error {
	return nil
}

func validateNewQueueAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

