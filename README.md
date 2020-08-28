# terraform-provider-diceroll

An example provider that rolls x,y die.

```hlc
terraform {
  required_providers {
    diceroll = {
      source  = "briancain/diceroll"
      version = ">=0.1.5"
    }
  }
}

resource "diceroll_roll" "yahtzee" {
  quantity = 6
  sides = 6
  seed = "yahtzee!"
}

output "yahtzee_roll" {
  value = diceroll_roll.yahtzee
}
```

## Development

A `Makefile` has been provided for building and installing this terraform provider.o

Run the following command to build the provider:

```
make build
```

It will build the provider binary and drop it in your local directory.

If you wish to install the provider, and try out one of the examples, you'll want to install
it locally:

```
make install
```

Then you will be able to navigate to the example folder and test out a terraform configuration
for this provider.
