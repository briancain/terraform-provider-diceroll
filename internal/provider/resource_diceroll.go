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
		Update: schema.Noop,
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
										Optional: true,
									},
									"sides": &schema.Schema{
										Type:     schema.TypeInt,
										Required: true,
									},
									"description": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"result": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
	//dice := d.Get("dice").([]interface{})

	//for _, die := range dice {
	//	dc := die.(map[string]interface{})
	//}
	return nil
}

func resourceDiceRollDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
