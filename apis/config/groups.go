package config

import (
	"fmt"

	"github.com/open-policy-agent/gatekeeper/v3/apis/disambiguator"
)

var ConfigGroupName = fmt.Sprintf("config.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
