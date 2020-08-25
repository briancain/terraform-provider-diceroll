package main

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func resourceDiceRoll() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDiceRollCreate,
		ReadContext:   resourceDiceRollRead,
		UpdateContext: resourceDiceRollUpdate,
		DeleteContext: resourceDiceRollDelete,
		Schema: map[string]*schema.Schema{
			"die": &schema.Schema{
				Sides:    schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceDiceRollCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceDiceRollRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceDiceRollUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceDiceRollDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
