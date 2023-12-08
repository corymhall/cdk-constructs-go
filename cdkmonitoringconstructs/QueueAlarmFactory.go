package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type QueueAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxQueueIncomingMessagesCountAlarm(metric interface{}, props *MaxIncomingMessagesCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxQueueMessageAgeAlarm(metric interface{}, props *MaxMessageAgeThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxQueueMessageCountAlarm(metric interface{}, props *MaxMessageCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxQueueTimeToDrainMessagesAlarm(metric interface{}, props *MaxTimeToDrainThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinQueueIncomingMessagesCountAlarm(metric interface{}, props *MinIncomingMessagesCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinQueueMessageCountAlarm(metric interface{}, props *MinMessageCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for QueueAlarmFactory
type jsiiProxy_QueueAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_QueueAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewQueueAlarmFactory(alarmFactory AlarmFactory) QueueAlarmFactory {
	_init_.Initialize()

	if err := validateNewQueueAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_QueueAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.QueueAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewQueueAlarmFactory_Override(q QueueAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.QueueAlarmFactory",
		[]interface{}{alarmFactory},
		q,
	)
}

func (q *jsiiProxy_QueueAlarmFactory) AddMaxQueueIncomingMessagesCountAlarm(metric interface{}, props *MaxIncomingMessagesCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMaxQueueIncomingMessagesCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMaxQueueIncomingMessagesCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueAlarmFactory) AddMaxQueueMessageAgeAlarm(metric interface{}, props *MaxMessageAgeThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMaxQueueMessageAgeAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMaxQueueMessageAgeAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueAlarmFactory) AddMaxQueueMessageCountAlarm(metric interface{}, props *MaxMessageCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMaxQueueMessageCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMaxQueueMessageCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueAlarmFactory) AddMaxQueueTimeToDrainMessagesAlarm(metric interface{}, props *MaxTimeToDrainThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMaxQueueTimeToDrainMessagesAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMaxQueueTimeToDrainMessagesAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueAlarmFactory) AddMinQueueIncomingMessagesCountAlarm(metric interface{}, props *MinIncomingMessagesCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMinQueueIncomingMessagesCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMinQueueIncomingMessagesCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (q *jsiiProxy_QueueAlarmFactory) AddMinQueueMessageCountAlarm(metric interface{}, props *MinMessageCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := q.validateAddMinQueueMessageCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		q,
		"addMinQueueMessageCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

