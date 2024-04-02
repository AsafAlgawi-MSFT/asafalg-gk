package status

import (
	"fmt"

	"github.com/open-policy-agent/gatekeeper/v3/apis/disambiguator"
)

var StatusGroupName string = fmt.Sprintf("status.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
var ConstraintsGroupName = fmt.Sprintf("constraints.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
