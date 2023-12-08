package cdkmonitoringconstructs


// Experimental.
type ElastiCacheClusterMetricFactoryProps struct {
	// Cluster to monitor.
	// Default: - monitor all clusters.
	//
	// Experimental.
	ClusterId *string `field:"optional" json:"clusterId" yaml:"clusterId"`
}

