package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type CustomAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddCustomAlarm(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *CustomThreshold) *AlarmWithAnnotation
}

// The jsii proxy struct for CustomAlarmFactory
type jsiiProxy_CustomAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_CustomAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomAlarmFactory(alarmFactory AlarmFactory) CustomAlarmFactory {
	_init_.Initialize()

	if err := validateNewCustomAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomAlarmFactory_Override(c CustomAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomAlarmFactory",
		[]interface{}{alarmFactory},
		c,
	)
}

func (c *jsiiProxy_CustomAlarmFactory) AddCustomAlarm(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *CustomThreshold) *AlarmWithAnnotation {
	if err := c.validateAddCustomAlarmParameters(metric, alarmNameSuffix, disambiguator, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"addCustomAlarm",
		[]interface{}{metric, alarmNameSuffix, disambiguator, props},
		&returns,
	)

	return returns
}

