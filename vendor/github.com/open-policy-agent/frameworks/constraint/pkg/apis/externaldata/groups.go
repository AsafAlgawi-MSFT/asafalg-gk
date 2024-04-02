package externaldata

import (
	"fmt"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/disambiguator"
)

var ExternalDataGroupName = fmt.Sprintf("externaldata.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
