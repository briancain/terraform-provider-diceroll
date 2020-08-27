terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = "0.0.8"
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

  seed = "yahtzeetest"
}

output "yahtzee_roll" {
  value = diceroll_roll.yahtzee
}
