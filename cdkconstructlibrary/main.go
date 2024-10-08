// cdk-construct-library
package cdkconstructlibrary

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"cdk-construct-library.PlatformQueue",
		reflect.TypeOf((*PlatformQueue)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_PlatformQueue{}
			_jsii_.InitJsiiProxy(&j.Type__constructsConstruct)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"cdk-construct-library.PlatformQueueProps",
		reflect.TypeOf((*PlatformQueueProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"cdk-construct-library.SecurityStandard",
		reflect.TypeOf((*SecurityStandard)(nil)).Elem(),
		map[string]interface{}{
			"GDPR": SecurityStandard_GDPR,
			"PCI_DSS": SecurityStandard_PCI_DSS,
			"SOC_2": SecurityStandard_SOC_2,
		},
	)
}
