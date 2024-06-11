package main

type Constants struct {
	DefaultAPIVersion   string
	DefaultHost         string
	MissingAPIKeyMsg    string
	BaseUrlConstituents []string
	PropNames           struct {
		ApiKey     string
		Host       string
		Service    string
		ApiVersion string
		MimeType   string
		Profile    string
		Format     string
		Timeout    string
	}
}

var constants = Constants{
	DefaultAPIVersion:   "v2",
	DefaultHost:         "https://api.openrouteservice.org",
	MissingAPIKeyMsg:    "Please add your openrouteservice api_key.",
	BaseUrlConstituents: []string{"host", "service", "api_version", "mime_type"},
	PropNames: struct {
		ApiKey     string
		Host       string
		Service    string
		ApiVersion string
		MimeType   string
		Profile    string
		Format     string
		Timeout    string
	}{
		ApiKey:     "api_key",
		Host:       "host",
		Service:    "service",
		ApiVersion: "api_version",
		MimeType:   "mime_type",
		Profile:    "profile",
		Format:     "format",
		Timeout:    "timeout",
	},
}
