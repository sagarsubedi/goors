package main

type OrsIsochrones struct {
	OrsBase
}

func NewOrsIsochrones(args map[string]interface{}) *OrsIsochrones {
	oi := &OrsIsochrones{
		OrsBase: *NewOrsBase(args),
	}

	if _, exists := oi.DefaultArgs[constants.PropNames.Service]; !exists {
		oi.DefaultArgs[constants.PropNames.Service] = "isochrones"
	}
	if _, exists := args[constants.PropNames.ApiVersion]; !exists {
		oi.DefaultArgs[constants.PropNames.ApiVersion] = constants.DefaultAPIVersion
	}

	return oi
}

func (oi *OrsIsochrones) getBody(args map[string]interface{}) map[string]interface{} {
	options := make(map[string]interface{})

	if restrictions, exists := args["restrictions"]; exists {
		options["profile_params"] = map[string]interface{}{
			"restrictions": restrictions,
		}
		delete(args, "restrictions")
	}
	if avoidables, exists := args["avoidables"]; exists {
		options["avoid_features"] = avoidables
		delete(args, "avoidables")
	}
	if avoidPolygons, exists := args["avoid_polygons"]; exists {
		options["avoid_polygons"] = avoidPolygons
		delete(args, "avoid_polygons")
	}

	if len(options) > 0 {
		args["options"] = options
	}

	return args
}
