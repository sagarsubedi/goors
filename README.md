### WIP: Please not that the project is still a work in progress and won't work.

# The Go API to consume openrouteservice(s) painlessly!
This library lets you consume the openrouteservice API in Go applications. It allows you to painlessly consume the following services:

- Directions (routing)
- Geocoding | Reverse Geocoding | Structured Geocoding (powered by Pelias)
- Isochrones (accessibility)
- Time-distance matrix
- POIs (points of interest)
- Elevation (linestring or point)
- Optimization

See the examples in the examples folder.

***Note***: In order to use this client, you have to register for a token at openrouteservice. To understand the features of openrouteservice, please don't forget to read the docs. For visualization purposes on the map, please use openrouteservice maps.

## Documentation
This library uses the ORS API for request validation. To understand the input of each API specifically, please check the API Playground that provides interactive documentation.

## Installation and Usage
### Requirements:
- git
- go

## Install the library:
```go
go get github.com/sagarsubedi/openrouteservice-go
```

​
## Use the library:
```go
package main

import (
    "github.com/sagarsubedi/openrouteservice-go"
)

func main() {
    orsDirections := openrouteservice.NewOrsDirections(map[string]interface{}{
        "api_key": "XYZ",
    })
    // ...
}
```
​
## Pair with local openrouteservice instance:
```go
package main

import (
    "github.com/sagarsubedi/openrouteservice-go"
)

func main() {
    orsDirections := openrouteservice.NewOrsDirections(map[string]interface{}{
        "host": "<http://localhost:8082/ors>",
    })
    // ...
}
```
​
## Development Setup
### Clone the openrouteservice-go repository from GitHub into a development environment of your choice.
`git clone <https://github.com/sagarsubedi/openrouteservice-go.git>`
`cd openrouteservice-go`

​
### Make your openrouteservice API key available for tests and examples:
`echo "API_KEY=your_api_key_here" > .env`

​
### Running Tests
To run tests, use the Go testing tool:
`go test ./...`

​
## Commits and versioning
This app uses the commitizen plugin to generate standardized commit types/messages. After applying any change in a feature branch, use git add . and then npm run commit (instead of git commit ...). The plugin standard-version is used to generate changelog entries, version tag, and to bump the app version in go.mod.

## Deployment flow:
Apply the changes in a feature branch and test it locally.
Once the feature is ready, merge it to develop, deploy it to the testing environment.

Checkout in main, merge from develop and use npm run release to generate a release. This will generate a new release commit as well as a git tag and an entry in CHANGELOG.md.

For more details about commitizen and standard-version, see this article.
