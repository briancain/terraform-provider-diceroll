# terraform-provider-diceroll

A terraform provider to roll x number of y sided die :game_die:

## Argument Reference

* `description` - (Optional) This field can be used to give a short description
of the die used in your resource.
* `seed` - (Required) This string is used to seed the random number generator
used to roll the die.
* `sides` - (Required, defaults to 1) Can be used to define how many sides
your die will have. There is no real upper limit to this number beyond what
Terraform supports as a `TypeInt`.
* `quantity` - (Optional, defaults to 1) The number of die to roll together
in the resource.

### Computed Arguments

The following values are automatically computed by Terraform given the arguments
above defined in a resource.

* `computed_total` - (Int) This field is the sum of all die rolled for a given resource.
* `result` - (Array of Ints) This field contains the resulting roll for each die
defined inside the resource.

## Usage

Below is a same module that rolls a handful of dice. Yahtzee rolls 6,6 sided die.

```hcl
terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = "0.1.2"
    }
  }
}

resource "diceroll_roll" "yahtzee" {
  quantity = 6
  sides = 6
  seed = "yahtzee!"
}

resource "diceroll_roll" "dnd" {
  quantity = 4
  sides = 20
  seed = "dungeonsanddragons"
}

output "yahtzee_roll" {
  value = diceroll_roll.yahtzee
}

output "dnd_roll" {
  value = diceroll_roll.dnd
}
```

Then, if you wish to roll these dice:

```shell
brian@localghost:example % terraform apply
An execution plan has been generated and is shown below.
Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

  # diceroll_roll.dnd will be created
  + resource "diceroll_roll" "dnd" {
      + calculated_total = (known after apply)
      + id               = (known after apply)
      + quantity         = 4
      + result           = (known after apply)
      + seed             = "dungeonsanddragons"
      + sides            = 20
    }

  # diceroll_roll.yahtzee will be created
  + resource "diceroll_roll" "yahtzee" {
      + calculated_total = (known after apply)
      + id               = (known after apply)
      + quantity         = 6
      + result           = (known after apply)
      + seed             = "yahtzee!"
      + sides            = 6
    }

Plan: 2 to add, 0 to change, 0 to destroy.

Do you want to perform these actions?
  Terraform will perform the actions described above.
  Only 'yes' will be accepted to approve.

  Enter a value: yes

diceroll_roll.dnd: Creating...
diceroll_roll.yahtzee: Creating...
diceroll_roll.dnd: Creation complete after 0s [id=3692300b2a466f57]
diceroll_roll.yahtzee: Creation complete after 0s [id=6ae6a5b3eecadac1]

Apply complete! Resources: 2 added, 0 changed, 0 destroyed.

Outputs:

dnd_roll = {
  "calculated_total" = 48
  "id" = "3692300b2a466f57"
  "quantity" = 4
  "result" = [
    18,
    9,
    5,
    16,
  ]
  "seed" = "dungeonsanddragons"
  "sides" = 20
}
yahtzee_roll = {
  "calculated_total" = 15
  "id" = "6ae6a5b3eecadac1"
  "quantity" = 6
  "result" = [
    5,
    2,
    2,
    1,
    4,
    1,
  ]
  "seed" = "yahtzee!"
  "sides" = 6
```
