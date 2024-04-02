package mutations

import (
	"fmt"

	"github.com/open-policy-agent/gatekeeper/v3/apis/disambiguator"
)

var MutationGroupName string = fmt.Sprintf("mutations.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
