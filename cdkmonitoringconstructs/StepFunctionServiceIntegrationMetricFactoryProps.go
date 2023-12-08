package cdkmonitoringconstructs


// Experimental.
type StepFunctionServiceIntegrationMetricFactoryProps struct {
	// Experimental.
	ServiceIntegrationResourceArn *string `field:"required" json:"serviceIntegrationResourceArn" yaml:"serviceIntegrationResourceArn"`
	// Default: - average.
	//
	// Experimental.
	RateComputationMethod RateComputationMethod `field:"optional" json:"rateComputationMethod" yaml:"rateComputationMethod"`
}

