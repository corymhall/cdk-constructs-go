package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// A dashboard widget that can be customized using a Lambda.
// Experimental.
type CustomWidget interface {
	awscloudwatch.ConcreteWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	Warnings() *[]*string
	// Any warnings that are produced as a result of putting together this widget.
	// Experimental.
	WarningsV2() *map[string]*string
	// The amount of horizontal grid units the widget will take up.
	// Experimental.
	Width() *float64
	// Experimental.
	X() *float64
	// Experimental.
	SetX(val *float64)
	// Experimental.
	Y() *float64
	// Experimental.
	SetY(val *float64)
	// Copy the warnings from the given metric.
	// Experimental.
	CopyMetricWarnings(ms ...awscloudwatch.IMetric)
	// Place the widget at a given position.
	// Experimental.
	Position(x *float64, y *float64)
	// Return the widget JSON for use in the dashboard.
	// Experimental.
	ToJson() *[]interface{}
}

// The jsii proxy struct for CustomWidget
type jsiiProxy_CustomWidget struct {
	internal.Type__awscloudwatchConcreteWidget
}

func (j *jsiiProxy_CustomWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CustomWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewCustomWidget(props *CustomWidgetProps) CustomWidget {
	_init_.Initialize()

	if err := validateNewCustomWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CustomWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewCustomWidget_Override(c CustomWidget, props *CustomWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CustomWidget",
		[]interface{}{props},
		c,
	)
}

func (j *jsiiProxy_CustomWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_CustomWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (c *jsiiProxy_CustomWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		c,
		"copyMetricWarnings",
		args,
	)
}

func (c *jsiiProxy_CustomWidget) Position(x *float64, y *float64) {
	if err := c.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"position",
		[]interface{}{x, y},
	)
}

func (c *jsiiProxy_CustomWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		c,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

