terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = ">=0.1.3"
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
