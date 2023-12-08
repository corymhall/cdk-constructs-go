//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
)

func (d *jsiiProxy_DynamoTableMetricFactory) validateMetricAverageSuccessfulRequestLatencyInMillisParameters(operation awsdynamodb.Operation) error {
	if operation == "" {
		return fmt.Errorf("parameter operation is required, but nil was provided")
	}

	return nil
}

func validateNewDynamoTableMetricFactoryParameters(metricFactory MetricFactory, props *DynamoTableMetricFactoryProps) error {
	if metricFactory == nil {
		return fmt.Errorf("parameter metricFactory is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

