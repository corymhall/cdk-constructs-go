package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type AuroraAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxServerlessDatabaseCapacity(metric interface{}, props *HighServerlessDatabaseCapacityThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for AuroraAlarmFactory
type jsiiProxy_AuroraAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AuroraAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAuroraAlarmFactory(alarmFactory AlarmFactory) AuroraAlarmFactory {
	_init_.Initialize()

	if err := validateNewAuroraAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_AuroraAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AuroraAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewAuroraAlarmFactory_Override(a AuroraAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AuroraAlarmFactory",
		[]interface{}{alarmFactory},
		a,
	)
}

func (a *jsiiProxy_AuroraAlarmFactory) AddMaxServerlessDatabaseCapacity(metric interface{}, props *HighServerlessDatabaseCapacityThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := a.validateAddMaxServerlessDatabaseCapacityParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addMaxServerlessDatabaseCapacity",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

