package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
)

// Experimental.
type MonitoringHeaderWidget interface {
	HeaderWidget
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

// The jsii proxy struct for MonitoringHeaderWidget
type jsiiProxy_MonitoringHeaderWidget struct {
	jsiiProxy_HeaderWidget
}

func (j *jsiiProxy_MonitoringHeaderWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringHeaderWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringHeaderWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringHeaderWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringHeaderWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitoringHeaderWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewMonitoringHeaderWidget(props *MonitoringHeaderWidgetProps) MonitoringHeaderWidget {
	_init_.Initialize()

	if err := validateNewMonitoringHeaderWidgetParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitoringHeaderWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringHeaderWidget",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewMonitoringHeaderWidget_Override(m MonitoringHeaderWidget, props *MonitoringHeaderWidgetProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.MonitoringHeaderWidget",
		[]interface{}{props},
		m,
	)
}

func (j *jsiiProxy_MonitoringHeaderWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_MonitoringHeaderWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (m *jsiiProxy_MonitoringHeaderWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		m,
		"copyMetricWarnings",
		args,
	)
}

func (m *jsiiProxy_MonitoringHeaderWidget) Position(x *float64, y *float64) {
	if err := m.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"position",
		[]interface{}{x, y},
	)
}

func (m *jsiiProxy_MonitoringHeaderWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		m,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

