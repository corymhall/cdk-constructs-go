//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SingleWidgetDashboardSegment) validateWidgetsForDashboardParameters(name *string) error {
	return nil
}

func validateNewSingleWidgetDashboardSegmentParameters(widget awscloudwatch.IWidget) error {
	return nil
}

