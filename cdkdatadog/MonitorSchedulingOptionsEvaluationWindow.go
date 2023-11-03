package cdkdatadog


// Configuration options for the evaluation window.
//
// If `hour_starts` is set, no other fields may be set. Otherwise, `day_starts` and `month_starts` must be set together.
type MonitorSchedulingOptionsEvaluationWindow struct {
	// The time of the day at which a one day cumulative evaluation window starts.
	//
	// Must be defined in UTC time in `HH:mm` format.
	DayStarts *string `field:"optional" json:"dayStarts" yaml:"dayStarts"`
	// The minute of the hour at which a one hour cumulative evaluation window starts.
	HourStarts *float64 `field:"optional" json:"hourStarts" yaml:"hourStarts"`
	// The day of the month at which a one month cumulative evaluation window starts.
	MonthStarts *float64 `field:"optional" json:"monthStarts" yaml:"monthStarts"`
}

