package expansion

import (
	"fmt"

	"github.com/open-policy-agent/gatekeeper/v3/apis/disambiguator"
)

var ExpansionGroupName = fmt.Sprintf("expansion.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
