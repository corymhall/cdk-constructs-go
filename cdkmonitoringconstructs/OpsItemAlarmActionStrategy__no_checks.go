//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OpsItemAlarmActionStrategy) validateAddAlarmActionsParameters(props *AlarmActionStrategyProps) error {
	return nil
}

func validateNewOpsItemAlarmActionStrategyParameters(severity awscloudwatchactions.OpsItemSeverity) error {
	return nil
}

