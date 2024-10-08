package cdkconstructlibrary


// Properties for the PlatformQueue Construct.
type PlatformQueueProps struct {
	// Whether the queue has a dead letter queue (dlq).
	Dlq *bool `field:"optional" json:"dlq" yaml:"dlq"`
	// Whether the queue is first-in first-out (fifo).
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// The security standard for the queue.
	SecurityStandard SecurityStandard `field:"optional" json:"securityStandard" yaml:"securityStandard"`
}

