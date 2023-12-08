package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Line graph widget with both left and right axes.
//
// The purpose of this custom class is to make the properties more strict.
// It will avoid graphs with undefined axes and dimensions.
// Experimental.
type DoubleAxisGraphWidget interface {
	awscloudwatch.GraphWidget
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
	// Add another metric to the left Y axis of the GraphWidget.
	// Experimental.
	AddLeftMetric(metric awscloudwatch.IMetric)
	// Add another metric to the right Y axis of the GraphWidget.
	// Experimental.
	AddRightMetric(metric awscloudwatch.IMetric)
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

// The jsii proxy struct for DoubleAxisGraphWidget
type jsiiProxy_DoubleAxisGraphWidget struct {
	internal.Type__awscloudwatchGraphWidget
}

func (j *jsiiProxy_DoubleAxisGraphWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DoubleAxisGraphWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DoubleAxisGraphWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DoubleAxisGraphWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DoubleAxisGraphWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DoubleAxisGraphWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewDoubleAxisGraphWidget(props *DoubleAxisGraphWidgetProps) DoubleAxisGraphWidget {
	_init_.Initialize()

	if err := validateNewDoubleAxisGraphWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DoubleAxisGraphWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DoubleAxisGraphWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewDoubleAxisGraphWidget_Override(d DoubleAxisGraphWidget, props *DoubleAxisGraphWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DoubleAxisGraphWidget",
		[]interface{}{props},
		d,
	)
}

func (j *jsiiProxy_DoubleAxisGraphWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_DoubleAxisGraphWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (d *jsiiProxy_DoubleAxisGraphWidget) AddLeftMetric(metric awscloudwatch.IMetric) {
	if err := d.validateAddLeftMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addLeftMetric",
		[]interface{}{metric},
	)
}

func (d *jsiiProxy_DoubleAxisGraphWidget) AddRightMetric(metric awscloudwatch.IMetric) {
	if err := d.validateAddRightMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addRightMetric",
		[]interface{}{metric},
	)
}

func (d *jsiiProxy_DoubleAxisGraphWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		d,
		"copyMetricWarnings",
		args,
	)
}

func (d *jsiiProxy_DoubleAxisGraphWidget) Position(x *float64, y *float64) {
	if err := d.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"position",
		[]interface{}{x, y},
	)
}

func (d *jsiiProxy_DoubleAxisGraphWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		d,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

