# 0007 - separators-vs-terminators, 2025-02-16

## Issue

The majority of programming languages use commas to separate expressions in
contexts where a sequence of expressions occurs e.g. parameter lists, argument
lists, initializer lists. However strict separators, as opposed to terminators,
mean that adding/deleting to the start/end of the list requires adjusting the 
following/preceding comma.

To solve this issue some modern programming languages also allow commas to be
used as a terminator. This means that all the expressions including the start/end
have the same structure. However, Monotype has a much looser grammar than a
typical programming language and such optional behaviour may increase the scope
for unwanted ambiguity.


## Factors

- Reducing ambiguity
- Increasing teachability
- Commas-as-separators is always allowed in mainstream programming languages
- Commas-as-terminators is sometimes allowed
- Semi-colons-as-terminators is always allowed
- Semi-colons-as separators is sometimes allowed
- Programmers can object to mandatory terminators 
- I plan to allow newlines to replace semi-colons


## Options and Outcome

- Commas as separators, semi-colons as terminators
- Allow commas and semi-colons to be optional terminators
- A hybrid

Outcome: Commas as separators, semi-colons as terminators

## Consequences

- Slightly stricter grammar
- Easy to teach
- 

## Pros and Cons of Options

### Option 1

- Pros
  - Straightforward to explain the policy.
  - Conforms to C/C++/C# in the majority of contexts and precisely to Java/Go.
  - Other programming languages include it.

- Cons
  - Adjusting first/last expression in a comma-separated sequence is different
    from others.
  - Python 1-tuple hack `(EXPR,)` must be done with semi-colon `(EXPR;)` or 
    some other device.

- Interesting
  - Semi-colons are allowed in all contexts where commas are allowed so the
    programmer has a choice.
  - No extra verbosity as long as newlines->semi-colons is supported.

### Option 2

- Pros
  - Majority of programming languages support trailing commas in some context
  - Possible to arrange for every expression in an argument list to be in the
    same format of `EXPR ,`.
  - Supports the Python 1-tuple hack `(EXPR,)`

- Cons
  - Accidentally deleted expressions (esp. 1 letter variable names) may
    still compile.

- Interesting
  - No option can be terser

### Option 3

- Cons
  - Least simple to teach
  - Will have some ambuiguity

