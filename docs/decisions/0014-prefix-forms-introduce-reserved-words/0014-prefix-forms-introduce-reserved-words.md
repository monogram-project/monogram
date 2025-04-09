# 0014 - Prefix forms introduce reserved words, 2025-04-09

## Issue

Any identifier can act as a prefix form if it is followed by `!`, the prefix
marker. Should that identifier be treated as a dynamically discovered 
reserved word.

To give a concrete example - is this a valid monogram text?
```
include! "My file"
include := 999
```

If `include!` marks `include` as reserved then it is definitely invalid.
However, if it does not then the second use of `include` is just an ordinary
identifier.

## Factors

- The syntax should protect against accidents to whatever extent is
  practical with such a free-form language.

- And at the same time it should not make decisions that compromise the
  freedom to use it how developers want.

- Simplicity, consistency and ease of learning is a top priority.


## Options

- Option 1: Mark as reserved.
- Option 2: Allow both uses.

## Pros and Cons of Options

### Option 1

- Pros
    - It keeps the language simpler and easier to learn.
    - Accidentally leaving off the `!` does not lead to the keyword being
      treated as an ordinary identifier, making the language less error prone.

- Cons
    - It takes away some potentially useful freedom to use identifiers as
      both variables and forms.

- Interesting
    - It means that the same limitations that apply to `endendXXX` must now
      be applied to prefix-forms.
    - Note that 

### Option 2

This is simply the flip side of Option 1.


## Outcome

I have selected Option 1, to treat `foo!` as a dynamic discovery of both `foo`
and `endfoo` as an open/close pair of form keywords. This keeps the notation 
simple and easy to explain. The other issues were secondary and well-balanced.
