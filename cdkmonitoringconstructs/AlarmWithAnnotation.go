package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Representation of an alarm with additional information.
// Experimental.
type AlarmWithAnnotation struct {
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
	// Experimental.
	AlarmDescription *string `field:"required" json:"alarmDescription" yaml:"alarmDescription"`
	// Experimental.
	AlarmLabel *string `field:"required" json:"alarmLabel" yaml:"alarmLabel"`
	// Experimental.
	AlarmName *string `field:"required" json:"alarmName" yaml:"alarmName"`
	// Experimental.
	AlarmNameSuffix *string `field:"required" json:"alarmNameSuffix" yaml:"alarmNameSuffix"`
	// Experimental.
	AlarmRuleWhenAlarming awscloudwatch.IAlarmRule `field:"required" json:"alarmRuleWhenAlarming" yaml:"alarmRuleWhenAlarming"`
	// Experimental.
	AlarmRuleWhenInsufficientData awscloudwatch.IAlarmRule `field:"required" json:"alarmRuleWhenInsufficientData" yaml:"alarmRuleWhenInsufficientData"`
	// Experimental.
	AlarmRuleWhenOk awscloudwatch.IAlarmRule `field:"required" json:"alarmRuleWhenOk" yaml:"alarmRuleWhenOk"`
	// Experimental.
	Annotation *awscloudwatch.HorizontalAnnotation `field:"required" json:"annotation" yaml:"annotation"`
}

