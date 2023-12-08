package cdkmonitoringconstructs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/corymhall/cdk-constructs-go/cdkmonitoringconstructs/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

// Experimental.
type DynamoTableMetricFactory interface {
	// Experimental.
	MetricFactory() MetricFactory
	// Experimental.
	Table() awsdynamodb.ITable
	// Experimental.
	MetricAverageSuccessfulRequestLatencyInMillis(operation awsdynamodb.Operation) interface{}
	// Experimental.
	MetricConsumedReadCapacityUnits() interface{}
	// Experimental.
	MetricConsumedWriteCapacityUnits() interface{}
	// Experimental.
	MetricProvisionedReadCapacityUnits() interface{}
	// Experimental.
	MetricProvisionedWriteCapacityUnits() interface{}
	// Experimental.
	MetricReadCapacityUtilizationPercentage() interface{}
	// Experimental.
	MetricSearchAverageSuccessfulRequestLatencyInMillis() awscloudwatch.IMetric
	// This represents the number of requests that resulted in a 500 (server error) error code.
	//
	// It summarizes across the basic CRUD operations:
	// GetItem, BatchGetItem, Scan, Query, GetRecords, PutItem, DeleteItem, UpdateItem, BatchWriteItem
	//
	// It’s usually equal to zero.
	// Experimental.
	MetricSystemErrorsCount() interface{}
	// Experimental.
	MetricThrottledReadRequestCount() interface{}
	// Experimental.
	MetricThrottledWriteRequestCount() interface{}
	// Experimental.
	MetricWriteCapacityUtilizationPercentage() interface{}
}

// The jsii proxy struct for DynamoTableMetricFactory
type jsiiProxy_DynamoTableMetricFactory struct {
	_ byte // padding
}

func (j *jsiiProxy_DynamoTableMetricFactory) MetricFactory() MetricFactory {
	var returns MetricFactory
	_jsii_.Get(
		j,
		"metricFactory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DynamoTableMetricFactory) Table() awsdynamodb.ITable {
	var returns awsdynamodb.ITable
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}


// Experimental.
func NewDynamoTableMetricFactory(metricFactory MetricFactory, props *DynamoTableMetricFactoryProps) DynamoTableMetricFactory {
	_init_.Initialize()

	if err := validateNewDynamoTableMetricFactoryParameters(metricFactory, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_DynamoTableMetricFactory{}

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableMetricFactory",
		[]interface{}{metricFactory, props},
		&j,
	)

	return &j
}

// Experimental.
func NewDynamoTableMetricFactory_Override(d DynamoTableMetricFactory, metricFactory MetricFactory, props *DynamoTableMetricFactoryProps) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-monitoring-constructs.DynamoTableMetricFactory",
		[]interface{}{metricFactory, props},
		d,
	)
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricAverageSuccessfulRequestLatencyInMillis(operation awsdynamodb.Operation) interface{} {
	if err := d.validateMetricAverageSuccessfulRequestLatencyInMillisParameters(operation); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricAverageSuccessfulRequestLatencyInMillis",
		[]interface{}{operation},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricConsumedReadCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricConsumedReadCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricConsumedWriteCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricConsumedWriteCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricProvisionedReadCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricProvisionedReadCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricProvisionedWriteCapacityUnits() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricProvisionedWriteCapacityUnits",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricReadCapacityUtilizationPercentage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricReadCapacityUtilizationPercentage",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricSearchAverageSuccessfulRequestLatencyInMillis() awscloudwatch.IMetric {
	var returns awscloudwatch.IMetric

	_jsii_.Invoke(
		d,
		"metricSearchAverageSuccessfulRequestLatencyInMillis",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricSystemErrorsCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricSystemErrorsCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricThrottledReadRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricThrottledReadRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricThrottledWriteRequestCount() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricThrottledWriteRequestCount",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DynamoTableMetricFactory) MetricWriteCapacityUtilizationPercentage() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"metricWriteCapacityUtilizationPercentage",
		nil, // no parameters
		&returns,
	)

	return returns
}

