package main

type OrsElevation struct {
	OrsBase
}

var orsUtil = OrsUtil{}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func NewOrsElevation(args map[string]interface{}) *OrsElevation {
	orsElevation := &OrsElevation{
		OrsBase: *NewOrsBase(args),
	}
	return orsElevation
}

func (oe *OrsElevation) GeneratePayload(args map[string]interface{}) map[string]interface{} {
	payload := make(map[string]interface{})
	for key, value := range args {
		if !contains(constants.BaseUrlConstituents, key) {
			payload[key] = value
		}
	}
	return payload
}

func (oe *OrsElevation) ElevationPromise() (interface{}, error) {
	oe.ArgsCache = orsUtil.SaveArgsToCache(oe.RequestArgs)
	payload := oe.GeneratePayload(oe.RequestArgs)
	return oe.createRequest(payload)
}

func (oe *OrsElevation) LineElevation(reqArgs map[string]interface{}) (interface{}, error) {
	oe.RequestArgs = reqArgs
	oe.checkHeaders()

	_, servicePresentInDefaultArgs := oe.DefaultArgs[constants.PropNames.Service]
	_, servicePresentInRequestArgs := oe.RequestArgs[constants.PropNames.Service]

	if !servicePresentInDefaultArgs && !servicePresentInRequestArgs {
		oe.RequestArgs[constants.PropNames.Service] = "elevation/line"
	}
	oe.RequestArgs = orsUtil.FillArgs(oe.DefaultArgs, oe.RequestArgs)
	return oe.ElevationPromise()
}

func (oe *OrsElevation) PointElevation(reqArgs map[string]interface{}) (interface{}, error) {
	oe.RequestArgs = reqArgs
	oe.checkHeaders()

	_, servicePresentInDefaultArgs := oe.DefaultArgs[constants.PropNames.Service]
	_, servicePresentInRequestArgs := oe.RequestArgs[constants.PropNames.Service]

	if !servicePresentInDefaultArgs && !servicePresentInRequestArgs {
		oe.RequestArgs[constants.PropNames.Service] = "elevation/point"
	}
	oe.RequestArgs = orsUtil.FillArgs(oe.DefaultArgs, oe.RequestArgs)
	return oe.ElevationPromise()
}
