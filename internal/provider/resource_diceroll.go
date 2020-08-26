package diceroll

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDiceRoll() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiceRollCreate,
		Read:   resourceDiceRollRead,
		Update: resourceDiceRollUpdate,
		Delete: resourceDiceRollDelete,
		Schema: map[string]*schema.Schema{
			"dice": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"die": &schema.Schema{
							Type:     schema.TypeList,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"sides": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"quantity": &schema.Schema{
							Type:     schema.TypeInt,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceDiceRollCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(fmt.Sprintf("%x", rand.Int()))
	return resourceDiceRollRead(d, m)
}

func resourceDiceRollRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDiceRollUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceDiceRollRead(d, m)
}

func resourceDiceRollDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
