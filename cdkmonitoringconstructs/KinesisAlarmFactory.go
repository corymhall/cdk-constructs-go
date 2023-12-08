package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type KinesisAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddFirehoseStreamExceedSafetyThresholdAlarm(metric interface{}, metricName *string, quotaName *string, props *FirehoseStreamLimitThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddIteratorMaxAgeAlarm(metric interface{}, props *MaxIteratorAgeThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddProvisionedReadThroughputExceededAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddProvisionedWriteThroughputExceededAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddPutRecordsFailedAlarm(metric interface{}, props *RecordsFailedThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddPutRecordsThrottledAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for KinesisAlarmFactory
type jsiiProxy_KinesisAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_KinesisAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewKinesisAlarmFactory(alarmFactory AlarmFactory) KinesisAlarmFactory {
	_init_.Initialize()

	if err := validateNewKinesisAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_KinesisAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewKinesisAlarmFactory_Override(k KinesisAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.KinesisAlarmFactory",
		[]interface{}{alarmFactory},
		k,
	)
}

func (k *jsiiProxy_KinesisAlarmFactory) AddFirehoseStreamExceedSafetyThresholdAlarm(metric interface{}, metricName *string, quotaName *string, props *FirehoseStreamLimitThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddFirehoseStreamExceedSafetyThresholdAlarmParameters(metric, metricName, quotaName, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addFirehoseStreamExceedSafetyThresholdAlarm",
		[]interface{}{metric, metricName, quotaName, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAlarmFactory) AddIteratorMaxAgeAlarm(metric interface{}, props *MaxIteratorAgeThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddIteratorMaxAgeAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addIteratorMaxAgeAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAlarmFactory) AddProvisionedReadThroughputExceededAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddProvisionedReadThroughputExceededAlarmParameters(metric, props, disambiguator); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addProvisionedReadThroughputExceededAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAlarmFactory) AddProvisionedWriteThroughputExceededAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddProvisionedWriteThroughputExceededAlarmParameters(metric, props, disambiguator); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addProvisionedWriteThroughputExceededAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAlarmFactory) AddPutRecordsFailedAlarm(metric interface{}, props *RecordsFailedThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddPutRecordsFailedAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addPutRecordsFailedAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KinesisAlarmFactory) AddPutRecordsThrottledAlarm(metric interface{}, props *RecordsThrottledThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := k.validateAddPutRecordsThrottledAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		k,
		"addPutRecordsThrottledAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

