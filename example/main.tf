provider "diceroll" {
  source = "briancain/diceroll"
  version = "0.0.4"
}

resource "dice" "Yahtzee" {
  dice {
    die {
      id = 1
      sides = 6
    }
    quantity = 6
  }
}
