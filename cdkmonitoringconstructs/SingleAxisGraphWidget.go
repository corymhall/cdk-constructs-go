package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Line graph widget with one axis only (left).
//
// If there is just one metric, it will hide the legend to save space.
// The purpose of this custom class is to make the properties more strict.
// It will avoid graphs with undefined axis and dimensions.
// Experimental.
type SingleAxisGraphWidget interface {
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

// The jsii proxy struct for SingleAxisGraphWidget
type jsiiProxy_SingleAxisGraphWidget struct {
	internal.Type__awscloudwatchGraphWidget
}

func (j *jsiiProxy_SingleAxisGraphWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleAxisGraphWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleAxisGraphWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleAxisGraphWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleAxisGraphWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SingleAxisGraphWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewSingleAxisGraphWidget(props *SingleAxisGraphWidgetProps) SingleAxisGraphWidget {
	_init_.Initialize()

	if err := validateNewSingleAxisGraphWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SingleAxisGraphWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SingleAxisGraphWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewSingleAxisGraphWidget_Override(s SingleAxisGraphWidget, props *SingleAxisGraphWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SingleAxisGraphWidget",
		[]interface{}{props},
		s,
	)
}

func (j *jsiiProxy_SingleAxisGraphWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_SingleAxisGraphWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (s *jsiiProxy_SingleAxisGraphWidget) AddLeftMetric(metric awscloudwatch.IMetric) {
	if err := s.validateAddLeftMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addLeftMetric",
		[]interface{}{metric},
	)
}

func (s *jsiiProxy_SingleAxisGraphWidget) AddRightMetric(metric awscloudwatch.IMetric) {
	if err := s.validateAddRightMetricParameters(metric); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addRightMetric",
		[]interface{}{metric},
	)
}

func (s *jsiiProxy_SingleAxisGraphWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		s,
		"copyMetricWarnings",
		args,
	)
}

func (s *jsiiProxy_SingleAxisGraphWidget) Position(x *float64, y *float64) {
	if err := s.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"position",
		[]interface{}{x, y},
	)
}

func (s *jsiiProxy_SingleAxisGraphWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		s,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

