package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRenderResources(t *testing.T) {
	trueRef := true
	cfg := &SiteConfig{
		ProjectKey:   "key",
		ClientSecret: "${data.sops.values[\"my-secret\"]}",
		ProjectSettings: &CommercetoolsProjectSettings{
			Countries: []string{"NL", "DE"},
		},
		Frontend: &CommercetoolsFrontendSettings{
			CreateCredentials: &trueRef,
		},
		TaxCategories: []CommercetoolsTaxCategory{
			{
				Key:  "low",
				Name: "Low Tax",
				Rates: []CommercetoolsTax{
					{
						Country:         "NL",
						Amount:          0.8,
						Name:            "Low",
						IncludedInPrice: &trueRef,
					},
				},
			},
		},
		Zones: []CommercetoolsZone{
			{
				Name:        "Primary",
				Description: "Primary zone",
				Locations: []CommercetoolsZoneLocation{
					{
						Country: "NL",
					},
				},
			},
		},
	}
	data, err := renderResources(cfg, "0.1.0")
	require.NoError(t, err)

	assert.Contains(t, data, `client_secret = data.sops.values["my-secret"]`)
}