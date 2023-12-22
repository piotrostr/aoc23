# Advent of Code 2023

lets see how many puzzles I can do before work piles up and I won't have any more playtime

## Day one

easy part one, doing this no copilot so had to remind myself some syntax like reading a file lol

the part two was a bit weird, I found on reddit that
`oneight` should result with `18` rather than just leave it
with an eight, Ill now change my `Calibrate` function

## Day two

A lot simpler than day one, day one is probs one of the harder days in the first
week I have ever done

## Day three

My algorithm was to find the numbers first and then check
each number coordinates (O(2n^2) since two loops for each of the iterations)

I got stuck here for a bit, two issues:

1. ID issue

    The issue with my initial approach (passing the test-case) was that there might
    be duplicate numbers, so ID cannot be the number string (`123`, `5432`, etc),
    but the number with the `(i, j)` pointers to the column and row, so say
    `123-2-3`

2. `'\t'` issue

     My input had another tab at the beginning, instead of

    ```go
    func Foo() {
      input := `some
    multiline
    input
    without
    \t`
      // ...
    }
    ```

    I had

    ```go
    func Bar() {
      input := `some
      multiline
      input
      with accidental
      \t`
      // ...
    }
    ```

    resulting in the input having `'\t'` as the first char at each line

The second part with asterisks was ez tho

## Day Four

This was a fun one, I got stuck at the second part of the question for a little
bit since it required adding a mutable field to the `type Card struct {}`, which
was not mutable, so after seeing that I cannot `Increment()` the `Card.Count` I
had to change the mapping `map[int]Card` to `map[int]*Card` and it was dandy

The algorithm was not optimal, it ran for like 2 seconds for the large input, so
I assume a fun exercise could be to `pprof` it
