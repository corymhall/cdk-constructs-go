package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ConnectionAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxConnectionCountAlarm(metric interface{}, props *HighConnectionCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinConnectionCountAlarm(metric interface{}, props *LowConnectionCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for ConnectionAlarmFactory
type jsiiProxy_ConnectionAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ConnectionAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewConnectionAlarmFactory(alarmFactory AlarmFactory) ConnectionAlarmFactory {
	_init_.Initialize()

	if err := validateNewConnectionAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_ConnectionAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ConnectionAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewConnectionAlarmFactory_Override(c ConnectionAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ConnectionAlarmFactory",
		[]interface{}{alarmFactory},
		c,
	)
}

func (c *jsiiProxy_ConnectionAlarmFactory) AddMaxConnectionCountAlarm(metric interface{}, props *HighConnectionCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := c.validateAddMaxConnectionCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"addMaxConnectionCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ConnectionAlarmFactory) AddMinConnectionCountAlarm(metric interface{}, props *LowConnectionCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := c.validateAddMinConnectionCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		c,
		"addMinConnectionCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

