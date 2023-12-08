package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscodebuild"
)

// Experimental.
type CodeBuildProjectMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Project() awscodebuild.IProject
	// Experimental.
	MetricBuildCount() interface{}
	// Experimental.
	MetricDurationP50InSeconds() interface{}
	// Experimental.
	MetricDurationP90InSeconds() interface{}
	// Experimental.
	MetricDurationP99InSeconds() interface{}
	// Experimental.
	MetricFailedBuildCount() interface{}
	// Experimental.
	MetricFailedBuildRate() interface{}
	// Experimental.
	MetricSucceededBuildCount() interface{}
}

// The jsii proxy struct for CodeBuildProjectMetricFactory
type jsiiProxy_CodeBuildProjectMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_CodeBuildProjectMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CodeBuildProjectMetricFactory) Project() awscodebuild.IProject {
	var returns awscodebuild.IProject
	_jsii_.Get(
		j,
		"project",
		&returns,
	)
	return returns
}


// Experimental.
func NewCodeBuildProjectMetricFactory(metricFactory MetricFactory, props *CodeBuildProjectMetricFactoryProps) CodeBuildProjectMetricFactory {
	_init_.Initialize()

	if err := validateNewCodeBuildProjectMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CodeBuildProjectMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CodeBuildProjectMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCodeBuildProjectMetricFactory_Override(c CodeBuildProjectMetricFactory, metricFactory MetricFactory, props *CodeBuildProjectMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CodeBuildProjectMetricFactory",
		[]interface{}{metricFactory, props},
		c,
	)
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricBuildCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricBuildCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricDurationP50InSeconds() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricDurationP50InSeconds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricDurationP90InSeconds() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricDurationP90InSeconds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricDurationP99InSeconds() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricDurationP99InSeconds",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricFailedBuildCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricFailedBuildCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricFailedBuildRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricFailedBuildRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CodeBuildProjectMetricFactory) MetricSucceededBuildCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricSucceededBuildCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

