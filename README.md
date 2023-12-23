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

## Day Five

Probs the most difficult day till now, the first time when the brute force is
not enough

Took me some time to refresh traversing the mappings in the loop, kinda like
traversing/reversing a binary tree or a linked list

After passing the test case the submission is running for like 3mins+, so there
has to be an optimization made to allow for the large ranges in the input, e.g.

```plaintext
seed-to-soil map:
0 1894195346 315486903
1184603419 2977305241 40929361
1225532780 597717 4698739
1988113706 1603988885 78481073
679195301 529385087 505408118
1781158512 2285166785 39457705
352613463 2324624490 326581838
1820616217 1738931330 104130014
2066594779 2671974456 78036460
1288754536 1682469958 56461372
1371340411 3442489267 409818101
3341036988 1092718505 511270380
315486903 1857068786 37126560
1924746231 2209682249 49360033
1345215908 2259042282 26124503
2917167497 2651206328 20768128
1230231519 1034793205 57925300
2144631239 3421335965 21153302
2689873172 2750010916 227294325
1974106264 1843061344 14007442
2165784541 5296456 524088631
1288156819 0 597717
2937935625 3018234602 403101363
```

Rather than brute-forcing out all of the combinations, and doing map lookup, it
has to be dynamically done, i.e. the range has to be used for comparison and
rather than store the entire range in memory, just build the single
corresponding number

seednum (num) 123
range 123412345 20 150 - dststart srcstart len

```pseudo
func correspondingnum(num) {
  for dststart, srcstart, len in _ranges {
    if srcstart <= num <= srcstart + len  {
      return dststart + (num - srcstart)
    }
  }
  return num
}
```
