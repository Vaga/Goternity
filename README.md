
# Goternity

## What is Eternity II ?

The Eternity II puzzle, aka E2 or E II, is a puzzle competition which was released on 28 July 2007.[1] It was published by Christopher Monckton, and is marketed and copyrighted by TOMY UK Ltd. A $2 million prize was offered for the first complete solution. The competition ended at noon on 31 December 2010, with no solution being found.
[Wikipedia - Eternity II Puzzle](https://en.wikipedia.org/wiki/Eternity_II_puzzle)

## Install

```
$ cd PATH/goternity
$ go get
$ go run main.go
```

## Help

```
Usage of ./goternity:
  -in="new": Input file : [name].goternity
  -out="result.goternity": Output file : [name].goternity
  -render="render.png": Render file : [name].png
```

## Alogrithm

```
Initialize population with random candidate solutions
Evaluate each candidate     

while termination criterion has not been reached
{
    Selection;
    Reproduction;
    Crossover;
    Mutation;
    Evaluation;
}
```
