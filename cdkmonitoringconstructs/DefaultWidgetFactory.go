package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type DefaultWidgetFactory interface {
	IWidgetFactory
	// Create widget representing an alarm detail.
	// Experimental.
	CreateAlarmDetailWidget(alarm *AlarmWithAnnotation) awscloudwatch.IWidget
}

// The jsii proxy struct for DefaultWidgetFactory
type jsiiProxy_DefaultWidgetFactory struct {
	jsiiProxy_IWidgetFactory
}

// Experimental.
func NewDefaultWidgetFactory() DefaultWidgetFactory {
	_init_.Initialize()

	j := jsiiProxy_DefaultWidgetFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultWidgetFactory",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDefaultWidgetFactory_Override(d DefaultWidgetFactory) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DefaultWidgetFactory",
		nil, // no parameters
		d,
	)
}

func (d *jsiiProxy_DefaultWidgetFactory) CreateAlarmDetailWidget(alarm *AlarmWithAnnotation) awscloudwatch.IWidget {
	if err := d.validateCreateAlarmDetailWidgetParameters(alarm); err != nil {
		panic(err)
	}
	var returns awscloudwatch.IWidget

	_jsii_.Invoke(
		d,
		"createAlarmDetailWidget",
		[]interface{}{alarm},
		&returns,
	)

	return returns
}

