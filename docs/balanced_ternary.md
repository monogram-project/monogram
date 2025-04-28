# More About Balanced Ternary Notation

Balanced ternary is a numeral system that uses three digits that represent zero,
one and minus one. In Monogram these are `0`, `1`, and `T`. Here, `T` represents
`-1`, making it a "balanced" system because the digits are symmetric around
zero. This system is useful in certain mathematical and computational contexts
due to its ability to represent negative values without requiring a separate
sign.

## Balanced Ternary Digits

- `0`: Represents zero.
- `1`: Represents positive one.
- `T`: Represents negative one (`-1`).

## Examples of Balanced Ternary Integers

In Monogram we introduce balanced ternary literals with the prefix `0t`. Here
are some examples of integers represented in balanced ternary:

- `0t10` = `1 × 3^1 + 0 × 3^0 = 3`
- `0tT1` = `T × 3^1 + 1 × 3^0 = -3 + 1 = -2`
- `0t1T0` = `1 × 3^2 + T × 3^1 + 0 × 3^0 = 9 - 3 = 6`

## Examples of Balanced Ternary Fractions

Balanced ternary can also represent fractional values. Here are some examples:

- `0t.1` = `1 × 3^-1 = 1/3`
- `0t1.T`= `1 × 3^0 + T × 3^-1 = 1 - 1/3 = 2/3`
- `0tT.01` = `T × 3^0 + 0 × 3^-1 + 1 × 3^-2 = -1 + 1/9 = -8/9`

## Example of Balanced Ternary with Exponents

- `0tTTe-2` = `(T x 3^1 + T x 3^0) x 3^-2 = -4/9`

Note that the exponents are written in familiar decimal notation.

## Why Use Balanced Ternary?

Balanced ternary has unique properties that make it useful in certain
applications:

- It eliminates the need for a separate negative sign.
- Arithmetic operations can be simplified in some contexts.
- It has applications in computer science, particularly in ternary computing
  systems.

Balanced ternary offers an elegant, symmetrical alternative when representing
numbers, both integers and fractions.
