package main

import jsoniter "github.com/json-iterator/go"

type OrsDirections struct {
	OrsBase
}

func NewOrsDirections(args map[string]interface{}) *OrsDirections {

	od := &OrsDirections{
		OrsBase: *NewOrsBase(args),
	}

	if _, exists := od.DefaultArgs[constants.PropNames.Service]; !exists {
		od.DefaultArgs[constants.PropNames.Service] = "directions"
	}

	if _, exists := args[constants.PropNames.ApiVersion]; !exists {
		od.DefaultArgs[constants.PropNames.ApiVersion] = constants.DefaultAPIVersion
	}

	return od
}

func GetBody(args map[string]interface{}) map[string]interface{} {
	if options, exists := args["options"]; exists {
		if _, ok := options.(string); ok {
			var optionsMaps map[string]interface{}
			if err := jsoniter.Unmarshal([]byte(options.(string)), &optionsMaps); err == nil {
				args["options"] = optionsMaps
			}
		}
	}

	if restrictions, exists := args["restrictions"]; exists {
		if args["options"] == nil {
			args["options"] = make(map[string]interface{})
		}

		if optionsMap, ok := args["options"].(map[string]interface{}); ok {
			optionsMap["profile_params"] = map[string]interface{}{
				"restrictions": restrictions,
			}

			delete(args, "restrictions")
		}
	}

	if avoidables, exists := args["avoidables"]; exists {
		if args["options"] == nil {
			args["options"] = make(map[string]interface{})
		}
		if optionsMap, ok := args["options"].(map[string]interface{}); ok {
			optionsMap["avoid_features"] = avoidables
			delete(args, "avoidables")
		}
	}

	return args

}
