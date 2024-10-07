//go:build !no_runtime_type_checking

package cdkconstructlibrary

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validatePlatformQueue_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewPlatformQueueParameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

