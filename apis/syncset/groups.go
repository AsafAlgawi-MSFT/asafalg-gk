package syncset

import (
	"fmt"

	"github.com/open-policy-agent/gatekeeper/v3/apis/disambiguator"
)

var SyncSetGroupNama = fmt.Sprintf("syncset.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
