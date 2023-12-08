package cdkmonitoringconstructs


// Experimental.
type GlueJobMetricFactoryProps struct {
	// Experimental.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

