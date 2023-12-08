//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateResolveProjectNameParameters(project awscodebuild.IProject) error {
	return nil
}

func (c *jsiiProxy_CodeBuildProjectMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewCodeBuildProjectMonitoringParameters(scope MonitoringScope, props *CodeBuildProjectMonitoringProps) error {
	return nil
}

