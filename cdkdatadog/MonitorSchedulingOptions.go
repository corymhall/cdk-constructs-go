package cdkdatadog


// Configuration options for scheduling.
type MonitorSchedulingOptions struct {
	EvaluationWindow *MonitorSchedulingOptionsEvaluationWindow `field:"optional" json:"evaluationWindow" yaml:"evaluationWindow"`
}

