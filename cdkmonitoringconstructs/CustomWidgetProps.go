package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties of a custom widget.
// Experimental.
type CustomWidgetProps struct {
	// Lambda providing the widget contents.
	//
	// The Lambda function should return HTML with widget code.
	// The simplest Lambda example:
	// ```typescript
	// exports.handler = function (event, context, callback) {
	//    return callback(null, "<h1>Hello! This is a custom widget.</h1><pre>" + JSON.stringify(event, null, 2) + "</pre>");
	// };
	// ```.
	// Experimental.
	Handler awslambda.IFunction `field:"required" json:"handler" yaml:"handler"`
	// Arguments to pass to the Lambda.
	//
	// These arguments will be available in the Lambda context.
	// Experimental.
	HandlerParams interface{} `field:"optional" json:"handlerParams" yaml:"handlerParams"`
	// Height of the widget.
	// Default: - 6.
	//
	// Experimental.
	Height *float64 `field:"optional" json:"height" yaml:"height"`
	// Title for the graph.
	// Experimental.
	Title *string `field:"optional" json:"title" yaml:"title"`
	// Whether the widget should be updated (by calling the Lambda again) on refresh.
	// Default: - true.
	//
	// Experimental.
	UpdateOnRefresh *bool `field:"optional" json:"updateOnRefresh" yaml:"updateOnRefresh"`
	// Whether the widget should be updated (by calling the Lambda again) on resize.
	// Default: - true.
	//
	// Experimental.
	UpdateOnResize *bool `field:"optional" json:"updateOnResize" yaml:"updateOnResize"`
	// Whether the widget should be updated (by calling the Lambda again) on time range change.
	// Default: - true.
	//
	// Experimental.
	UpdateOnTimeRangeChange *bool `field:"optional" json:"updateOnTimeRangeChange" yaml:"updateOnTimeRangeChange"`
	// Width of the widget, in a grid of 24 units wide.
	// Default: - 6.
	//
	// Experimental.
	Width *float64 `field:"optional" json:"width" yaml:"width"`
}

