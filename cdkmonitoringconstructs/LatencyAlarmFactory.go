package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type LatencyAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddDurationAlarm(metric interface{}, latencyType LatencyType, props *DurationThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation
	// Experimental.
	AddIntegrationLatencyAlarm(metric interface{}, latencyType LatencyType, props *LatencyThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation
	// Experimental.
	AddJvmGarbageCollectionDurationAlarm(metric interface{}, latencyType LatencyType, props *DurationThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation
	// Experimental.
	AddLatencyAlarm(metric interface{}, latencyType LatencyType, props *LatencyThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation
}

// The jsii proxy struct for LatencyAlarmFactory
type jsiiProxy_LatencyAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_LatencyAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewLatencyAlarmFactory(alarmFactory AlarmFactory) LatencyAlarmFactory {
	_init_.Initialize()

	if err := validateNewLatencyAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_LatencyAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.LatencyAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewLatencyAlarmFactory_Override(l LatencyAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.LatencyAlarmFactory",
		[]interface{}{alarmFactory},
		l,
	)
}

func (l *jsiiProxy_LatencyAlarmFactory) AddDurationAlarm(metric interface{}, latencyType LatencyType, props *DurationThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation {
	if err := l.validateAddDurationAlarmParameters(metric, latencyType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"addDurationAlarm",
		[]interface{}{metric, latencyType, props, disambiguator, additionalAlarmNameSuffix},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LatencyAlarmFactory) AddIntegrationLatencyAlarm(metric interface{}, latencyType LatencyType, props *LatencyThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation {
	if err := l.validateAddIntegrationLatencyAlarmParameters(metric, latencyType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"addIntegrationLatencyAlarm",
		[]interface{}{metric, latencyType, props, disambiguator, additionalAlarmNameSuffix},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LatencyAlarmFactory) AddJvmGarbageCollectionDurationAlarm(metric interface{}, latencyType LatencyType, props *DurationThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation {
	if err := l.validateAddJvmGarbageCollectionDurationAlarmParameters(metric, latencyType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"addJvmGarbageCollectionDurationAlarm",
		[]interface{}{metric, latencyType, props, disambiguator, additionalAlarmNameSuffix},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LatencyAlarmFactory) AddLatencyAlarm(metric interface{}, latencyType LatencyType, props *LatencyThreshold, disambiguator *string, additionalAlarmNameSuffix *string) *AlarmWithAnnotation {
	if err := l.validateAddLatencyAlarmParameters(metric, latencyType, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		l,
		"addLatencyAlarm",
		[]interface{}{metric, latencyType, props, disambiguator, additionalAlarmNameSuffix},
		&returns,
	)

	return returns
}

