# 0018 - Non-Finite Numeric Literals, 2025-04-26

## Issue

Monogram should support non-finite values in numeric literals, including
positive infinity (`∞`, `0n1`), negative infinity (`-∞`, `-0n1`), and nullity
(`⦰`, `0n0`). These values are distinguished between integer and floating-point
domains, where the presence of a fractional point denotes floating-point values
(e.g., `0n1` for integer infinity vs `0n1.0` for floating-point infinity). 

A decision was required on how to handle the conversion of these non-finite
values when the `--decimal` flag is used, ensuring compatibility with other
programming environments like Python and Rust.

## Factors

- **Interoperability**: Python and Rust adhere to the IEEE 754 standard,
  representing non-finite floating-point values as `"inf"`, `"-inf"`, and
  `"nan"`. To ease adoption of Monogram notation, these outputs should be
  compatible with such environments.
- **Clarity in Trans-Integer vs Trans-Real Values**: Monogram distinguishes
  between integer and floating-point domains for non-finite values, but current
  programming environments generally lack this distinction.
- **Practicality**: The `--decimal` option is intended to simplify parsing,
  reducing the need for users to implement additional conversion logic.

## Options

- Option 1: Output `"inf"`, `"-inf"`, and `"nan"` for all non-finite values
  - Convert both trans-integer and trans-real non-finite values into `"inf"`,
    `"-inf"`, and `"nan"` when the `--decimal` flag is used, ensuring full
    compatibility with IEEE 754-compliant languages.

- Option 2: Distinguish between trans-integer and trans-real values
  - Use distinct strings for trans-integer values, such as `"i-inf"` and
    `"i-nan"`, while retaining `"inf"`, `"-inf"`, and `"nan"` for floating-point
    values. This approach preserves the distinction but requires additional
    parsing logic.

- Option 3: Retain Monogram literals for non-finite values
  - Output Monogram's native notation (`∞`, `-∞`, `⦰`) under `--decimal`.
    While this maintains the distinction between domains, it requires users to
    implement conversion logic for compatibility with other languages.

## Pros and Cons of Options

### Option 1: Output `"inf"`, `"-inf"`, and `"nan"` for all non-finite values

- **Pros**:
  - Fully compatible with Python, Rust, and other IEEE 754-compliant languages.
  - Simplifies parsing for adopters, reducing barriers to use.
- **Cons**:
  - Loses the distinction between trans-integer and trans-real domains.
- **Interesting**:
  - Aligns with Monogram’s goal of easing notation adoption across diverse
    environments.

### Option 2: Distinguish between trans-integer and trans-real values

- **Pros**:
  - Preserves the distinction between trans-integer and trans-real values.
- **Cons**:
  - Introduces complexity for users who must parse these distinct strings manually.
  - Could reduce interoperability with existing programming environments.
- **Interesting**:
  - Provides a trailblazing approach to extended arithmetic but risks alienating
    adopters.

### Option 3: Retain Monogram literals for non-finite values

- **Pros**:
  - Maintains Monogram’s native notation and clarity.
- **Cons**:
  - Requires users to implement additional conversion logic.
  - Reduces interoperability with languages lacking Monogram notation support.
- **Interesting**:
  - Reinforces Monogram’s identity but limits adoption in conventional environments.

## Outcome and Consequences

**Decision**: Adopt **Option 1**: Convert non-finite values to `"inf"`,
`"-inf"`, and `"nan"` under the `--decimal` flag.

This decision ensures compatibility with Python, Rust, and other IEEE
754-compliant environments, reducing barriers to adoption. While it sacrifices
the distinction between trans-integer and trans-real values, this can be
documented for users who require advanced parsing or extended arithmetic
support.

## Additional Notes

- Monogram will preserve the distinction between integer and floating-point
  non-finite values (`∞` vs `∞.0` and `0n1` vs `0n1.0`), offering clarity
  for advanced users.
- Documentation will highlight the mapping of Monogram literals to IEEE 754
  standard values under the `--decimal` flag.
