package expansion

import (
	"flag"
	"os"
)

var ExpansionGroupName = getGroupNameFromEnvVars()

func init() {
	flag.String("expansion-group-name", "expansion.gatekeeper.sh", "expansion group name")
}

func getGroupNameFromEnvVars() string {
	value, exists := os.LookupEnv("EXPANSION_GROUP_NAME")
	if exists {
		return value
	}

	return "expansion.gatekeeper.sh"
}
