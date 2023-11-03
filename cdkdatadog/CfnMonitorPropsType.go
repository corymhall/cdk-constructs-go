package cdkdatadog


// The type of the monitor.
type CfnMonitorPropsType string

const (
	// audit alert.
	CfnMonitorPropsType_AUDIT_ALERT CfnMonitorPropsType = "AUDIT_ALERT"
	// composite.
	CfnMonitorPropsType_COMPOSITE CfnMonitorPropsType = "COMPOSITE"
	// event alert.
	CfnMonitorPropsType_EVENT_ALERT CfnMonitorPropsType = "EVENT_ALERT"
	// event-v2 alert.
	CfnMonitorPropsType_EVENT_V2_ALERT CfnMonitorPropsType = "EVENT_V2_ALERT"
	// log alert.
	CfnMonitorPropsType_LOG_ALERT CfnMonitorPropsType = "LOG_ALERT"
	// metric alert.
	CfnMonitorPropsType_METRIC_ALERT CfnMonitorPropsType = "METRIC_ALERT"
	// process alert.
	CfnMonitorPropsType_PROCESS_ALERT CfnMonitorPropsType = "PROCESS_ALERT"
	// query alert.
	CfnMonitorPropsType_QUERY_ALERT CfnMonitorPropsType = "QUERY_ALERT"
	// service check.
	CfnMonitorPropsType_SERVICE_CHECK CfnMonitorPropsType = "SERVICE_CHECK"
	// synthetics alert.
	CfnMonitorPropsType_SYNTHETICS_ALERT CfnMonitorPropsType = "SYNTHETICS_ALERT"
	// trace-analytics alert.
	CfnMonitorPropsType_TRACE_ANALYTICS_ALERT CfnMonitorPropsType = "TRACE_ANALYTICS_ALERT"
	// slo alert.
	CfnMonitorPropsType_SLO_ALERT CfnMonitorPropsType = "SLO_ALERT"
	// rum alert.
	CfnMonitorPropsType_RUM_ALERT CfnMonitorPropsType = "RUM_ALERT"
	// ci-pipelines alert.
	CfnMonitorPropsType_CI_PIPELINES_ALERT CfnMonitorPropsType = "CI_PIPELINES_ALERT"
	// error-tracking alert.
	CfnMonitorPropsType_ERROR_TRACKING_ALERT CfnMonitorPropsType = "ERROR_TRACKING_ALERT"
	// ci-tests alert.
	CfnMonitorPropsType_CI_TESTS_ALERT CfnMonitorPropsType = "CI_TESTS_ALERT"
)

