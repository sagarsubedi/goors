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

func TestGeneratePayload(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))
	orsElevation := NewOrsElevation(map[string]interface{}{
		constants.PropNames.ApiKey: os.Getenv("API_KEY"),
	})

	payload := orsElevation.GeneratePayload(map[string]interface{}{
		"host":              "host",
		"service":           "service",
		"notUrlConstituent": "notUrlConstituent",
	})

	expected := map[string]interface{}{
		"notUrlConstituent": "notUrlConstituent",
	}

	assert.Equal(t, expected, payload)
}

func TestLineElevation_SetsCorrectService(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))

	orsElevation := NewOrsElevation(map[string]interface{}{
		constants.PropNames.ApiKey: os.Getenv("API_KEY"),
	})

	_, err := orsElevation.LineElevation(map[string]interface{}{
		"format_in": "encodedpolyline",
		"geometry":  "u`rgFswjpAKD",
	})

	assert.NoError(t, err)
	assert.Equal(t, "elevation/line", orsElevation.RequestArgs[constants.PropNames.Service])
}

func TestLineElevation_FailsWithoutParameters(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))

	orsElevation := NewOrsElevation(map[string]interface{}{
		constants.PropNames.ApiKey:  os.Getenv("API_KEY"),
		constants.PropNames.Service: "elevation/line",
	})

	_, err := orsElevation.LineElevation(map[string]interface{}{})
	assert.Error(t, err)
}

func TestPointElevation_SetsCorrectService(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))

	orsElevation := NewOrsElevation(map[string]interface{}{
		constants.PropNames.ApiKey: os.Getenv("API_KEY"),
	})

	_, err := orsElevation.PointElevation(map[string]interface{}{
		"format_in": "point",
		"geometry":  []interface{}{13.331273, 38.10849},
	})

	assert.NoError(t, err)
	assert.Equal(t, "elevation/point", orsElevation.RequestArgs[constants.PropNames.Service])
}

func TestPointElevation_FailsWithoutParameters(t *testing.T) {
	t.Setenv("API_KEY", os.Getenv("API_KEY"))

	orsElevation := NewOrsElevation(map[string]interface{}{
		constants.PropNames.ApiKey:  os.Getenv("API_KEY"),
		constants.PropNames.Service: "elevation/point",
	})

	_, err := orsElevation.PointElevation(map[string]interface{}{})
	assert.Error(t, err)
}
