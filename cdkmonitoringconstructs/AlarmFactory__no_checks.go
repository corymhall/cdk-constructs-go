//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlarmFactory) validateAddAlarmParameters(metric interface{}, props *AddAlarmProps) error {
	return nil
}

func (a *jsiiProxy_AlarmFactory) validateAddCompositeAlarmParameters(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) error {
	return nil
}

func (a *jsiiProxy_AlarmFactory) validateCreateAnnotationParameters(props *AlarmAnnotationStrategyProps) error {
	return nil
}

func (a *jsiiProxy_AlarmFactory) validateDetermineCompositeAlarmRuleParameters(alarms *[]*AlarmWithAnnotation, props *AddCompositeAlarmProps) error {
	return nil
}

func (a *jsiiProxy_AlarmFactory) validateGenerateDescriptionParameters(alarmDescription *string) error {
	return nil
}

func validateNewAlarmFactoryParameters(alarmScope constructs.Construct, props *AlarmFactoryProps) error {
	return nil
}

