//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DynamoAlarmFactory) validateAddConsumedCapacityAlarmParameters(metric interface{}, capacityType CapacityType, props *ConsumedCapacityThreshold) error {
	return nil
}

func (d *jsiiProxy_DynamoAlarmFactory) validateAddThrottledEventsAlarmParameters(metric interface{}, capacityType CapacityType, props *ThrottledEventsThreshold) error {
	return nil
}

func validateNewDynamoAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

