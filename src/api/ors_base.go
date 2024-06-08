package utils

import (
	"bytes"
	"errors"
	"net/http"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

type OrsBase struct {
	DefaultArgs   map[string]interface{}
	RequestArgs   map[string]interface{}
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

	// TODO: call setRequestDefaults
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
