//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func validateNewSecretsManagerMetricFactoryParameters(metricFactory MetricFactory) error {
	if metricFactory == nil {
		return fmt.Errorf("parameter metricFactory is required, but nil was provided")
	}

	return nil
}

