//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitoringScope) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func validateMonitoringScope_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewMonitoringScopeParameters(scope constructs.Construct, id *string) error {
	return nil
}

