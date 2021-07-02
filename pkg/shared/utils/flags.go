package utils

import "flag"

// GetStringFlag func
func GetStringFlag(name string) string {
	return flag.Lookup(name).Value.(flag.Getter).Get().(string)
}

// GetIntFlag func
func GetIntFlag(name string) int {
	return flag.Lookup(name).Value.(flag.Getter).Get().(int)
}

// GetBoolFlag func
func GetBoolFlag(name string) bool {
	return flag.Lookup(name).Value.(flag.Getter).Get().(bool)
}
