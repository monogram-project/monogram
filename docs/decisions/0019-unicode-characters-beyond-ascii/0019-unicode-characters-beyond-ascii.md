# 0019 - Unicode characters beyond ASCII, 2025-05-20

## Issue

Should we expand the characters used in Monogram beyond Latin-1? And if so, what
is the basis on which we make that decision?

## Factors

- Low adoption rates of programming languages that have gone beyond ASCII, let
  alone Latin-1 e.g. APL, J, K, Uiua.
- Difficulty in typing arbitrary Unicode characters.
- Expansion in available syntax for improved readability.
- Broad adoption of UTF-8 in text editors.
- The existing inclusion of symbols for non-finite values, paired with
  ASCII equivalents (`∞`, `0n1` etc).

## Options

- Option 1: Restrict the input to 7-bit ASCII
- Option 2: Allow UTF-8
- Option 3: Allow UTF-8 when paired with ASCII equivalents

## Pros and Cons of Options

### Option 1: Restrict the input to 7-bit ASCII
- Pros
  - Follows established and successful conventions
  - Simplicity and clarity of policy
- Cons
  - Limits syntax expansion
  - Misses the chance to use fun and helpful symbols

### Option 2: Allow UTF-8
- Pros
  - Makes available fun and helpful symbols
- Cons
  - At the expense of a lack of familiarity
  - Awkward to type for a high percentage of programmers

### Option 3: Allow UTF-8 when paired with ASCII equivalents
- Pros
  - Makes available fun and helpful symbols
- Cons
  - Two versions of every symbol is a nuisance to remember
    - So they must be memorable
- Interesting
  - Having an ASCII version means that a VSCode plug-in could perform
    automatic substitution.


## Outcome and Consequences

Option 2 would be a receipe for failure. Option 1 is no fun. Option 3 wins on 
balance as the disadvantage is relatively minor.

