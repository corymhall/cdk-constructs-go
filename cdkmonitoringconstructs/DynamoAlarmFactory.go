package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type DynamoAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddConsumedCapacityAlarm(metric interface{}, capacityType CapacityType, props *ConsumedCapacityThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddThrottledEventsAlarm(metric interface{}, capacityType CapacityType, props *ThrottledEventsThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for DynamoAlarmFactory
type jsiiProxy_DynamoAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamoAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoAlarmFactory(alarmFactory AlarmFactory) DynamoAlarmFactory {
	_init_.Initialize()

	if err := validateNewDynamoAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoAlarmFactory_Override(d DynamoAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoAlarmFactory",
		[]interface{}{alarmFactory},
		d,
	)
}

func (d *jsiiProxy_DynamoAlarmFactory) AddConsumedCapacityAlarm(metric interface{}, capacityType CapacityType, props *ConsumedCapacityThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := d.validateAddConsumedCapacityAlarmParameters(metric, capacityType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		d,
		"addConsumedCapacityAlarm",
		[]interface{}{metric, capacityType, props, disambiguator},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoAlarmFactory) AddThrottledEventsAlarm(metric interface{}, capacityType CapacityType, props *ThrottledEventsThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := d.validateAddThrottledEventsAlarmParameters(metric, capacityType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		d,
		"addThrottledEventsAlarm",
		[]interface{}{metric, capacityType, props, disambiguator},
		&returns,
	)

	return returns
}

