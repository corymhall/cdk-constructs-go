package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type AlarmMatrixWidgetProps struct {
	// list of alarms to show.
	// Experimental.
	Alarms *[]awscloudwatch.IAlarm `field:"required" json:"alarms" yaml:"alarms"`
	// desired height.
	// Default: - auto calculated based on alarm number (3 to 8).
	//
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// widget title.
	// Default: - no title.
	//
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
}

