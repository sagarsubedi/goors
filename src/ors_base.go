package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type OrsBase struct {
	DefaultArgs   map[string]interface{}
	RequestArgs   map[string]interface{}
	HttpArgs      map[string]interface{}
	ArgsCache     map[string]interface{}
	CustomHeaders map[string]string
}

func NewOrsBase(args map[string]interface{}) *OrsBase {
	obj := &OrsBase{
		DefaultArgs:   make(map[string]interface{}),
		RequestArgs:   make(map[string]interface{}),
		ArgsCache:     make(map[string]interface{}),
		CustomHeaders: make(map[string]string),
	}

	obj.setRequestDefaults(args)
	return obj
}

func (ob *OrsBase) setRequestDefaults(args map[string]interface{}) {
	ob.DefaultArgs[constants.PropNames.Host] = constants.DefaultHost
	if host, ok := args[constants.PropNames.Host]; ok {
		ob.DefaultArgs[constants.PropNames.Host] = host
	}

	if service, ok := args[constants.PropNames.Service]; ok {
		ob.DefaultArgs[constants.PropNames.Service] = service
	}

	if apiKey, ok := args[constants.PropNames.ApiKey]; ok {
		ob.DefaultArgs[constants.PropNames.ApiKey] = apiKey
	} else {
		panic(errors.New(constants.MissingAPIKeyMsg))
	}
}

func (ob *OrsBase) checkHeaders() {
	if customHeaders, ok := ob.RequestArgs["customHeaders"].(map[string]string); ok {
		ob.CustomHeaders = customHeaders
		delete(ob.RequestArgs, "customHeaders")
	}

	if _, exists := ob.CustomHeaders["Content-type"]; !exists {
		ob.CustomHeaders["Content-Type"] = "application/json"
	}
}

func (ob *OrsBase) fetchRequest(body interface{}, client *http.Client) (*http.Response, error) {
	orsUtil := OrsUtil{}
	url := orsUtil.PrepareUrl(ob.ArgsCache)

	if ob.ArgsCache[constants.PropNames.Service] == "pois" {
		if strings.Contains(url, "?") {
			url += "&"
		} else {
			url += "?"
		}
	}

	authorization := map[string]string{"Authorization": ob.ArgsCache[constants.PropNames.ApiKey].(string)}

	jsonBody, err := jsoniter.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	for key, value := range authorization {
		req.Header.Set(key, value)
	}

	for key, value := range ob.CustomHeaders {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}

func (ob *OrsBase) createRequest(body interface{}) (interface{}, error) {
	client := &http.Client{Timeout: 5 * time.Second}

	response, err := ob.fetchRequest(body, client)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(response.Status)
	}

	var result interface{}
	if ob.ArgsCache["format"] == "gpx" {
		result, err = io.ReadAll(response.Body)
	} else {
		err = jsoniter.NewDecoder(response.Body).Decode(&result)
	}

	if err != nil {
		return nil, err
	}
	return result, nil

}

// * getBody() is overwritten in directions and isochrones
func (ob *OrsBase) getBody(map[string]interface{}) map[string]interface{} {
	return ob.HttpArgs
}

func (ob *OrsBase) calculate(reqArgs map[string]interface{}) (interface{}, error) {
	util := OrsUtil{}
	ob.RequestArgs = reqArgs
	ob.checkHeaders()
	ob.RequestArgs = util.FillArgs(ob.DefaultArgs, ob.RequestArgs)
	ob.ArgsCache = util.SaveArgsToCache(ob.RequestArgs)
	ob.HttpArgs = util.PrepareRequest(ob.RequestArgs)
	postBody := ob.getBody(ob.HttpArgs)

	return ob.createRequest(postBody)
}
