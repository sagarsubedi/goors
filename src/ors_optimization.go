package main

type OrsOptimization struct {
	OrsBase
}

func NewOrsOptimization(args map[string]interface{}) *OrsOptimization {
	oo := &OrsOptimization{
		OrsBase: *NewOrsBase(args),
	}

	if _, exists := oo.DefaultArgs[constants.PropNames.Service]; !exists {
		oo.DefaultArgs[constants.PropNames.Service] = "optimization"
	}
	if _, exists := args[constants.PropNames.ApiVersion]; !exists {
		oo.DefaultArgs[constants.PropNames.ApiVersion] = constants.DefaultAPIVersion
	}

	return oo
}

func (oo *OrsOptimization) generatePayload(args map[string]interface{}) map[string]interface{} {
	payload := make(map[string]interface{})

	for key, val := range args {
		if !contains(constants.BaseUrlConstituents, key) {
			payload[key] = val
		}
	}

	return payload
}

func (oo *OrsOptimization) optimizationPromise() (interface{}, error) {
	oo.ArgsCache = orsUtil.SaveArgsToCache(oo.RequestArgs)

	payload := oo.generatePayload(oo.RequestArgs)

	return oo.createRequest(payload)
}

func (oo *OrsOptimization) Optimize(reqArgs map[string]interface{}) (interface{}, error) {
	oo.RequestArgs = reqArgs

	oo.checkHeaders()

	if _, exists := oo.DefaultArgs[constants.PropNames.Service]; !exists {
		if _, exists := reqArgs[constants.PropNames.Service]; !exists {
			reqArgs[constants.PropNames.Service] = "optimization"
		}
	}

	oo.RequestArgs = orsUtil.FillArgs(oo.DefaultArgs, oo.RequestArgs)

	return oo.optimizationPromise()
}
