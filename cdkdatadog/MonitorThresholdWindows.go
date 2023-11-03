package cdkdatadog


type MonitorThresholdWindows struct {
	// How long an anomalous metric must be normal before recovering from an alert state.
	RecoveryWindow *string `field:"optional" json:"recoveryWindow" yaml:"recoveryWindow"`
	// How long a metric must be anomalous before triggering an alert.
	TriggerWindow *string `field:"optional" json:"triggerWindow" yaml:"triggerWindow"`
}

