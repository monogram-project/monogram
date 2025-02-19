# 0008 - Strings with single, double and backtick quote-marks, 2025-02-18

## Issue

Programming languages have three quote marks: the single quote, the double
quote and the backtick. The string and double quotes are typically used as
alternative string quotes but the backtick gets a range of uses:

- Template parameters (Javascript)
- Interpolating the results of a shell command (PHP, Ruby, Bash)
- Raw string literals (Go)
- Escaping keywords (Swift)

What is the proper way to handle these in Monogram?

## Factors

- Ease of learning
- Ease of recall
- Flexibility of syntax for different intentions

## Options

1. Treat all three as string-like quotes e.g. "foo", 'foo' and `foo` and the
   syntax tree simply records which quote was employed.
2. Reserve single quotes for characters, double quotes for strings, backtticks
   for interpolation.

## Outcome and Consequences

We choose Option 1: all three are quotes for string-like tokens. This creates
a lot of potential redundancy that has to be handled by tree-builder/analyser.

## Pros and Cons of Options

### Option 1

- Pros
  - Easy to learn and recall: when all three are treated the same way there is
    much less to remember
  - The syntax does not commit to any particular usage (e.g. shell interpolation)
    and so is very flexible
- Cons
  - Any interpretation of the differences is handled by the post-processor.
- Interesting
  - It is a parallel solution to the way brackets are handled.

### Option 2

- Pros
  - Matching syntax onto existing uses helps with familiarity for experienced
    programmers
- Cons
  - But only for programmers from a particular school of languages
  - Getting away from pre-determined meanings is at the core of Monogram