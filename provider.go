package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns a Terraform provider schema.
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dice_roll": resourceDiceRoll(),
		},
	}
}
