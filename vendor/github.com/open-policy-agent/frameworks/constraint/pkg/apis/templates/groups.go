package templates

import (
	"flag"
	"os"
)

var TemplateGroupName string = getGroupFromEnvVars()

func init() {
	flag.String("template-group-name", "templates.gatekeeper.sh", "Templates Group name")
}

func getGroupFromEnvVars() string {
	value, exists := os.LookupEnv("TEMPLATE_GROUP_NAME")
	if exists {
		return value
	}

	return "templates.gatekeeper.sh"
}
