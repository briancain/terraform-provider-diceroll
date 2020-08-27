terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = "0.0.9"
    }
  }
}

resource "diceroll_roll" "yahtzee" {
  quantity = 6
  sides = 6
  seed = "yahtzeetest"
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
