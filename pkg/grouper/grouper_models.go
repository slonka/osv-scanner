// Deprecated: this is now private and should not be used outside the scanner
package grouper

import (
	"strings"

	"github.com/google/osv-scanner/pkg/models"
)

// Deprecated: this is now private and should not be used outside the scanner
type IDAliases struct {
	ID      string
	Aliases []string
}

// Deprecated: this is now private and should not be used outside the scanner
func ConvertVulnerabilityToIDAliases(c []models.Vulnerability) []IDAliases {
	output := []IDAliases{}
	for _, v := range c {
		idAliases := IDAliases{
			ID:      v.ID,
			Aliases: v.Aliases,
		}

		// For Debian Security Advisory data,
		// all related CVEs should be bundled together, as they are part of this DSA.
		// TODO(gongh@): Revisit and provide a universal way to handle all Linux distro advisories.
		if strings.Split(v.ID, "-")[0] == "DSA" {
			idAliases.Aliases = append(idAliases.Aliases, v.Related...)
		}

		output = append(output, idAliases)
	}

	return output
}
