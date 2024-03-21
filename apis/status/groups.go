package status

import (
	"flag"
	"os"
)

var StatusGroupName string = getStatusGroupNameFromEnvVars()
var ConstraintGroupName string = getConstraintGroupNameFromEnvVars()

func getStatusGroupNameFromEnvVars() string {
	value, exists := os.LookupEnv("STATUS_GROUP_NAME")
	if exists {
		return value
	}

	return "status.gatekeeper.sh"
}

func getConstraintGroupNameFromEnvVars() string {
	value, exists := os.LookupEnv("CONSTRAINTS_GROUP_NAME")
	if exists {
		return value
	}

	return "constraints.gatekeeper.sh"
}

func init() {
	flag.String("status-group-name", "status.gatekeeper.sh", "Status group name")
}
