package diceroll

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDiceRoll() *schema.Resource {
	return &schema.Resource{
		Create: resourceDiceRollCreate,
		Read:   resourceDiceRollRead,
		Update: schema.Noop,
		Delete: schema.RemoveFromState,
		Schema: map[string]*schema.Schema{
			"seed": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
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
	//dice := d.Get("dice").([]interface{})
	quantity := d.Get("quantity").(int)
	seed := d.Get("seed").(string)
	rand := NewRand(seed)
	result := make([]interface{}, 0, quantity)

	//for _, die := range dice {
	//	dc := die.(map[string]interface{})
	//  dv := dc["die"].([]interface{})[0]

	//}
	d.SetId(fmt.Sprintf("%x", rand.Int()))
	d.Set("result", result)
	return resourceDiceRollRead(d, m)
}

func resourceDiceRollRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
