//go:build !no_runtime_type_checking

package cdkmonitoringconstructs

import (
	"fmt"
)

func (i *jsiiProxy_IEC2MetricFactoryStrategy) validateCreateMetricsParameters(metricFactory MetricFactory, metricName *string, statistic MetricStatistic) error {
	if metricFactory == nil {
		return fmt.Errorf("parameter metricFactory is required, but nil was provided")
	}

	if metricName == nil {
		return fmt.Errorf("parameter metricName is required, but nil was provided")
	}

	if statistic == "" {
		return fmt.Errorf("parameter statistic is required, but nil was provided")
	}

	return nil
}

