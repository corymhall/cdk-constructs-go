//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SecretsManagerAlarmFactory) validateAddChangeInSecretCountAlarmParameters(metric interface{}, props *ChangeInSecretCountThreshold) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerAlarmFactory) validateAddMaxSecretCountAlarmParameters(metric interface{}, props *MaxSecretCountThreshold) error {
	return nil
}

func (s *jsiiProxy_SecretsManagerAlarmFactory) validateAddMinSecretCountAlarmParameters(metric interface{}, props *MinSecretCountThreshold) error {
	return nil
}

func validateNewSecretsManagerAlarmFactoryParameters(alarmFactory AlarmFactory) error {
	return nil
}

