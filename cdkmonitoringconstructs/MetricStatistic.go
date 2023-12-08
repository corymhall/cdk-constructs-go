package cdkmonitoringconstructs


// Metric aggregation statistic to be used with the IMetric objects.
// See: https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Statistics-definitions.html
//
// Experimental.
type MetricStatistic string

const (
	// 50th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P50 MetricStatistic = "P50"
	// 70th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P70 MetricStatistic = "P70"
	// 90th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P90 MetricStatistic = "P90"
	// 95th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P95 MetricStatistic = "P95"
	// 99th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P99 MetricStatistic = "P99"
	// 99.9th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P999 MetricStatistic = "P999"
	// 99.99th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P9999 MetricStatistic = "P9999"
	// 100th percentile of all datapoints.
	// Experimental.
	MetricStatistic_P100 MetricStatistic = "P100"
	// trimmed mean;
	//
	// calculates the average after removing the 50% of data points with the highest values.
	// Experimental.
	MetricStatistic_TM50 MetricStatistic = "TM50"
	// trimmed mean;
	//
	// calculates the average after removing the 30% of data points with the highest values.
	// Experimental.
	MetricStatistic_TM70 MetricStatistic = "TM70"
	// trimmed mean;
	//
	// calculates the average after removing the 10% of data points with the highest values.
	// Experimental.
	MetricStatistic_TM90 MetricStatistic = "TM90"
	// trimmed mean;
	//
	// calculates the average after removing the 5% of data points with the highest values.
	// Experimental.
	MetricStatistic_TM95 MetricStatistic = "TM95"
	// trimmed mean;
	//
	// calculates the average after removing the 1% of data points with the highest values.
	// Experimental.
	MetricStatistic_TM99 MetricStatistic = "TM99"
	// trimmed mean;
	//
	// calculates the average after removing the 0.1% of data points with the highest values
	// Experimental.
	MetricStatistic_TM999 MetricStatistic = "TM999"
	// trimmed mean;
	//
	// calculates the average after removing the 0.01% of data points with the highest values
	// Experimental.
	MetricStatistic_TM9999 MetricStatistic = "TM9999"
	// trimmed mean;
	//
	// calculates the average after removing the 1% lowest data points and the 1% highest data points.
	// Experimental.
	MetricStatistic_TM99_BOTH MetricStatistic = "TM99_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 5% lowest data points and the 5% highest data points.
	// Experimental.
	MetricStatistic_TM95_BOTH MetricStatistic = "TM95_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 10% lowest data points and the 10% highest data points.
	// Experimental.
	MetricStatistic_TM90_BOTH MetricStatistic = "TM90_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 15% lowest data points and the 15% highest data points.
	// Experimental.
	MetricStatistic_TM85_BOTH MetricStatistic = "TM85_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 20% lowest data points and the 20% highest data points.
	// Experimental.
	MetricStatistic_TM80_BOTH MetricStatistic = "TM80_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 25% lowest data points and the 25% highest data points.
	// Experimental.
	MetricStatistic_TM75_BOTH MetricStatistic = "TM75_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 30% lowest data points and the 30% highest data points.
	// Experimental.
	MetricStatistic_TM70_BOTH MetricStatistic = "TM70_BOTH"
	// trimmed mean;
	//
	// calculates the average after removing the 95% lowest data points.
	// Experimental.
	MetricStatistic_TM95_TOP MetricStatistic = "TM95_TOP"
	// trimmed mean;
	//
	// calculates the average after removing the 99% lowest data points.
	// Experimental.
	MetricStatistic_TM99_TOP MetricStatistic = "TM99_TOP"
	// trimmed mean;
	//
	// calculates the average after removing the 99.9% lowest data points
	// Experimental.
	MetricStatistic_TM999_TOP MetricStatistic = "TM999_TOP"
	// trimmed mean;
	//
	// calculates the average after removing the 99.99% lowest data points
	// Experimental.
	MetricStatistic_TM9999_TOP MetricStatistic = "TM9999_TOP"
	// winsorized mean;
	//
	// calculates the average while treating the 50% of the highest values to be equal to the value at the 50th percentile.
	// Experimental.
	MetricStatistic_WM50 MetricStatistic = "WM50"
	// winsorized mean;
	//
	// calculates the average while treating the 30% of the highest values to be equal to the value at the 70th percentile.
	// Experimental.
	MetricStatistic_WM70 MetricStatistic = "WM70"
	// winsorized mean;
	//
	// calculates the average while treating the 10% of the highest values to be equal to the value at the 90th percentile.
	// Experimental.
	MetricStatistic_WM90 MetricStatistic = "WM90"
	// winsorized mean;
	//
	// calculates the average while treating the 5% of the highest values to be equal to the value at the 95th percentile.
	// Experimental.
	MetricStatistic_WM95 MetricStatistic = "WM95"
	// winsorized mean;
	//
	// calculates the average while treating the 1% of the highest values to be equal to the value at the 99th percentile.
	// Experimental.
	MetricStatistic_WM99 MetricStatistic = "WM99"
	// winsorized mean;
	//
	// calculates the average while treating the 0.1% of the highest values to be equal to the value at the 99.9th percentile
	// Experimental.
	MetricStatistic_WM999 MetricStatistic = "WM999"
	// winsorized mean;
	//
	// calculates the average while treating the 0.01% of the highest values to be equal to the value at the 99.99th percentile
	// Experimental.
	MetricStatistic_WM9999 MetricStatistic = "WM9999"
	// winsorized mean;
	//
	// calculates the average while treating the highest 1% of data points to be the value of the 99% boundary, and treating the lowest 1% of data points to be the value of the 1% boundary.
	// Experimental.
	MetricStatistic_WM99_BOTH MetricStatistic = "WM99_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 5% of data points to be the value of the 95% boundary, and treating the lowest 5% of data points to be the value of the 5% boundary.
	// Experimental.
	MetricStatistic_WM95_BOTH MetricStatistic = "WM95_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 10% of data points to be the value of the 90% boundary, and treating the lowest 10% of data points to be the value of the 10% boundary.
	// Experimental.
	MetricStatistic_WM90_BOTH MetricStatistic = "WM90_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 15% of data points to be the value of the 85% boundary, and treating the lowest 15% of data points to be the value of the 15% boundary.
	// Experimental.
	MetricStatistic_WM85_BOTH MetricStatistic = "WM85_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 20% of data points to be the value of the 80% boundary, and treating the lowest 20% of data points to be the value of the 20% boundary.
	// Experimental.
	MetricStatistic_WM80_BOTH MetricStatistic = "WM80_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 25% of data points to be the value of the 75% boundary, and treating the lowest 25% of data points to be the value of the 25% boundary.
	// Experimental.
	MetricStatistic_WM75_BOTH MetricStatistic = "WM75_BOTH"
	// winsorized mean;
	//
	// calculates the average while treating the highest 30% of data points to be the value of the 70% boundary, and treating the lowest 30% of data points to be the value of the 30% boundary.
	// Experimental.
	MetricStatistic_WM70_BOTH MetricStatistic = "WM70_BOTH"
	// minimum of all datapoints.
	// Experimental.
	MetricStatistic_MIN MetricStatistic = "MIN"
	// maximum of all datapoints.
	// Experimental.
	MetricStatistic_MAX MetricStatistic = "MAX"
	// sum of all datapoints.
	// Experimental.
	MetricStatistic_SUM MetricStatistic = "SUM"
	// average of all datapoints.
	// Experimental.
	MetricStatistic_AVERAGE MetricStatistic = "AVERAGE"
	// number of datapoints.
	// Experimental.
	MetricStatistic_N MetricStatistic = "N"
)

