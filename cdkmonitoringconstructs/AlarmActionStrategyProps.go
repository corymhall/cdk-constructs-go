package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Properties necessary to append actions to an alarm.
// Experimental.
type AlarmActionStrategyProps struct {
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
	// Experimental.
	Alarm awscloudwatch.AlarmBase `field:"required" json:"alarm" yaml:"alarm"`
}

