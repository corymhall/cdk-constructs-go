package cdkmonitoringconstructs


// Experimental.
type CompositeAlarmOperator string

const (
	// trigger only if all the alarms are triggered.
	// Experimental.
	CompositeAlarmOperator_AND CompositeAlarmOperator = "AND"
	// trigger if any of the alarms is triggered.
	// Experimental.
	CompositeAlarmOperator_OR CompositeAlarmOperator = "OR"
)

