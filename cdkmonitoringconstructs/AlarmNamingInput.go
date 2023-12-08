package cdkmonitoringconstructs


// Experimental.
type AlarmNamingInput struct {
	// Experimental.
	AlarmNameSuffix *string `field:"required" json:"alarmNameSuffix" yaml:"alarmNameSuffix"`
	// Experimental.
	AlarmDedupeStringSuffix *string `field:"optional" json:"alarmDedupeStringSuffix" yaml:"alarmDedupeStringSuffix"`
	// Experimental.
	AlarmNameOverride *string `field:"optional" json:"alarmNameOverride" yaml:"alarmNameOverride"`
	// Experimental.
	DedupeStringOverride *string `field:"optional" json:"dedupeStringOverride" yaml:"dedupeStringOverride"`
	// Experimental.
	Disambiguator *string `field:"optional" json:"disambiguator" yaml:"disambiguator"`
}

