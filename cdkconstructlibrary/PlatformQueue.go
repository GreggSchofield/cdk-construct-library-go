package cdkconstructlibrary

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/GreggSchofield/cdk-construct-library-go/cdkconstructlibrary/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/GreggSchofield/cdk-construct-library-go/cdkconstructlibrary/internal"
)

type PlatformQueue interface {
	constructs.Construct
	// The tree node.
	Node() constructs.Node
	// Returns a string representation of this construct.
	ToString() *string
}

// The jsii proxy struct for PlatformQueue
type jsiiProxy_PlatformQueue struct {
	internal.Type__constructsConstruct
}

func (j *jsiiProxy_PlatformQueue) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


func NewPlatformQueue(scope constructs.Construct, id *string) PlatformQueue {
	_init_.Initialize()

	if err := validateNewPlatformQueueParameters(scope, id); err != nil {
		panic(err)
	}
	j := jsiiProxy_PlatformQueue{}

	_jsii_.Create(
		"cdk-construct-library.PlatformQueue",
		[]interface{}{scope, id},
		&j,
	)

	return &j
}

func NewPlatformQueue_Override(p PlatformQueue, scope constructs.Construct, id *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk-construct-library.PlatformQueue",
		[]interface{}{scope, id},
		p,
	)
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func PlatformQueue_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePlatformQueue_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"cdk-construct-library.PlatformQueue",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PlatformQueue) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

