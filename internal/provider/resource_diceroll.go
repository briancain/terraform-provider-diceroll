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
				ForceNew: true,
				Optional: true,
			},
			"quantity": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"sides": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
				ForceNew: true,
			},
			"result": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func resourceDiceRollCreate(d *schema.ResourceData, m interface{}) error {
	seed := d.Get("seed").(string)
	rand := NewRand(seed)
	sides := d.Get("sides").(int)
	quantity := d.Get("quantity").(int)
	result := make([]int, quantity)

	for i := 0; i < quantity; i++ {
		r := rand.Intn(sides - 1) + 1 // Dice don't have a 0
		result[i] = r
	}

	d.Set("result", result)
	d.SetId(fmt.Sprintf("%x", rand.Int()))
	return resourceDiceRollRead(d, m)
}

func resourceDiceRollRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
