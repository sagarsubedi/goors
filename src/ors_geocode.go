package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type FocusPoint [2]float64
type BoundaryBBox [2][2]float64
type Point struct {
	LatLng [2]float64
}
type BoundaryCircle struct {
	LatLng [2]float64
	Radius float64
}
type Sources []string
type Layers []string
type BoundaryCountry string

type OrsGeocode struct {
	OrsBase
	lookupParameter map[string]func(key string, val interface{}) string
}

func NewOrsGeocode(args map[string]interface{}) *OrsGeocode {
	og := &OrsGeocode{
		OrsBase: *NewOrsBase(args),
	}
	og.lookupParameterFuncs()
	return og
}

func (og *OrsGeocode) lookupParameterFuncs() {
	og.lookupParameter = map[string]func(key string, val interface{}) string{
		"api_key": func(key string, val interface{}) string {
			return key + "=" + val.(string)
		},
		"text": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"focus_point": func(key string, val interface{}) string {
			valSlice := val.(FocusPoint)
			return "&focus.point.lon=" + fmt.Sprintf("%f", valSlice[1]) + "&focus.point.lat=" + fmt.Sprintf("%f", valSlice[0])
		},
		"boundary_bbox": func(key string, val interface{}) string {
			valSlice := val.(BoundaryBBox)
			return "&boundary.rect.min_lon=" + fmt.Sprintf("%f", valSlice[0][1]) +
				"&boundary.rect.min_lat=" + fmt.Sprintf("%f", valSlice[0][0]) +
				"&boundary.rect.max_lon=" + fmt.Sprintf("%f", valSlice[1][1]) +
				"&boundary.rect.max_lat=" + fmt.Sprintf("%f", valSlice[1][0])
		},
		"point": func(key string, val interface{}) string {
			point := val.(Point)
			return "&point.lon=" + fmt.Sprintf("%f", point.LatLng[1]) + "&point.lat=" + fmt.Sprintf("%f", point.LatLng[0])
		},
		"boundary_circle": func(key string, val interface{}) string {
			circle := val.(BoundaryCircle)
			return "&boundary.circle.lon=" + fmt.Sprintf("%f", circle.LatLng[1]) + "&boundary.circle.lat=" + fmt.Sprintf("%f", circle.LatLng[0]) + "&boundary.circle.radius=" + fmt.Sprintf("%f", circle.Radius)
		},
		"sources": func(key string, val interface{}) string {
			sources := val.(Sources)
			return "&sources=" + strings.Join(sources, ",")
		},
		"layers": func(key string, val interface{}) string {
			layers := val.(Layers)
			return "&layers=" + strings.Join(layers, ",")
		},
		"boundary_country": func(key string, val interface{}) string {
			return "&boundary.country=" + string(val.(BoundaryCountry))
		},
		"size": func(key string, val interface{}) string {
			return "&" + key + "=" + fmt.Sprintf("%v", val)
		},
		"address": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"neighbourhood": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"borough": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"locality": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"county": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"region": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"postalcode": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
		"country": func(key string, val interface{}) string {
			return "&" + key + "=" + url.QueryEscape(val.(string))
		},
	}
}

func (og *OrsGeocode) GetParametersAsQueryString(args map[string]interface{}) string {
	var queryString string
	for key, value := range args {
		if !contains(constants.BaseUrlConstituents, key) {
			queryString += og.lookupParameter[key](key, value)
		}
	}
	return queryString
}

func (og *OrsGeocode) FetchGetRequest() (*http.Response, error) {
	url := orsUtil.PrepareUrl(og.RequestArgs)
	url += "?" + og.GetParametersAsQueryString(og.RequestArgs)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range og.CustomHeaders {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}

func (og *OrsGeocode) GeocodePromise() (interface{}, error) {
	response, err := og.FetchGetRequest()
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func (og *OrsGeocode) Geocode(reqArgs map[string]interface{}) (interface{}, error) {
	og.RequestArgs = reqArgs
	og.checkHeaders()

	if _, exists := og.DefaultArgs[constants.PropNames.Service]; !exists {
		if _, exists := og.RequestArgs[constants.PropNames.Service]; !exists {
			og.RequestArgs[constants.PropNames.Service] = "geocode/search"
		}
	}

	og.RequestArgs = orsUtil.FillArgs(og.DefaultArgs, og.RequestArgs)

	return og.GeocodePromise()
}

func (og *OrsGeocode) reverseGeocode(reqArgs map[string]interface{}) (interface{}, error) {
	og.RequestArgs = reqArgs
	og.checkHeaders()

	if _, exists := og.DefaultArgs[constants.PropNames.Service]; !exists {
		if _, exists := og.RequestArgs[constants.PropNames.Service]; !exists {
			og.RequestArgs[constants.PropNames.Service] = "geocode/reverse"
		}
	}

	og.RequestArgs = orsUtil.FillArgs(og.DefaultArgs, og.RequestArgs)

	return og.GeocodePromise()
}

func (og *OrsGeocode) structuredGeocode(reqArgs map[string]interface{}) (interface{}, error) {
	og.RequestArgs = reqArgs
	og.checkHeaders()

	if _, exists := og.DefaultArgs[constants.PropNames.Service]; !exists {
		if _, exists := og.RequestArgs[constants.PropNames.Service]; !exists {
			og.RequestArgs[constants.PropNames.Service] = "geocode/search/structured"
		}
	}

	og.RequestArgs = orsUtil.FillArgs(og.DefaultArgs, og.RequestArgs)

	return og.GeocodePromise()
}
