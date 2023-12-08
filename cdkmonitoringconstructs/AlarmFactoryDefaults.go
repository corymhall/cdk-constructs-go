package cdkmonitoringconstructs


// Experimental.
type AlarmFactoryDefaults struct {
	// Enables the configured CloudWatch alarm ticketing actions for either all severities, or per severity.
	// Experimental.
	ActionsEnabled interface{} `field:"required" json:"actionsEnabled" yaml:"actionsEnabled"`
	// Global prefix for all alarm names.
	//
	// This should be something unique to avoid potential collisions.
	// This is ignored if an alarm's dedupeStringOverride is declared.
	// Experimental.
	AlarmNamePrefix *string `field:"required" json:"alarmNamePrefix" yaml:"alarmNamePrefix"`
	// Default alarm action used for each alarm, unless it is overridden.
	// Default: - no action.
	//
	// Experimental.
	Action IAlarmActionStrategy `field:"optional" json:"action" yaml:"action"`
	// Custom strategy to name alarms.
	// Default: - default behaviour (no change).
	//
	// Experimental.
	AlarmNamingStrategy IAlarmNamingStrategy `field:"optional" json:"alarmNamingStrategy" yaml:"alarmNamingStrategy"`
	// Custom strategy to create annotations for alarms.
	// Default: - default annotations.
	//
	// Experimental.
	AnnotationStrategy IAlarmAnnotationStrategy `field:"optional" json:"annotationStrategy" yaml:"annotationStrategy"`
	// Number of breaches required to transition into an ALARM state.
	// Default: - 3.
	//
	// Experimental.
	DatapointsToAlarm *float64 `field:"optional" json:"datapointsToAlarm" yaml:"datapointsToAlarm"`
	// Custom strategy to process dedupe strings of the alarms.
	// Default: - default behaviour (no change).
	//
	// Experimental.
	DedupeStringProcessor IAlarmDedupeStringProcessor `field:"optional" json:"dedupeStringProcessor" yaml:"dedupeStringProcessor"`
	// Optional alarm action for each disambiguator.
	// Default: - Global alarm action if defined.
	//
	// Experimental.
	DisambiguatorAction *map[string]IAlarmActionStrategy `field:"optional" json:"disambiguatorAction" yaml:"disambiguatorAction"`
	// An optional link included in the generated ticket description body.
	// Experimental.
	DocumentationLink *string `field:"optional" json:"documentationLink" yaml:"documentationLink"`
	// Number of periods to consider when checking the number of breaching datapoints.
	// Default: - Same as datapointsToAlarm.
	//
	// Experimental.
	EvaluationPeriods *float64 `field:"optional" json:"evaluationPeriods" yaml:"evaluationPeriods"`
	// An optional link included in the generated ticket description body.
	// Experimental.
	RunbookLink *string `field:"optional" json:"runbookLink" yaml:"runbookLink"`
	// If this is defined as false and dedupeStringOverride is undefined, the alarm prefix will be part of the dedupe string.
	//
	// This essentially stops the dedupe of different errors together.
	// Default: - undefined (true).
	//
	// Experimental.
	UseDefaultDedupeForError *bool `field:"optional" json:"useDefaultDedupeForError" yaml:"useDefaultDedupeForError"`
	// If this is defined as false and dedupeStringOverride is undefined, the alarm prefix will be part of the dedupe string.
	//
	// This essentially stops the dedupe of different latency issues together.
	// Default: - undefined (true).
	//
	// Experimental.
	UseDefaultDedupeForLatency *bool `field:"optional" json:"useDefaultDedupeForLatency" yaml:"useDefaultDedupeForLatency"`
}

