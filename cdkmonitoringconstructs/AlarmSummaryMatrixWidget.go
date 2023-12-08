package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type AlarmSummaryMatrixWidget interface {
	awscloudwatch.ConcreteWidget
	// The amount of vertical grid units the widget will take up.
	// Experimental.
	Height() *float64
	// Experimental.
	Props() *AlarmSummaryMatrixWidgetProps
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

// The jsii proxy struct for AlarmSummaryMatrixWidget
type jsiiProxy_AlarmSummaryMatrixWidget struct {
	internal.Type__awscloudwatchConcreteWidget
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) Props() *AlarmSummaryMatrixWidgetProps {
	var returns *AlarmSummaryMatrixWidgetProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewAlarmSummaryMatrixWidget(props *AlarmSummaryMatrixWidgetProps) AlarmSummaryMatrixWidget {
	_init_.Initialize()

	if err := validateNewAlarmSummaryMatrixWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_AlarmSummaryMatrixWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.AlarmSummaryMatrixWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewAlarmSummaryMatrixWidget_Override(a AlarmSummaryMatrixWidget, props *AlarmSummaryMatrixWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.AlarmSummaryMatrixWidget",
		[]interface{}{props},
		a,
	)
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_AlarmSummaryMatrixWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (a *jsiiProxy_AlarmSummaryMatrixWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		a,
		"copyMetricWarnings",
		args,
	)
}

func (a *jsiiProxy_AlarmSummaryMatrixWidget) Position(x *float64, y *float64) {
	if err := a.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"position",
		[]interface{}{x, y},
	)
}

func (a *jsiiProxy_AlarmSummaryMatrixWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		a,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

