package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ThroughputAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMinProcessedBytesAlarm(metric interface{}, props *MinProcessedBytesThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for ThroughputAlarmFactory
type jsiiProxy_ThroughputAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ThroughputAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewThroughputAlarmFactory(alarmFactory AlarmFactory) ThroughputAlarmFactory {
	_init_.Initialize()

	if err := validateNewThroughputAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_ThroughputAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ThroughputAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewThroughputAlarmFactory_Override(t ThroughputAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ThroughputAlarmFactory",
		[]interface{}{alarmFactory},
		t,
	)
}

func (t *jsiiProxy_ThroughputAlarmFactory) AddMinProcessedBytesAlarm(metric interface{}, props *MinProcessedBytesThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMinProcessedBytesAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMinProcessedBytesAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

