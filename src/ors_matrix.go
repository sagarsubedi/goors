package main

type OrsMatrix struct {
	OrsBase
}

func NewOrsMatrix(args map[string]interface{}) *OrsMatrix {
	om := &OrsMatrix{
		OrsBase: *NewOrsBase(args),
	}

	if _, exists := om.DefaultArgs[constants.PropNames.Service]; !exists {
		om.DefaultArgs[constants.PropNames.Service] = "matrix"
	}
	if _, exists := args[constants.PropNames.ApiVersion]; !exists {
		om.DefaultArgs[constants.PropNames.ApiVersion] = constants.DefaultAPIVersion
	}

	return om
}
