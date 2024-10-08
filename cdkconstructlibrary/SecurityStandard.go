package cdkconstructlibrary


// Enum for the different types of security standards.
type SecurityStandard string

const (
	// General Data Protection Regulation.
	SecurityStandard_GDPR SecurityStandard = "GDPR"
	// Payment Card Industry Data Security Standard.
	SecurityStandard_PCI_DSS SecurityStandard = "PCI_DSS"
	// System and Organization Controls 2.
	SecurityStandard_SOC_2 SecurityStandard = "SOC_2"
)

