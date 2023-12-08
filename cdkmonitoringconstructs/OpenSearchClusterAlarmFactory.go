package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type OpenSearchClusterAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddAutomatedSnapshotFailureAlarm(metric interface{}, props *OpenSearchClusterAutomatedSnapshotFailureThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddClusterIndexWritesBlockedAlarm(metric interface{}, props *OpenSearchClusterIndexWritesBlockedThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddClusterNodeCountAlarm(metric interface{}, props *OpenSearchClusterNodesThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddClusterStatusAlarm(metric interface{}, props *OpenSearchClusterStatusCustomization, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddKmsKeyErrorAlarm(metric interface{}, props *OpenSearchKmsKeyErrorThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddKmsKeyInaccessibleAlarm(metric interface{}, props *OpenSearchKmsKeyInaccessibleThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for OpenSearchClusterAlarmFactory
type jsiiProxy_OpenSearchClusterAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_OpenSearchClusterAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewOpenSearchClusterAlarmFactory(alarmFactory AlarmFactory) OpenSearchClusterAlarmFactory {
	_init_.Initialize()

	if err := validateNewOpenSearchClusterAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpenSearchClusterAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewOpenSearchClusterAlarmFactory_Override(o OpenSearchClusterAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.OpenSearchClusterAlarmFactory",
		[]interface{}{alarmFactory},
		o,
	)
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddAutomatedSnapshotFailureAlarm(metric interface{}, props *OpenSearchClusterAutomatedSnapshotFailureThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddAutomatedSnapshotFailureAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addAutomatedSnapshotFailureAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddClusterIndexWritesBlockedAlarm(metric interface{}, props *OpenSearchClusterIndexWritesBlockedThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddClusterIndexWritesBlockedAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addClusterIndexWritesBlockedAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddClusterNodeCountAlarm(metric interface{}, props *OpenSearchClusterNodesThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddClusterNodeCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addClusterNodeCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddClusterStatusAlarm(metric interface{}, props *OpenSearchClusterStatusCustomization, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddClusterStatusAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addClusterStatusAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddKmsKeyErrorAlarm(metric interface{}, props *OpenSearchKmsKeyErrorThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddKmsKeyErrorAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addKmsKeyErrorAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpenSearchClusterAlarmFactory) AddKmsKeyInaccessibleAlarm(metric interface{}, props *OpenSearchKmsKeyInaccessibleThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := o.validateAddKmsKeyInaccessibleAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		o,
		"addKmsKeyInaccessibleAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

