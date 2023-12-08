package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type TopicAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxMessagesPublishedAlarm(metric interface{}, props *HighMessagesPublishedThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMessageNotificationsFailedAlarm(metric interface{}, props *NotificationsFailedThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinMessagesPublishedAlarm(metric interface{}, props *LowMessagesPublishedThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for TopicAlarmFactory
type jsiiProxy_TopicAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_TopicAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewTopicAlarmFactory(alarmFactory AlarmFactory) TopicAlarmFactory {
	_init_.Initialize()

	if err := validateNewTopicAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_TopicAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.TopicAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewTopicAlarmFactory_Override(t TopicAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.TopicAlarmFactory",
		[]interface{}{alarmFactory},
		t,
	)
}

func (t *jsiiProxy_TopicAlarmFactory) AddMaxMessagesPublishedAlarm(metric interface{}, props *HighMessagesPublishedThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMaxMessagesPublishedAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMaxMessagesPublishedAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicAlarmFactory) AddMessageNotificationsFailedAlarm(metric interface{}, props *NotificationsFailedThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMessageNotificationsFailedAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMessageNotificationsFailedAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TopicAlarmFactory) AddMinMessagesPublishedAlarm(metric interface{}, props *LowMessagesPublishedThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := t.validateAddMinMessagesPublishedAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		t,
		"addMinMessagesPublishedAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

