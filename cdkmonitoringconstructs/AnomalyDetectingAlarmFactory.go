package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type AnomalyDetectingAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddAlarmWhenOutOfBand(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *AnomalyDetectionThreshold) *AlarmWithAnnotation
}

// The jsii proxy struct for AnomalyDetectingAlarmFactory
type jsiiProxy_AnomalyDetectingAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_AnomalyDetectingAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewAnomalyDetectingAlarmFactory(alarmFactory AlarmFactory) AnomalyDetectingAlarmFactory {
	_init_.Initialize()

	if err := validateNewAnomalyDetectingAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_AnomalyDetectingAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AnomalyDetectingAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewAnomalyDetectingAlarmFactory_Override(a AnomalyDetectingAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AnomalyDetectingAlarmFactory",
		[]interface{}{alarmFactory},
		a,
	)
}

func (a *jsiiProxy_AnomalyDetectingAlarmFactory) AddAlarmWhenOutOfBand(metric interface{}, alarmNameSuffix *string, disambiguator *string, props *AnomalyDetectionThreshold) *AlarmWithAnnotation {
	if err := a.validateAddAlarmWhenOutOfBandParameters(metric, alarmNameSuffix, disambiguator, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		a,
		"addAlarmWhenOutOfBand",
		[]interface{}{metric, alarmNameSuffix, disambiguator, props},
		&returns,
	)

	return returns
}

