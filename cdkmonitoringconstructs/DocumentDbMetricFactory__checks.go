//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (d *jsiiProxy_DocumentDbMetricFactory) validateMetricReadLatencyInMillisParameters(latencyType LatencyType) error {
	if latencyType == "" {
		return fmt.Errorf("parameter latencyType is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DocumentDbMetricFactory) validateMetricWriteLatencyInMillisParameters(latencyType LatencyType) error {
	if latencyType == "" {
		return fmt.Errorf("parameter latencyType is required, but nil was provided")
	}

	return nil
}

func validateNewDocumentDbMetricFactoryParameters(metricFactory MetricFactory, props *DocumentDbMetricFactoryProps) error {
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
