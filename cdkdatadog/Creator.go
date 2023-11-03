package cdkdatadog


type Creator struct {
	// Email of the creator of the monitor.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Handle of the creator of the monitor.
	Handle *string `field:"optional" json:"handle" yaml:"handle"`
	// Name of the creator of the monitor.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

