package cdkmonitoringconstructs


// Properties necessary to create a composite alarm and configure it.
// Experimental.
type AddCompositeAlarmProps struct {
	// Disambiguator is a string that differentiates this alarm from other similar ones.
	// Experimental.
	Disambiguator *string `field:"required" json:"disambiguator" yaml:"disambiguator"`
	// Allows to override the default action strategy.
	// Default: - default action will be used.
	//
	// Experimental.
	ActionOverride IAlarmActionStrategy `field:"optional" json:"actionOverride" yaml:"actionOverride"`
	// Enables the configured CloudWatch alarm ticketing actions.
	// Default: - the same as monitoring facade default.
	//
	// Experimental.
	ActionsEnabled *bool `field:"optional" json:"actionsEnabled" yaml:"actionsEnabled"`
	// If this is defined, the default resource-specific alarm dedupe string will be set and this will be added as a suffix.
	//
	// This allows you to specify the same dedupe string for a family of alarms.
	// Cannot be defined at the same time as alarmDedupeStringOverride.
	// Default: - undefined (no suffix).
	//
	// Experimental.
	AlarmDedupeStringSuffix *string `field:"optional" json:"alarmDedupeStringSuffix" yaml:"alarmDedupeStringSuffix"`
	// Alarm description is included in the ticket and therefore should describe what happened, with as much context as possible.
	// Default: - no description.
	//
	// Experimental.
	AlarmDescription *string `field:"optional" json:"alarmDescription" yaml:"alarmDescription"`
	// A text included in the generated ticket description body, which fully replaces the generated text.
	// Default: - default auto-generated content only.
	//
	// Experimental.
	AlarmDescriptionOverride *string `field:"optional" json:"alarmDescriptionOverride" yaml:"alarmDescriptionOverride"`
	// If this is defined, the alarm name is set to this exact value.
	//
	// Please be aware that you need to specify prefix for different stages (Beta, Prod...) and realms (EU, NA...) manually.
	// Experimental.
	AlarmNameOverride *string `field:"optional" json:"alarmNameOverride" yaml:"alarmNameOverride"`
	// Suffix added to base alarm name.
	//
	// Alarm names need to be unique.
	// Default: - no suffix.
	//
	// Experimental.
	AlarmNameSuffix *string `field:"optional" json:"alarmNameSuffix" yaml:"alarmNameSuffix"`
	// Logical operator used to aggregate the status individual alarms.
	// Default: - OR.
	//
	// Experimental.
	CompositeOperator CompositeAlarmOperator `field:"optional" json:"compositeOperator" yaml:"compositeOperator"`
	// This allows user to attach custom parameters to this alarm, which can later be accessed from the "useCreatedAlarms" method.
	// Default: - no parameters.
	//
	// Experimental.
	CustomParams *map[string]interface{} `field:"optional" json:"customParams" yaml:"customParams"`
	// This allows user to attach custom values to this alarm, which can later be accessed from the "useCreatedAlarms" method.
	// Default: - no tags.
	//
	// Experimental.
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// If this is defined, the alarm dedupe string is set to this exact value.
	//
	// Please be aware that you need to handle deduping for different stages (Beta, Prod...) and realms (EU, NA...) manually.
	// Default: - undefined (no override).
	//
	// Experimental.
	DedupeStringOverride *string `field:"optional" json:"dedupeStringOverride" yaml:"dedupeStringOverride"`
	// An optional link included in the generated ticket description body.
	// Default: - no additional link will be added.
	//
	// Experimental.
	DocumentationLink *string `field:"optional" json:"documentationLink" yaml:"documentationLink"`
	// Indicates whether the alarming range of values should be highlighted in the widget.
	// Default: - false.
	//
	// Experimental.
	FillAlarmRange *bool `field:"optional" json:"fillAlarmRange" yaml:"fillAlarmRange"`
	// An optional link included in the generated ticket description body.
	// Default: - no additional link will be added.
	//
	// Experimental.
	RunbookLink *string `field:"optional" json:"runbookLink" yaml:"runbookLink"`
}

