package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type AgeAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddDaysSinceUpdateAlarm(metric interface{}, props *DaysSinceUpdateThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddDaysToExpiryAlarm(metric interface{}, props *DaysToExpiryThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddIteratorMaxAgeAlarm(metric interface{}, props *MaxAgeThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for AgeAlarmFactory
type jsiiProxy_AgeAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AgeAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAgeAlarmFactory(alarmFactory AlarmFactory) AgeAlarmFactory {
	_init_.Initialize()

	if err := validateNewAgeAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_AgeAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AgeAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewAgeAlarmFactory_Override(a AgeAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AgeAlarmFactory",
		[]interface{}{alarmFactory},
		a,
	)
}

func (a *jsiiProxy_AgeAlarmFactory) AddDaysSinceUpdateAlarm(metric interface{}, props *DaysSinceUpdateThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := a.validateAddDaysSinceUpdateAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addDaysSinceUpdateAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgeAlarmFactory) AddDaysToExpiryAlarm(metric interface{}, props *DaysToExpiryThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := a.validateAddDaysToExpiryAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addDaysToExpiryAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AgeAlarmFactory) AddIteratorMaxAgeAlarm(metric interface{}, props *MaxAgeThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := a.validateAddIteratorMaxAgeAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addIteratorMaxAgeAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

