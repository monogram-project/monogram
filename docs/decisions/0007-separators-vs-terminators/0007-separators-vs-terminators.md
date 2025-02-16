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
typical programming language and such optional behaviour increase the scope
for unwanted ambiguity.




## Factors
List the factors that should be considered...

## Options and Outcome
List possible options considered and the outcome of the decision...

## Consequences
The impact of the decision...

## Pros and Cons of Options

### Option 1
- Pros
- Cons
- Interesting

### Option 2
Etc ...

## Additional Notes
