package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssynthetics"
)

// Experimental.
type SyntheticsCanaryMetricFactory interface {
	// Experimental.
	Canary() awssynthetics.Canary
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	RateComputationMethod() RateComputationMethod
	// Experimental.
	Metric4xxErrorCount() interface{}
	// Experimental.
	Metric4xxErrorRate() interface{}
	// Experimental.
	Metric5xxFaultCount() interface{}
	// Experimental.
	Metric5xxFaultRate() interface{}
	// Experimental.
	MetricLatencyAverageInMillis() interface{}
	// Experimental.
	MetricSuccessInPercent() interface{}
}

// The jsii proxy struct for SyntheticsCanaryMetricFactory
type jsiiProxy_SyntheticsCanaryMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_SyntheticsCanaryMetricFactory) Canary() awssynthetics.Canary {
	var returns awssynthetics.Canary
	_jsii_.Get(
		j,
		"canary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SyntheticsCanaryMetricFactory) RateComputationMethod() RateComputationMethod {
	var returns RateComputationMethod
	_jsii_.Get(
		j,
		"rateComputationMethod",
		&returns,
	)
	return returns
}


// Experimental.
func NewSyntheticsCanaryMetricFactory(metricFactory MetricFactory, props *SyntheticsCanaryMetricFactoryProps) SyntheticsCanaryMetricFactory {
	_init_.Initialize()

	if err := validateNewSyntheticsCanaryMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_SyntheticsCanaryMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.SyntheticsCanaryMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryMetricFactory_Override(s SyntheticsCanaryMetricFactory, metricFactory MetricFactory, props *SyntheticsCanaryMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.SyntheticsCanaryMetricFactory",
		[]interface{}{metricFactory, props},
		s,
	)
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) Metric4xxErrorCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metric4xxErrorCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) Metric4xxErrorRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metric4xxErrorRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) Metric5xxFaultCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metric5xxFaultCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) Metric5xxFaultRate() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metric5xxFaultRate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) MetricLatencyAverageInMillis() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricLatencyAverageInMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SyntheticsCanaryMetricFactory) MetricSuccessInPercent() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"metricSuccessInPercent",
		nil, // no parameters
		&returns,
	)

	return returns
}

