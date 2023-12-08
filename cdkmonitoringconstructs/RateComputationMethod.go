package cdkmonitoringconstructs


// Enumeration of different rate computation methods.
// Experimental.
type RateComputationMethod string

const (
	// Number of occurrences relative to requests.
	//
	// Less sensitive than per-second when TPS > 1.
	// Experimental.
	RateComputationMethod_AVERAGE RateComputationMethod = "AVERAGE"
	// Number of occurrences per second.
	//
	// More sensitive than average when TPS > 1.
	// Experimental.
	RateComputationMethod_PER_SECOND RateComputationMethod = "PER_SECOND"
	// Number of occurrences per minute.
	// Experimental.
	RateComputationMethod_PER_MINUTE RateComputationMethod = "PER_MINUTE"
	// Number of occurrences per hour.
	// Experimental.
	RateComputationMethod_PER_HOUR RateComputationMethod = "PER_HOUR"
	// Number of occurrences per day.
	// Experimental.
	RateComputationMethod_PER_DAY RateComputationMethod = "PER_DAY"
)

