terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = "0.0.7"
    }
  }
}

resource "diceroll_roll" "yahtzee" {
  dice {
    die {
      id    = 1
      sides = 6
    }
    quantity = 6
  }
}

output "roll_dice" {
  value = dice.yahtzee
}
