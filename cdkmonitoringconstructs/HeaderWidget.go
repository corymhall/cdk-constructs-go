package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/internal"
)

// Experimental.
type HeaderWidget interface {
	awscloudwatch.TextWidget
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

// The jsii proxy struct for HeaderWidget
type jsiiProxy_HeaderWidget struct {
	internal.Type__awscloudwatchTextWidget
}

func (j *jsiiProxy_HeaderWidget) Height() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"height",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HeaderWidget) Warnings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"warnings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HeaderWidget) WarningsV2() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"warningsV2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HeaderWidget) Width() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"width",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HeaderWidget) X() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"x",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HeaderWidget) Y() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"y",
		&returns,
	)
	return returns
}


// Experimental.
func NewHeaderWidget(text *string, level HeaderLevel, description *string, descriptionHeight *float64) HeaderWidget {
	_init_.Initialize()

	if err := validateNewHeaderWidgetParameters(text); err != nil {
		panic(err)
	}
	j := jsiiProxy_HeaderWidget{}

	_jsii_.Create(
		"cdk-monitoring-constructs.HeaderWidget",
		[]interface{}{text, level, description, descriptionHeight},
		&j,
	)

	return &j
}

// Experimental.
func NewHeaderWidget_Override(h HeaderWidget, text *string, level HeaderLevel, description *string, descriptionHeight *float64) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.HeaderWidget",
		[]interface{}{text, level, description, descriptionHeight},
		h,
	)
}

func (j *jsiiProxy_HeaderWidget)SetX(val *float64) {
	_jsii_.Set(
		j,
		"x",
		val,
	)
}

func (j *jsiiProxy_HeaderWidget)SetY(val *float64) {
	_jsii_.Set(
		j,
		"y",
		val,
	)
}

func (h *jsiiProxy_HeaderWidget) CopyMetricWarnings(ms ...awscloudwatch.IMetric) {
	args := []interface{}{}
	for _, a := range ms {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		h,
		"copyMetricWarnings",
		args,
	)
}

func (h *jsiiProxy_HeaderWidget) Position(x *float64, y *float64) {
	if err := h.validatePositionParameters(x, y); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"position",
		[]interface{}{x, y},
	)
}

func (h *jsiiProxy_HeaderWidget) ToJson() *[]interface{} {
	var returns *[]interface{}

	_jsii_.Invoke(
		h,
		"toJson",
		nil, // no parameters
		&returns,
	)

	return returns
}

