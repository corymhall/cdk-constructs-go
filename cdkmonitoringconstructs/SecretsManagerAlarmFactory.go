package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type SecretsManagerAlarmFactory interface {
	// Experimental.
	AlarmFactory() AlarmFactory
	// Experimental.
	AddChangeInSecretCountAlarm(metric interface{}, props *ChangeInSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMaxSecretCountAlarm(metric interface{}, props *MaxSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation
	// Experimental.
	AddMinSecretCountAlarm(metric interface{}, props *MinSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation
}

// The jsii proxy struct for SecretsManagerAlarmFactory
type jsiiProxy_SecretsManagerAlarmFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SecretsManagerAlarmFactory) AlarmFactory() AlarmFactory {
	var returns AlarmFactory
	_jsii_.Get(
		j,
		"alarmFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewSecretsManagerAlarmFactory(alarmFactory AlarmFactory) SecretsManagerAlarmFactory {
	_init_.Initialize()

	if err := validateNewSecretsManagerAlarmFactoryParameters(alarmFactory); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecretsManagerAlarmFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerAlarmFactory",
		[]interface{}{alarmFactory},
		&j,
	)

	return &j
}

// Experimental.
func NewSecretsManagerAlarmFactory_Override(s SecretsManagerAlarmFactory, alarmFactory AlarmFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SecretsManagerAlarmFactory",
		[]interface{}{alarmFactory},
		s,
	)
}

func (s *jsiiProxy_SecretsManagerAlarmFactory) AddChangeInSecretCountAlarm(metric interface{}, props *ChangeInSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := s.validateAddChangeInSecretCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"addChangeInSecretCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerAlarmFactory) AddMaxSecretCountAlarm(metric interface{}, props *MaxSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := s.validateAddMaxSecretCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"addMaxSecretCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecretsManagerAlarmFactory) AddMinSecretCountAlarm(metric interface{}, props *MinSecretCountThreshold, disambiguator *string) *AlarmWithAnnotation {
	if err := s.validateAddMinSecretCountAlarmParameters(metric, props); err != nil {
		panic(err)
	}
	var returns *AlarmWithAnnotation

	_jsii_.Invoke(
		s,
		"addMinSecretCountAlarm",
		[]interface{}{metric, props, disambiguator},
		&returns,
	)

	return returns
}

