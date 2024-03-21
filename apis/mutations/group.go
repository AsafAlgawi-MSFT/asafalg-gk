package mutations

import (
	"flag"
	"os"
)

var MutationGroupName string = getGroupFromEnvVars()

func init() {
	flag.String("mutation-group-name", "mutations.gatekeeper.sh", "mutations group name")
}

func getGroupFromEnvVars() string {
	value, exists := os.LookupEnv("MUTATIONS_GROUP_NAME")
	if exists {
		return value
	}

	return "mutations.gatekeeper.sh"
}
