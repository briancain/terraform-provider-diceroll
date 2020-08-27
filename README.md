# terraform-provider-diceroll

An example provider that rolls x,y die.

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
