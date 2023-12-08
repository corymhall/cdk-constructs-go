package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type ElastiCacheAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxEvictedItemsCountAlarm(metric interface{}, props *MaxItemsCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxItemsCountAlarm(metric interface{}, props *MaxItemsCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxUsedSwapMemoryAlarm(metric interface{}, props *MaxUsedSwapMemoryThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinFreeableMemoryAlarm(metric interface{}, props *MinFreeableMemoryThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for ElastiCacheAlarmFactory
type jsiiProxy_ElastiCacheAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_ElastiCacheAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewElastiCacheAlarmFactory(alarmFactory AlarmFactory) ElastiCacheAlarmFactory {
	_init_.Initialize()

	if err := validateNewElastiCacheAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_ElastiCacheAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewElastiCacheAlarmFactory_Override(e ElastiCacheAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.ElastiCacheAlarmFactory",
		[]interface{}{alarmFactory},
		e,
	)
}

func (e *jsiiProxy_ElastiCacheAlarmFactory) AddMaxEvictedItemsCountAlarm(metric interface{}, props *MaxItemsCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddMaxEvictedItemsCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addMaxEvictedItemsCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheAlarmFactory) AddMaxItemsCountAlarm(metric interface{}, props *MaxItemsCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddMaxItemsCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addMaxItemsCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheAlarmFactory) AddMaxUsedSwapMemoryAlarm(metric interface{}, props *MaxUsedSwapMemoryThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddMaxUsedSwapMemoryAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addMaxUsedSwapMemoryAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ElastiCacheAlarmFactory) AddMinFreeableMemoryAlarm(metric interface{}, props *MinFreeableMemoryThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := e.validateAddMinFreeableMemoryAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		e,
		"addMinFreeableMemoryAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

