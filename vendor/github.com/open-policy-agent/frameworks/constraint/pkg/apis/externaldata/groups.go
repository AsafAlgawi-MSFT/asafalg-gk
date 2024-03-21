package externaldata

import "os"

var ExternalDataGroupName = getGroupFromEnvVars()

func getGroupFromEnvVars() string {
	value, exists := os.LookupEnv("EXTERNALDATA_GROUP_NAME")
	if exists {
		return value
	}

	return "externaldata.gatekeeper.sh"
}
