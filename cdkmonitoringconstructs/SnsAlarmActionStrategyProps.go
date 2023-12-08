package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
)

// Experimental.
type SnsAlarmActionStrategyProps struct {
	// Target topic used when the alarm is triggered.
	// Experimental.
	OnAlarmTopic awssns.ITopic `field:"required" json:"onAlarmTopic" yaml:"onAlarmTopic"`
	// Optional target topic for when the alarm goes into the INSUFFICIENT_DATA state.
	// Default: - no notification sent.
	//
	// Experimental.
	OnInsufficientDataTopic awssns.ITopic `field:"optional" json:"onInsufficientDataTopic" yaml:"onInsufficientDataTopic"`
	// Optional target topic for when the alarm goes into the OK state.
	// Default: - no notification sent.
	//
	// Experimental.
	OnOkTopic awssns.ITopic `field:"optional" json:"onOkTopic" yaml:"onOkTopic"`
}

