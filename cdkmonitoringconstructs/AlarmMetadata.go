package cdkmonitoringconstructs


// Metadata of an alarm.
// Experimental.
type AlarmMetadata struct {
	// Experimental.
	Action IAlarmActionStrategy `field:"required" json:"action" yaml:"action"`
	// Experimental.
	CustomParams *map[string]interface{} `field:"optional" json:"customParams" yaml:"customParams"`
	// Experimental.
	CustomTags *[]*string `field:"optional" json:"customTags" yaml:"customTags"`
	// Experimental.
	DedupeString *string `field:"optional" json:"dedupeString" yaml:"dedupeString"`
	// Experimental.
	Disambiguator *string `field:"optional" json:"disambiguator" yaml:"disambiguator"`
}

