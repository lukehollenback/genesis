# Genesis

![Status: In Development](https://img.shields.io/badge/Status-In%20Development-blue.svg)

A simple [genetic algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm) writting in Go.

Simulates the evolution of generations of creatures with varying traits through breeding,
mutation, and natural selection. The end goal is a creature that has *all* known superior traits,
and *no* known inferior traits, is bred.

_As a note, this was a programming excercise – not a biology excercise. At best, this was a gross
simplification of reality fueled by
[casual reductionism](https://en.wikipedia.org/wiki/Fallacy_of_the_single_cause). At worst, it
was an excercise with no corrolation to reality at all._

Ultimately, this project serves to:

* Demonstrate use of many of Go's key language features.
* Demonstrate a (albeit simple) genetic algorithm at work.

## Example Output

```
$ go run .
--------------------------------------------------------------------------------
God  = 10101100110011110001110111001110 (score = 32)
Adam = 01110101000011111001110000010110 (score = 19)
Eve  = 00101111001010010111001001110000 (score = 12)
--------------------------------------------------------------------------------
Breeding next generation (1) (of 495 children)...
  ↝        Best Child = 00100100000011011001010101000110 (score = 23)
  ↝ Second Best Child = 00101100000011011000110000000110 (score = 22)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (2) (of 642 children)...
  ↝        Best Child = 00101100000011010001010101001100 (score = 25)
  ↝ Second Best Child = 00001100000011010001110101000110 (score = 25)
...Finished.
```

...


```
--------------------------------------------------------------------------------
Breeding next generation (8) (of 963 children)...
  ↝        Best Child = 10101000110011110001110111001110 (score = 31)
  ↝ Second Best Child = 10001100110011110001110101001110 (score = 30)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (9) (of 157 children)...
  ↝        Best Child = 10101100010011110001110111000110 (score = 30)
  ↝ Second Best Child = 10001100110011110001110101001110 (score = 30)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (10) (of 254 children)...
  ↝        Best Child = 10101100110011110001110111001110 (score = 32)
  ↝ Second Best Child = 10101100110011110001110111101110 (score = 31)
...Finished.
--------------------------------------------------------------------------------
Found best possible child after 10 generations.
Best Possible Child = 10101100110011110001110111001110 (score = 32)
```
