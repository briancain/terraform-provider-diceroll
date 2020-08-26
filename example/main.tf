terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = "0.0.5"
    }
  }
}

resource "dice_roll" "yahtzee" {
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
