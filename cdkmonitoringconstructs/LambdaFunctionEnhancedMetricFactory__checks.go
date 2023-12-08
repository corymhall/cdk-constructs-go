//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

func validateNewLambdaFunctionEnhancedMetricFactoryParameters(metricFactory MetricFactory, lambdaFunction awslambda.IFunction) error {
	if metricFactory == nil {
		return fmt.Errorf("parameter metricFactory is required, but nil was provided")
	}

	if lambdaFunction == nil {
		return fmt.Errorf("parameter lambdaFunction is required, but nil was provided")
	}

	return nil
}

