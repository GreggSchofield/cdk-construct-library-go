//go:build no_runtime_type_checking

package cdkconstructlibrary

// Building without runtime type checking enabled, so all the below just return nil

func validatePlatformQueue_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPlatformQueueParameters(scope constructs.Construct, id *string, props *PlatformQueueProps) error {
	return nil
}

