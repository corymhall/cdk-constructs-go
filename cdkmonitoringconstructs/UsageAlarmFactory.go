package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type UsageAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddMaxCpuUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string, usageType UsageType, additionalAlarmNameSuffix *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxDiskUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxFileDescriptorPercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxHeapMemoryAfterGCUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxMasterCpuUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxMasterMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxThreadCountUsageAlarm(percentMetric interface{}, props *UsageCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxUsageCountAlarm(metric interface{}, props *UsageCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, usageType UsageType, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinUsageCountAlarm(percentMetric interface{}, props *MinUsageCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for UsageAlarmFactory
type jsiiProxy_UsageAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_UsageAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewUsageAlarmFactory(alarmFactory AlarmFactory) UsageAlarmFactory {
	_init_.Initialize()

	if err := validateNewUsageAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_UsageAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.UsageAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewUsageAlarmFactory_Override(u UsageAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.UsageAlarmFactory",
		[]interface{}{alarmFactory},
		u,
	)
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxCpuUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string, usageType UsageType, additionalAlarmNameSuffix *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxCpuUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxCpuUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator, usageType, additionalAlarmNameSuffix},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxDiskUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxDiskUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxDiskUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxFileDescriptorPercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxFileDescriptorPercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxFileDescriptorPercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxHeapMemoryAfterGCUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxHeapMemoryAfterGCUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxHeapMemoryAfterGCUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxMasterCpuUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxMasterCpuUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxMasterCpuUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxMasterMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxMasterMemoryUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxMasterMemoryUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxMemoryUsagePercentAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxMemoryUsagePercentAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxThreadCountUsageAlarm(percentMetric interface{}, props *UsageCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxThreadCountUsageAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxThreadCountUsageAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMaxUsageCountAlarm(metric interface{}, props *UsageCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMaxUsageCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMaxUsageCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMemoryUsagePercentAlarm(percentMetric interface{}, props *UsageThreshold, usageType UsageType, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMemoryUsagePercentAlarmParameters(percentMetric, props, usageType); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMemoryUsagePercentAlarm",
		[]interface{}{percentMetric, props, usageType, disambiguator},
		&returns,
	)

	return returns
}

func (u *jsiiProxy_UsageAlarmFactory) AddMinUsageCountAlarm(percentMetric interface{}, props *MinUsageCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := u.validateAddMinUsageCountAlarmParameters(percentMetric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		u,
		"addMinUsageCountAlarm",
		[]interface{}{percentMetric, props, disambiguator},
		&returns,
	)

	return returns
}

