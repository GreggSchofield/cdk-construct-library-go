package cdkconstructlibrary


type SecurityStandard string

const (
	SecurityStandard_GDPR SecurityStandard = "GDPR"
	SecurityStandard_PCI_DSS SecurityStandard = "PCI_DSS"
	// System and Organization Controls 2.
	SecurityStandard_SOC_2 SecurityStandard = "SOC_2"
)

