package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type KinesisDataAnalyticsAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddDowntimeAlarm(metric interface{}, props *MaxDowntimeThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddFullRestartAlarm(metric interface{}, props *FullRestartCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for KinesisDataAnalyticsAlarmFactory
type jsiiProxy_KinesisDataAnalyticsAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_KinesisDataAnalyticsAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisDataAnalyticsAlarmFactory(alarmFactory AlarmFactory) KinesisDataAnalyticsAlarmFactory {
	_init_.Initialize()

	if err := validateNewKinesisDataAnalyticsAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisDataAnalyticsAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisDataAnalyticsAlarmFactory_Override(k KinesisDataAnalyticsAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisDataAnalyticsAlarmFactory",
		[]interface{}{alarmFactory},
		k,
	)
}

func (k *jsiiProxy_KinesisDataAnalyticsAlarmFactory) AddDowntimeAlarm(metric interface{}, props *MaxDowntimeThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddDowntimeAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addDowntimeAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisDataAnalyticsAlarmFactory) AddFullRestartAlarm(metric interface{}, props *FullRestartCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddFullRestartAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addFullRestartAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

