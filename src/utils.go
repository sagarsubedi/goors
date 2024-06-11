package main

import (
	"strings"
)

type OrsUtil struct{}

func (ou *OrsUtil) FillArgs(defaultArgs, requestArgs map[string]interface{}) map[string]interface{} {
	for k, v := range defaultArgs {
		if _, exists := requestArgs[k]; !exists {
			requestArgs[k] = v
		}
	}
	return requestArgs
}

func (ou *OrsUtil) SaveArgsToCache(args map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		constants.PropNames.Host:       args[constants.PropNames.Host],
		constants.PropNames.ApiVersion: args[constants.PropNames.ApiVersion],
		constants.PropNames.Profile:    args[constants.PropNames.Profile],
		constants.PropNames.Format:     args[constants.PropNames.Format],
		constants.PropNames.Service:    args[constants.PropNames.Service],
		constants.PropNames.ApiKey:     args[constants.PropNames.ApiKey],
		constants.PropNames.MimeType:   args[constants.PropNames.MimeType],
	}
}

func (ou *OrsUtil) PrepareRequest(args map[string]interface{}) map[string]interface{} {
	delete(args, constants.PropNames.MimeType)
	delete(args, constants.PropNames.Host)
	delete(args, constants.PropNames.ApiVersion)
	delete(args, constants.PropNames.Service)
	delete(args, constants.PropNames.ApiKey)
	delete(args, constants.PropNames.Profile)
	delete(args, constants.PropNames.Format)
	delete(args, constants.PropNames.Timeout)
	return args
}

func (ou *OrsUtil) PrepareUrl(args map[string]interface{}) string {
	url := args[constants.PropNames.Host].(string)
	urlPathParts := []string{
		args[constants.PropNames.ApiVersion].(string),
		args[constants.PropNames.Service].(string),
		args[constants.PropNames.Profile].(string),
		args[constants.PropNames.Format].(string),
	}

	urlPathPartsStr := strings.Join(urlPathParts, "/")
	urlPathPartsStr = strings.ReplaceAll(urlPathPartsStr, "//", "/")

	// ? The beginning and end of urlPathPartsStr cannot be a slash
	if urlPathPartsStr[0] == '/' {
		urlPathPartsStr = urlPathPartsStr[1:]
	}
	if end := urlPathPartsStr[len(urlPathPartsStr)-1]; end == '/' {
		urlPathPartsStr = urlPathPartsStr[:len(urlPathPartsStr)-1]
	}

	url += "/" + urlPathPartsStr
	return url
}
