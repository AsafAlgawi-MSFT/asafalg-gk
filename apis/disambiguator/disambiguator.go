package disambiguator

import (
	"fmt"
	"os"
)

const GatekeeprApiSuffix = "gatekeeper.sh"

var Disambiguator string = readDisambiguatorFromEnvVars()

func readDisambiguatorFromEnvVars() string {
	value, exists := os.LookupEnv("GATEKEEPER_API_NAME_DISAMBIGUATOR")
	if exists {
		return fmt.Sprintf("%s-", value)
	}

	return ""
}
