# 0016 - Balanced Ternary Literals, 2025-04-26

## Issue

Balanced ternary notation (BTN) needed to be incorporated as a numeric literal
in Monogram. A decision was required on the notation structure, particularly how
to represent the three values `1`, `0`, and `-1`, while ensuring the notation
aligns with Monogram's thematic clarity, consistency, and usability.

## Factors

- **Consistency**: Monogram’s numeric literal design for other bases adheres to
  strict conventions.
- **Clarity**: The notation must be intuitive, legible, and teachable, even for
  users unfamiliar with BTN.
- **Legacy**: Existing conventions in BTN include representations like `0, T, +`
  (libraries) and `0, -, +` (educational resources).
- **Thematic Cohesion**: Monogram is opinionated, striving for the "best
  balanced" decision rather than supporting multiple interchangeable options.
- **Practicality**: The notation should align with Monogram’s broader syntax for
  numeric literals and be suitable for programming tasks.

## Options

- Option 1: Adopt `1` for `1`, `T` for `-1`, and `0` for `0`. e.g. `0t1T0`.
- Option 2: Use `+` for `1`, `-` for `-1`, and `0` for `0`, reflecting
  educational conventions. e.g. `0t+-0`.
- Option 3: Use `+` for `1`, `T` for `-1`, and `0` for `0`, based on conventions
  from existing libraries. e.g. `0t+T0`.
- Option 4: Support both notations as synonyms
  - Permit `1` and `+` as interchangeable for `1`, and `T` and `-` as
    interchangeable for `-1`. Optionally require users to stick to one
    convention within a single numeric literal.

## Pros and Cons of Options

### Option 1: Use `1`, `T`, and `0`

- **Pros**:
  - Aligns with library conventions, familiar to developers.
  - Avoids overloading `+` and `-`, maintaining clarity in numeric literals.
  - Consistent with Monogram’s theme of offering a single, clear solution.
- **Cons**:
  - Slightly less intuitive for those educated using `+` and `-`.
- **Interesting**:
  - `T` provides a unique and memorable representation of `-1`.

### Option 2: Use `+`, `-`, and `0`

- **Pros**:
  - Intuitive for users familiar with educational resources on BTN.
- **Cons**:
  - Deviates from Monogram’s established numeric literal conventions.
  - Visually ambiguous with existing `+` and `-` symbols in other contexts.
- **Interesting**:
  - Reflects the historical presentation of BTN in education.

### Option 3: Use `+`, `T`, and `0`

- **Pros**:
  - Mirrors library conventions, making it familiar to those already working
    with BTN in programming.
  - Maintains clarity by avoiding use of `-`, reducing potential ambiguity.
- **Cons**:
  - Using `+` may slightly deviate from Monogram’s design for other bases, which
    does not rely on symbols like `+` in literal notation.
  - Could create a learning curve for those accustomed to numeric literals
    without symbolic prefixes.
- **Interesting**:
  - Provides a direct bridge to existing BTN implementations in libraries.

### Option 4: Support both notations as synonyms

- **Pros**:
  - Provides flexibility, catering to different user preferences.
- **Cons**:
  - Increases complexity in implementation and user documentation.
  - Inconsistent with Monogram’s opinionated design philosophy.
- **Interesting**:
  - Supports legacy and modern conventions without requiring a decision.

## Outcome and Consequences

**Decision**: Adopt **Option 1**: Use `1`, `T`, and `0` to represent BTN values
in Monogram. 

This choice aligns with Monogram's focus on thematic clarity, consistency, and
usability, while honoring existing BTN conventions familiar to developers. It
avoids ambiguity and reduces learning curves for numeric literals across the
system.

## Additional Notes

- Documentation will clearly highlight `T` as representing `-1` to ensure 
  intuitive adoption.
- Examples will be provided for both integer and fractional BTN literals to
  reinforce user understanding.
