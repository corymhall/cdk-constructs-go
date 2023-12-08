//go:build no_runtime_type_checking

package cdkmonitoringconstructs

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MonitoringAspect) validateVisitParameters(node constructs.IConstruct) error {
	return nil
}

func validateNewMonitoringAspectParameters(monitoringFacade MonitoringFacade, props *MonitoringAspectProps) error {
	return nil
}

