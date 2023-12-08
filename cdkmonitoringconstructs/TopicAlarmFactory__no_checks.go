//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TopicAlarmFactory) validateAddMaxMessagesPublishedAlarmParameters(metric interface{}, props *HighMessagesPublishedThreshold) error {
	return nil
}

func (t *jsiiProxy_TopicAlarmFactory) validateAddMessageNotificationsFailedAlarmParameters(metric interface{}, props *NotificationsFailedThreshold) error {
	return nil
}

func (t *jsiiProxy_TopicAlarmFactory) validateAddMinMessagesPublishedAlarmParameters(metric interface{}, props *LowMessagesPublishedThreshold) error {
	return nil
}

func validateNewTopicAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

