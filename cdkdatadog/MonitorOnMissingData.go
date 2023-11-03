package cdkdatadog


// Controls how groups or monitors are treated if an evaluation does not return any data points.
//
// The default option results in different behavior depending on the monitor query type.
// For monitors using Count queries, an empty monitor evaluation is treated as 0 and is compared to the threshold conditions.
// For monitors using any query type other than Count, for example Gauge, Measure, or Rate, the monitor shows the last known status.
// This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.
type MonitorOnMissingData string

const (
	// default.
	MonitorOnMissingData_DEFAULT MonitorOnMissingData = "DEFAULT"
	// show_no_data.
	MonitorOnMissingData_SHOW_NO_DATA MonitorOnMissingData = "SHOW_NO_DATA"
	// show_and_notify_no_data.
	MonitorOnMissingData_SHOW_AND_NOTIFY_NO_DATA MonitorOnMissingData = "SHOW_AND_NOTIFY_NO_DATA"
	// resolve.
	MonitorOnMissingData_RESOLVE MonitorOnMissingData = "RESOLVE"
)

