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

output "yahtzee_roll" {
  value = diceroll_roll.yahtzee
}
