package cdkmonitoringconstructs


// Experimental.
type RdsClusterMetricFactoryProps struct {
	// database cluster (either this or `clusterIdentifier` need to be specified).
	// Experimental.
	Cluster interface{} `field:"optional" json:"cluster" yaml:"cluster"`
	// database cluster identifier (either this or `cluster` need to be specified).
	// Deprecated: please use `cluster` instead.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
}

