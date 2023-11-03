package cdkdatadog


type MonitorThresholds struct {
	// Threshold value for triggering an alert.
	Critical *float64 `field:"optional" json:"critical" yaml:"critical"`
	// Threshold value for recovering from an alert state.
	CriticalRecovery *float64 `field:"optional" json:"criticalRecovery" yaml:"criticalRecovery"`
	// Threshold value for recovering from an alert state.
	Ok *float64 `field:"optional" json:"ok" yaml:"ok"`
	// Threshold value for triggering a warning.
	Warning *float64 `field:"optional" json:"warning" yaml:"warning"`
	// Threshold value for recovering from a warning state.
	WarningRecovery *float64 `field:"optional" json:"warningRecovery" yaml:"warningRecovery"`
}

