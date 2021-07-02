package utils

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	stringName = "string"
	intName    = "int"
	boolName   = "bool"
)

// TestGetStringFlag func
func TestGetStringFlag(t *testing.T) {
	expected := ""
	flag.String(stringName, expected, "flag.StringVar")
	received := GetStringFlag(stringName)

	assert.Equal(t, expected, received)
}

// TestGetIntFlag func
func TestGetIntFlag(t *testing.T) {
	expected := 0
	flag.Int(intName, expected, "flag.IntVar")
	received := GetIntFlag(intName)

	assert.Equal(t, expected, received)
}

// TestGetBoolFlag func
func TestGetBoolFlag(t *testing.T) {
	expected := false
	flag.Bool(boolName, expected, "flag.BoolVar")
	received := GetBoolFlag(boolName)

	assert.Equal(t, expected, received)
}
