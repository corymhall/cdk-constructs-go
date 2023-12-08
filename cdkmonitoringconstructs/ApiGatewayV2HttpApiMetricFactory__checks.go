//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) validateMetricIntegrationLatencyInMillisParameters(latencyType LatencyType) error {
	if latencyType == "" {
		return fmt.Errorf("parameter latencyType is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_ApiGatewayV2HttpApiMetricFactory) validateMetricLatencyInMillisParameters(latencyType LatencyType) error {
	if latencyType == "" {
		return fmt.Errorf("parameter latencyType is required, but nil was provided")
	}

	return nil
}

func validateNewApiGatewayV2HttpApiMetricFactoryParameters(metricFactory MetricFactory, props *ApiGatewayV2HttpApiMetricFactoryProps) error {
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
