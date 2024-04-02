package templates

import (
	"fmt"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/disambiguator"
)

var TemplateGroupName string = fmt.Sprintf("templates.%s%s", disambiguator.Disambiguator, disambiguator.GatekeeprApiSuffix)
