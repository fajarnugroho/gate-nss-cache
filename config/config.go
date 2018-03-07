package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var appConfig config

type config struct {
	gateNSS nssHTTP
}

type nssHTTP struct {
	hostURL string
	apiKey  string
}

func Load() {
	gateConfigFile := os.Getenv("GATE_CONFIG_FILE")
	viper.SetConfigFile(gateConfigFile)

	viper.ReadInConfig()
	viper.AutomaticEnv()

	gateNSSConfig := nssHTTP{
		hostURL: mustGetString("NSS_HTTP.HOST_URL"),
		apiKey:  mustGetString("NSS_HTTP.API_KEY"),
	}

	appConfig = config{
		gateNSS: gateNSSConfig,
	}
}

func HostURL() string {
	return appConfig.gateNSS.hostURL
}

func ApiKey() string {
	return appConfig.gateNSS.apiKey
}

func UserURL() string {
	return fmt.Sprintf("%s/passwd?token=%s", HostURL(), ApiKey())
}

func mustGetString(key string) string {
	if !viper.IsSet(key) {
		panic(fmt.Sprintf("key %s is not set", key))
	}
	return viper.GetString(key)
}
