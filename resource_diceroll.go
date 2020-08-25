package main

import (
	//"fmt"
	//"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDiceRoll() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiceRollCreate,
		Read:   resourceDiceRollRead,
		Update: resourceDiceRollUpdate,
		Delete: resourceDiceRollDelete,
		Schema: map[string]*schema.Schema{
			"die_sides": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
		},
	}
}

func resourceDiceRollCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDiceRollRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDiceRollUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDiceRollDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
