package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Strategy for creating widgets.
// Experimental.
type IWidgetFactory interface {
	// Create widget representing an alarm detail.
	// Experimental.
	CreateAlarmDetailWidget(alarm *AlarmWithAnnotation) awscloudwatch.IWidget
}

// The jsii proxy for IWidgetFactory
type jsiiProxy_IWidgetFactory struct {
	_ byte // padding
}

func (i *jsiiProxy_IWidgetFactory) CreateAlarmDetailWidget(alarm *AlarmWithAnnotation) awscloudwatch.IWidget {
	if err := i.validateCreateAlarmDetailWidgetParameters(alarm); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IWidget

	_jsii_.Invoke(
		i,
		"createAlarmDetailWidget",
		[]interface{}{alarm},
		&returns,
	)

	return returns
}

