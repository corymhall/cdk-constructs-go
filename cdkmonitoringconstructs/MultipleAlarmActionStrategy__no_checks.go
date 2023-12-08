//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MultipleAlarmActionStrategy) validateAddAlarmActionsParameters(props *AlarmActionStrategyProps) error {
	return nil
}

func validateNewMultipleAlarmActionStrategyParameters(actions *[]IAlarmActionStrategy) error {
	return nil
}

