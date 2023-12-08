package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type TpsAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxTpsAlarm(metric interface{}, props *HighTpsThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinTpsAlarm(metric interface{}, props *LowTpsThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for TpsAlarmFactory
type jsiiProxy_TpsAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_TpsAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewTpsAlarmFactory(alarmFactory AlarmFactory) TpsAlarmFactory {
	_init_.Initialize()

	if err := validateNewTpsAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_TpsAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.TpsAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewTpsAlarmFactory_Override(t TpsAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.TpsAlarmFactory",
		[]interface{}{alarmFactory},
		t,
	)
}

func (t *jsiiProxy_TpsAlarmFactory) AddMaxTpsAlarm(metric interface{}, props *HighTpsThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMaxTpsAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMaxTpsAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TpsAlarmFactory) AddMinTpsAlarm(metric interface{}, props *LowTpsThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMinTpsAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMinTpsAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

