# terraform-provider-diceroll

This provider is still a work in progress. These docs will be updated once it is functional.

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
