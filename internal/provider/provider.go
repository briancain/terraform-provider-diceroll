package diceroll

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns a Terraform provider schema.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"diceroll_roll": resourceDiceRoll(),
		},
	}
}
