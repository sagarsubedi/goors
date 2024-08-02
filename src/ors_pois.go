package main

type OrsPois struct {
	OrsBase
}

func NewOrsPois(args map[string]interface{}) *OrsPois {
	op := &OrsPois{
		OrsBase: *NewOrsBase(args),
	}

	if _, exists := op.DefaultArgs[constants.PropNames.Service]; !exists {
		op.DefaultArgs[constants.PropNames.Service] = "pois"
	}

	return op
}

func (op *OrsPois) generatePayload(args map[string]interface{}) map[string]interface{} {
	payload := make(map[string]interface{})

	for key, val := range args {
		if !(contains(constants.BaseUrlConstituents, key) || key == constants.PropNames.ApiKey || key == constants.PropNames.Timeout) {
			payload[key] = val
		}
	}

	return payload
}

func (op *OrsPois) poisPromise() (interface{}, error) {
	// the request arg is required by the API as part of the body
	if op.RequestArgs["request"] == nil {
		op.RequestArgs["request"] = "pois"
	}

	op.ArgsCache = orsUtil.SaveArgsToCache(op.RequestArgs)

	if op.RequestArgs[constants.PropNames.Service] != nil {
		delete(op.RequestArgs, constants.PropNames.Service)
	}

	payload := op.generatePayload(op.RequestArgs)

	return op.createRequest(payload)
}

func (op *OrsPois) Pois(reqArgs map[string]interface{}) (interface{}, error) {
	op.RequestArgs = reqArgs

	op.checkHeaders()

	op.RequestArgs = orsUtil.FillArgs(op.DefaultArgs, op.RequestArgs)

	return op.poisPromise()
}
