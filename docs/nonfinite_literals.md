# Non-finite Literals in Monogram

## Symbols

The three non-finite values that we give names to correspond to the 
calculations of `1/0`, `-1/0` and `0/0`. The basic correspondence is:

- `1/0` = `∞`, positive infinity
- `-1/0` = `-∞`, negative infinity
- `0/0` = `⦰`

Depending on the context, these values can be interpreted as meaningless
(standard maths), as floating point infinity and NaN (IEE 754), or as infinity
and nullity as in [trans-real arithmetic](https://doi.org/10.36285/tm.91).

## ASCII Versions

To avoid the need to find these special symbols on standard keyboards we
provide ASCII equivalents.

- `1/0` = `0n1`, positive infinity
- `-1/0` = `-0n1`, negative infinity
- `0/0` = `0n0`

(The mnemonic behind this mapping is that these values work like the numerators
when it comes to multiplication.)

## Floating Point Variants

Note that floating point variants and ASCII alternatives are supported, see
below.

| Value    | Symbol | ASCII version |
| -------- | ------ | ------------- |
| 1/0      | ∞      | 0n1           |
| -1/0     | -∞   | -0n1          |
| 0/0      | ⦰      | 0n0           |
| 1.0/0.0  | ∞.0    | 0n1.0         |
| -1.0/0.0 | -∞.0 | -0n1.0        |
| 0.0/0.0  | ⦰.0    | 0n0.0         |

You might well ask, is there really a difference between positive infinity and
floating point positive infinity? We would answer, if you think there is a
useful difference between (say) integer 2 and floating point 2, then these are
also useful by analogous reasoning.
