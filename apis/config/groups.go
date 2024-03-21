package config

import (
	"flag"
	"os"
)

var ConfigGroupName = getGroupFromEnvVars()

func init() {
	flag.String("config-group-name", "config.gatekeeper.sh", "config group name")
}

func getGroupFromEnvVars() string {
	value, exists := os.LookupEnv("CONFIG_GROUP_NAME")
	if exists {
		return value
	}

	return "config.gatekeeper.sh"
}
