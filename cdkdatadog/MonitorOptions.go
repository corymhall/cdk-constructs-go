package cdkdatadog


type MonitorOptions struct {
	// Whether or not to include a sample of the logs.
	EnableLogsSample *bool `field:"optional" json:"enableLogsSample" yaml:"enableLogsSample"`
	// Whether or not to send a list of samples when the monitor triggers.
	//
	// This is only used by CI Test and Pipeline monitors.
	EnableSamples *bool `field:"optional" json:"enableSamples" yaml:"enableSamples"`
	// Message to include with a re-notification when renotify_interval is set.
	EscalationMessage *string `field:"optional" json:"escalationMessage" yaml:"escalationMessage"`
	// Time in seconds to delay evaluation.
	EvaluationDelay *float64 `field:"optional" json:"evaluationDelay" yaml:"evaluationDelay"`
	// The time span after which groups with missing data are dropped from the monitor state.
	//
	// The minimum value is one hour, and the maximum value is 72 hours.
	// Example values are: "60m", "1h", and "2d".
	// This option is only available for APM Trace Analytics, Audit Trail, CI, Error Tracking, Event, Logs, and RUM monitors.
	GroupRetentionDuration *string `field:"optional" json:"groupRetentionDuration" yaml:"groupRetentionDuration"`
	// Whether or not to include triggering tags into notification title.
	IncludeTags *bool `field:"optional" json:"includeTags" yaml:"includeTags"`
	// Whether or not changes to this monitor should be restricted to the creator or admins.
	Locked *bool `field:"optional" json:"locked" yaml:"locked"`
	// How long the test should be in failure before alerting (integer, number of seconds, max 7200).
	MinFailureDuration *float64 `field:"optional" json:"minFailureDuration" yaml:"minFailureDuration"`
	// Number of locations allowed to fail before triggering alert.
	MinLocationFailed *float64 `field:"optional" json:"minLocationFailed" yaml:"minLocationFailed"`
	// Time (in seconds) to skip evaluations for new groups.
	//
	// For example, this option can be used to skip evaluations for new hosts while they initialize. Must be a non negative integer.
	NewGroupDelay *float64 `field:"optional" json:"newGroupDelay" yaml:"newGroupDelay"`
	// Time in seconds to allow a host to start reporting data before starting the evaluation of monitor results.
	NewHostDelay *float64 `field:"optional" json:"newHostDelay" yaml:"newHostDelay"`
	// Number of minutes data stopped reporting before notifying.
	NoDataTimeframe *float64 `field:"optional" json:"noDataTimeframe" yaml:"noDataTimeframe"`
	NotificationPresetName MonitorNotificationPresetName `field:"optional" json:"notificationPresetName" yaml:"notificationPresetName"`
	// Whether or not to notify tagged users when changes are made to the monitor.
	NotifyAudit *bool `field:"optional" json:"notifyAudit" yaml:"notifyAudit"`
	// Controls what granularity a monitor alerts on.
	//
	// Only available for monitors with groupings.
	// For instance, a monitor grouped by `cluster`, `namespace`, and `pod` can be configured to only notify on each new `cluster` violating the alert conditions by setting `notify_by` to `["cluster"]`.
	// Tags mentioned in `notify_by` must be a subset of the grouping tags in the query.
	// For example, a query grouped by `cluster` and `namespace` cannot notify on `region`.
	// Setting `notify_by` to `[*]` configures the monitor to notify as a simple-alert.
	NotifyBy *[]*string `field:"optional" json:"notifyBy" yaml:"notifyBy"`
	// Whether or not to notify when data stops reporting.
	NotifyNoData *bool `field:"optional" json:"notifyNoData" yaml:"notifyNoData"`
	OnMissingData MonitorOnMissingData `field:"optional" json:"onMissingData" yaml:"onMissingData"`
	// Number of minutes after the last notification before the monitor re-notifies on the current status.
	RenotifyInterval *float64 `field:"optional" json:"renotifyInterval" yaml:"renotifyInterval"`
	// The number of times re-notification messages should be sent on the current status at the provided re-notification interval.
	RenotifyOccurrences *float64 `field:"optional" json:"renotifyOccurrences" yaml:"renotifyOccurrences"`
	// The types of monitor statuses for which re-notification messages are sent.
	RenotifyStatuses *[]MonitorOptionsRenotifyStatuses `field:"optional" json:"renotifyStatuses" yaml:"renotifyStatuses"`
	// Whether or not the monitor requires a full window of data before it is evaluated.
	RequireFullWindow *bool `field:"optional" json:"requireFullWindow" yaml:"requireFullWindow"`
	SchedulingOptions *MonitorSchedulingOptions `field:"optional" json:"schedulingOptions" yaml:"schedulingOptions"`
	// ID of the corresponding synthetics check.
	SyntheticsCheckId *float64 `field:"optional" json:"syntheticsCheckId" yaml:"syntheticsCheckId"`
	// The threshold definitions.
	Thresholds *MonitorThresholds `field:"optional" json:"thresholds" yaml:"thresholds"`
	// The threshold window definitions.
	ThresholdWindows *MonitorThresholdWindows `field:"optional" json:"thresholdWindows" yaml:"thresholdWindows"`
	// Number of hours of the monitor not reporting data before it automatically resolves.
	TimeoutH *float64 `field:"optional" json:"timeoutH" yaml:"timeoutH"`
	// List of requests that can be used in the monitor query.
	Variables *[]interface{} `field:"optional" json:"variables" yaml:"variables"`
}

