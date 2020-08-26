provider "diceroll" {
  source  = "briancain/diceroll"
  version = "0.0.5"
}

resource "dice" "yahtzee" {
  dice {
    die {
      id    = 1
      sides = 6
    }
    quantity = 6
  }
}

output "dice_roll" {
  value = dice.yahtzee
}
