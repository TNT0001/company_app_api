package infrastructure

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"go-api/pkg/shared/utils"
)

var (
	// ConfigCommonFile is common config file prefix.
	ConfigCommonFile = "app"
	// ConfigTypeDefault is default config type
	ConfigTypeDefault = "json"
)

// NewConfig to read config
func NewConfig() error {
	configPath := utils.GetStringFlag("dirConfig")
	viper.AddConfigPath(configPath) // path to look for the config file in
	viper.SetConfigName(ConfigCommonFile)
	viper.SetConfigType(ConfigTypeDefault) // viper.SetConfigType("YAML")としてもよい
	viper.AutomaticEnv()
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {
		// return wraperror.WithTrace(err, wraperror.Fields{"ConfigPath": ConfigPath, "ConfigCommonFile": ConfigCommonFile}, nil)
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("ConfigHandler file changed:", e.Name)
	})

	return nil
}

// SetConfig set value to config file.
func SetConfig(key string, value interface{}) {
	viper.Set(key, value)
}

// GetConfigString get string from config file.
func GetConfigString(key string) string {
	return viper.GetString(key)
}

// GetConfigInt get int from config file.
func GetConfigInt(key string) int {
	return viper.GetInt(key)
}

// GetConfigInt64 get int64 from config file.
func GetConfigInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetConfigBool get bool from config file.
func GetConfigBool(key string) bool {
	return viper.GetBool(key)
}

// GetConfigStringMap get bool from config file.
func GetConfigStringMap(key string) interface{} {
	return viper.GetStringMap(key)
}

// GetConfigByte get []byte from config file.
func GetConfigByte(key string) []byte {
	return []byte(viper.GetString(key))
}

// GetConfigEnv from env file.
func GetConfigEnv(key string) string {
	return os.Getenv(key)
}
