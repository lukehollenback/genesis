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

```bash
$ go run .
--------------------------------------------------------------------------------
God  = 00100110011111110011001010110011 (score = 32)
Adam = 11101100110100100010000100011101 (score = 15)
Eve  = 01111011111111101001101010000001 (score = 19)
--------------------------------------------------------------------------------
Breeding next generation (1) (of 534 children)...
  ↝        Best Child = 11110100011101100011101000110001 (score = 23)
  ↝ Second Best Child = 11101111000111110011001010011001 (score = 23)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (2) (of 12 children)...
  ↝        Best Child = 00110110000111010001101010010001 (score = 24)
  ↝ Second Best Child = 11101001011111110011101010111001 (score = 23)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (3) (of 96 children)...
  ↝        Best Child = 10100010001101110011001010110011 (score = 28)
  ↝ Second Best Child = 10100110011111010011101010110001 (score = 28)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (4) (of 337 children)...
  ↝        Best Child = 00100010011111110011001010110001 (score = 30)
  ↝ Second Best Child = 00100110011100110011001010110011 (score = 30)
...Finished.
--------------------------------------------------------------------------------
Breeding next generation (5) (of 756 children)...
  ↝        Best Child = 00100110011111110011001010110011 (score = 32)
  ↝ Second Best Child = 00100110011101110011001010110011 (score = 31)
...Finished.
--------------------------------------------------------------------------------
Found best possible child after 5 generations.
Best Possible Child = 00100110011111110011001010110011 (score = 32)
```
