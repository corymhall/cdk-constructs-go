package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type LogLevelAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddLogCountAlarm(metric interface{}, logLevel LogLevel, props *LogLevelCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for LogLevelAlarmFactory
type jsiiProxy_LogLevelAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_LogLevelAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewLogLevelAlarmFactory(alarmFactory AlarmFactory) LogLevelAlarmFactory {
	_init_.Initialize()

	if err := validateNewLogLevelAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogLevelAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LogLevelAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewLogLevelAlarmFactory_Override(l LogLevelAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LogLevelAlarmFactory",
		[]interface{}{alarmFactory},
		l,
	)
}

func (l *jsiiProxy_LogLevelAlarmFactory) AddLogCountAlarm(metric interface{}, logLevel LogLevel, props *LogLevelCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := l.validateAddLogCountAlarmParameters(metric, logLevel, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"addLogCountAlarm",
		[]interface{}{metric, logLevel, props, disambiguator},
		&returns,
	)

	return returns
}

