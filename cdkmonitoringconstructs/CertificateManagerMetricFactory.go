package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"
)

// Experimental.
type CertificateManagerMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	MetricDaysToExpiry() interface{}
}

// The jsii proxy struct for CertificateManagerMetricFactory
type jsiiProxy_CertificateManagerMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_CertificateManagerMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateManagerMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}


// Experimental.
func NewCertificateManagerMetricFactory(metricFactory MetricFactory, props *CertificateManagerMetricFactoryProps) CertificateManagerMetricFactory {
	_init_.Initialize()

	if err := validateNewCertificateManagerMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_CertificateManagerMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.CertificateManagerMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewCertificateManagerMetricFactory_Override(c CertificateManagerMetricFactory, metricFactory MetricFactory, props *CertificateManagerMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.CertificateManagerMetricFactory",
		[]interface{}{metricFactory, props},
		c,
	)
}

func (c *jsiiProxy_CertificateManagerMetricFactory) MetricDaysToExpiry() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		c,
		"metricDaysToExpiry",
		nil, // no parameters
		&returns,
	)

	return returns
}

