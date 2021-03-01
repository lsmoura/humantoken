# Human Token

HumanToken is a human-friendly token generator.
You can create tokens without ambiguous characters.

To install: `go get github.com/lsmoura/humantoken`

## API Reference

* `Generate(size int, r *rand.Rand)` -- generates a random `size` string. If `r` is `nil`, a new random object based on
  the current time will be generated for you.
* `TokenGenerator` interface has a single method `Generate(size int)`. This method generates a random human token
  just like the `Generate` standalone counterpart. Use one of the available generators to get an implementation of
  the interface
  
### Interface Generators
* `NewGenerator(r *rand.Rand)` generates a new random generator. If `r` is new, a new random generator will be created
  upon the creation of the interface.
* `NewCryptoGenerator()` generates a new generator based on the `crypto/rand` go package. This generator takes no
  parameters.
  
## Inspiration

This package is based on the [human_token](https://github.com/brianhempel/human_token) package by 
[@brianhempel](https://github.com/brianhempel), originally written for the Ruby programming language. 

## Author

This package was written by [Sergio Moura](https://sergio.moura.ca)


## License

Public Domain; no rights reserved.

No restrictions are placed on the use of HumanToken. That freedom also means, of course, that no
warranty of fitness is claimed; use HumanToken at your own risk.

This public domain dedication follows the the CC0 1.0 at https://creativecommons.org/publicdomain/zero/1.0/
