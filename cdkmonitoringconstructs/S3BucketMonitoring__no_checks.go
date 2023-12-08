//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3BucketMonitoring) validateAddAlarmParameters(alarm *AlarmWithAnnotation) error {
	return nil
}

func (s *jsiiProxy_S3BucketMonitoring) validateCreateAlarmFactoryParameters(alarmNamePrefix *string) error {
	return nil
}

func (s *jsiiProxy_S3BucketMonitoring) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewS3BucketMonitoringParameters(scope MonitoringScope, props *S3BucketMonitoringProps) error {
	return nil
}

