package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type TaskHealthAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddAvailabilityAlarm(metric interface{}, props *AvailabilityThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddHealthyTaskCountAlarm(metric interface{}, props *HealthyTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddHealthyTaskPercentAlarm(metric interface{}, props *HealthyTaskPercentThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinRunningTaskCountAlarm(metric interface{}, props *MinRunningTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddRunningTaskCountAlarm(metric interface{}, props *RunningTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddRunningTaskRateAlarm(metric interface{}, props *RunningTaskRateThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddUnhealthyTaskCountAlarm(metric interface{}, props *UnhealthyTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for TaskHealthAlarmFactory
type jsiiProxy_TaskHealthAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_TaskHealthAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewTaskHealthAlarmFactory(alarmFactory AlarmFactory) TaskHealthAlarmFactory {
	_init_.Initialize()

	if err := validateNewTaskHealthAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaskHealthAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.TaskHealthAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewTaskHealthAlarmFactory_Override(t TaskHealthAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.TaskHealthAlarmFactory",
		[]interface{}{alarmFactory},
		t,
	)
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddAvailabilityAlarm(metric interface{}, props *AvailabilityThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddAvailabilityAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addAvailabilityAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddHealthyTaskCountAlarm(metric interface{}, props *HealthyTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddHealthyTaskCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addHealthyTaskCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddHealthyTaskPercentAlarm(metric interface{}, props *HealthyTaskPercentThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddHealthyTaskPercentAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addHealthyTaskPercentAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddMinRunningTaskCountAlarm(metric interface{}, props *MinRunningTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMinRunningTaskCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMinRunningTaskCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddRunningTaskCountAlarm(metric interface{}, props *RunningTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddRunningTaskCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addRunningTaskCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddRunningTaskRateAlarm(metric interface{}, props *RunningTaskRateThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddRunningTaskRateAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addRunningTaskRateAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TaskHealthAlarmFactory) AddUnhealthyTaskCountAlarm(metric interface{}, props *UnhealthyTaskCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddUnhealthyTaskCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addUnhealthyTaskCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

