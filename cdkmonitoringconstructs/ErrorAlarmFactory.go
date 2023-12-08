package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ErrorAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddErrorCountAlarm(metric interface{}, errorType ErrorType, props *ErrorCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddErrorRateAlarm(metric interface{}, errorType ErrorType, props *ErrorRateThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for ErrorAlarmFactory
type jsiiProxy_ErrorAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ErrorAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewErrorAlarmFactory(alarmFactory AlarmFactory) ErrorAlarmFactory {
	_init_.Initialize()

	if err := validateNewErrorAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_ErrorAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ErrorAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewErrorAlarmFactory_Override(e ErrorAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ErrorAlarmFactory",
		[]interface{}{alarmFactory},
		e,
	)
}

func (e *jsiiProxy_ErrorAlarmFactory) AddErrorCountAlarm(metric interface{}, errorType ErrorType, props *ErrorCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddErrorCountAlarmParameters(metric, errorType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addErrorCountAlarm",
		[]interface{}{metric, errorType, props, disambiguator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ErrorAlarmFactory) AddErrorRateAlarm(metric interface{}, errorType ErrorType, props *ErrorRateThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddErrorRateAlarmParameters(metric, errorType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addErrorRateAlarm",
		[]interface{}{metric, errorType, props, disambiguator},
		&returns,
	)

	return returns
}

