package cdkmonitoringconstructs


// Experimental.
type OpenSearchClusterMetricFactoryProps struct {
	// Experimental.
	Domain interface{} `field:"required" json:"domain" yaml:"domain"`
	// Default: - true.
	//
	// Experimental.
	FillTpsWithZeroes *bool `field:"optional" json:"fillTpsWithZeroes" yaml:"fillTpsWithZeroes"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

