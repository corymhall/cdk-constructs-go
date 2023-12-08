package cdkmonitoringconstructs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Custom metric search.
// Experimental.
type CustomMetricSearch struct {
	// search dimensions (can be empty).
	// Experimental.
	DimensionsMap *map[string]*string `field:"required" json:"dimensionsMap" yaml:"dimensionsMap"`
	// search query (can be empty).
	// Experimental.
	SearchQuery *string `field:"required" json:"searchQuery" yaml:"searchQuery"`
	// metric statistic.
	// Experimental.
	Statistic MetricStatistic `field:"required" json:"statistic" yaml:"statistic"`
	// custom label for the metrics.
	// Default: - " ".
	//
	// Experimental.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// metric namespace.
	// Default: - none.
	//
	// Experimental.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// metric period.
	// Default: - global default.
	//
	// Experimental.
	Period awscdk.Duration `field:"optional" json:"period" yaml:"period"`
	// axis (right or left) on which to graph metric default: AxisPosition.LEFT.
	// Experimental.
	Position AxisPosition `field:"optional" json:"position" yaml:"position"`
}

