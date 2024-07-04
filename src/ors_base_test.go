package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func TestConstruction_PassesAPIKey(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))
	key := os.Getenv("API_KEY")
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey: key,
	})
	assert.Contains(t, base.DefaultArgs, constants.PropNames.ApiKey)
	assert.Equal(t, key, base.DefaultArgs[constants.PropNames.ApiKey])
}

func TestConstruction_FailsWithoutAPIKey(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, constants.MissingAPIKeyMsg, r.(error).Error())
		}
	}()
	NewOrsBase(map[string]interface{}{})
	t.Errorf("The code did not panic")
}

func TestConstruction_PassesHost(t *testing.T) {
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey: "test",
		constants.PropNames.Host:   "localhost:8080",
	})

	assert.Contains(t, base.DefaultArgs, constants.PropNames.Host)
	assert.Equal(t, "localhost:8080", base.DefaultArgs[constants.PropNames.Host])
}

func TestConstruction_PassesService(t *testing.T) {
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey:  "test",
		constants.PropNames.Service: "test-service",
	})

	assert.Contains(t, base.DefaultArgs, constants.PropNames.Service)
	assert.Equal(t, "test-service", base.DefaultArgs[constants.PropNames.Service])
}

func TestSetRequestDefaults_SetsDefaultArgsFromArgs(t *testing.T) {
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey: "test",
	})
	base.setRequestDefaults(map[string]interface{}{
		constants.PropNames.ApiKey:  "test",
		constants.PropNames.Service: "service",
		constants.PropNames.Host:    "host",
	})

	assert.Equal(t, "service", base.DefaultArgs[constants.PropNames.Service])
	assert.Equal(t, "host", base.DefaultArgs[constants.PropNames.Host])
}

func TestSetRequestDefaults_SetsDefaultHostWhenNotInArgs(t *testing.T) {
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey: "test",
	})
	base.setRequestDefaults(map[string]interface{}{
		constants.PropNames.ApiKey: "test",
	})

	assert.Equal(t, constants.DefaultHost, base.DefaultArgs[constants.PropNames.Host])
}

func TestCheckHeaders_SetsCustomHeaders(t *testing.T) {
	base := NewOrsBase(map[string]interface{}{
		constants.PropNames.ApiKey: "test",
	})
	base.RequestArgs = map[string]interface{}{
		"customHeaders": map[string]string{"customHeader": "value"},
	}
	base.checkHeaders()

	assert.NotEmpty(t, base.CustomHeaders)
	_, exists := base.RequestArgs["customHeaders"]
	assert.False(t, exists)
}
