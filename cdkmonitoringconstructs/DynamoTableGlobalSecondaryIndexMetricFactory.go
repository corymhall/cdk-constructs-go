package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type DynamoTableGlobalSecondaryIndexMetricFactory interface {
	// Experimental.
	DimensionsMap() *map[string]*string
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Table() awsdynamodb.ITable
	// Experimental.
	MetricConsumedReadCapacityUnits() interface{}
	// Experimental.
	MetricConsumedWriteCapacityUnits() interface{}
	// Experimental.
	MetricIndexConsumedWriteUnitsMetric() interface{}
	// Experimental.
	MetricProvisionedReadCapacityUnits() interface{}
	// Experimental.
	MetricProvisionedWriteCapacityUnits() interface{}
	// Experimental.
	MetricThrottledIndexRequestCount() interface{}
	// Experimental.
	MetricThrottledReadRequestCount() interface{}
	// Experimental.
	MetricThrottledWriteRequestCount() interface{}
}

// The jsii proxy struct for DynamoTableGlobalSecondaryIndexMetricFactory
type jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) DimensionsMap() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dimensionsMap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) Table() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoTableGlobalSecondaryIndexMetricFactory(metricFactory MetricFactory, props *DynamoTableGlobalSecondaryIndexMetricFactoryProps) DynamoTableGlobalSecondaryIndexMetricFactory {
	_init_.Initialize()

	if err := validateNewDynamoTableGlobalSecondaryIndexMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableGlobalSecondaryIndexMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoTableGlobalSecondaryIndexMetricFactory_Override(d DynamoTableGlobalSecondaryIndexMetricFactory, metricFactory MetricFactory, props *DynamoTableGlobalSecondaryIndexMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableGlobalSecondaryIndexMetricFactory",
		[]interface{}{metricFactory, props},
		d,
	)
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricConsumedReadCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricConsumedReadCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricConsumedWriteCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricConsumedWriteCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricIndexConsumedWriteUnitsMetric() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricIndexConsumedWriteUnitsMetric",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricProvisionedReadCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricProvisionedReadCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricProvisionedWriteCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricProvisionedWriteCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricThrottledIndexRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricThrottledIndexRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricThrottledReadRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricThrottledReadRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableGlobalSecondaryIndexMetricFactory) MetricThrottledWriteRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricThrottledWriteRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

