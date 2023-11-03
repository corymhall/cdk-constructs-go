package cdkdatadog


// Toggles the display of additional content sent in the monitor notification.
type MonitorNotificationPresetName string

const (
	// show_all.
	MonitorNotificationPresetName_SHOW_ALL MonitorNotificationPresetName = "SHOW_ALL"
	// hide_query.
	MonitorNotificationPresetName_HIDE_QUERY MonitorNotificationPresetName = "HIDE_QUERY"
	// hide_handles.
	MonitorNotificationPresetName_HIDE_HANDLES MonitorNotificationPresetName = "HIDE_HANDLES"
	// hide_all.
	MonitorNotificationPresetName_HIDE_ALL MonitorNotificationPresetName = "HIDE_ALL"
)

