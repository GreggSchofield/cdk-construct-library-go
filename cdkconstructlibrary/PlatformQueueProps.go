package cdkconstructlibrary


type PlatformQueueProps struct {
	Dlq *bool `field:"optional" json:"dlq" yaml:"dlq"`
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	SecurityStandard SecurityStandard `field:"optional" json:"securityStandard" yaml:"securityStandard"`
}

