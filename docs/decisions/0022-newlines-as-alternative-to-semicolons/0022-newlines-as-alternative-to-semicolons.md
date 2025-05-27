# 0021 - Newlines as alternative to semicolons, 2025-05-22

## Issue

If we use braces to enclose a series of statements they need to be separated
with semi-colons - unlike forms where statements may also be separated with
newlines as well as semi-colons.

Allowed:
```
{
    a = 1;
    b = 2;
}
```

Disallowed:
```
{
    a = 1
    b = 2
}
```

I would like to allow the same coding style within both contexts without
compromising the ability to use braces for expressions - especially JSON 
expressions.


## Factors

- Compatibility with JSON expressions
- Alignment of braces-content with form-content
  - Simplicity, teachability, consistency
- Alignment of parenthetical-content with bracketed-content with braces-content.
- Gotchas
  - Where you intend the content to be treated as an expression spread 
    over multiple lines but it gets unexpectedly treated as statements.

An example of this 'gotcha' is:

```
{
    1
    -2
}
```

Should this be equivalent to `{1-2}` or `{1;2;}`?

## Options

- Option 1: Continue allowing semi-colons but not newlines.
- Option 2: Allow both semi-colons and newlines as well as commas but:
  - Only allow semi-colons OR newlines OR commas
  - Insist that forms only use semi-colons OR newlines
  - In both cases flag which was used with the `separator` attribute with values
    - semicolon
    - comma
    - newline
    - undefined
  - A backslash at the end of the line (possibly followed by whitespace) will
    prevent the newline acting as a semi-colon.
- Option 3: As Option 2 but also allow parentheses and brackets to use 
  newlines as a separator.
- Option 4: As Option 3 but also allow forms to use commas as a separator.

## Pros and Cons of Options

### Option 1
- Pros
  - Keeps `()` and `[]` aligned with `{}`
  - JSON expressions work
  - No gotchas
- Cons
  - Inconsistent with role of braces as compound-statement.


### Option 2
- Pros
  - Consistent with role of braces as compound-statement.
  - JSON expressions work because the only ambiguous case is when there
    is a preview/infix operator and JSON does not have these!
- Cons
  - `()` and `[]` contents distinct from `{}`
  - A single "gotcha" - shown at the start of this record whose conditions
    are:
    - Single expression enclosed by braces 
    - Spread over multiple lines
    - One of those lines starts with an operator (ignoring whitespace)


### Option 3

- Pros
  - Keeps `()` and `[]` content aligned with `{}`
  - Consistent with role of braces as compound-statement.
  - JSON expressions work because the only ambiguous case is when there
    is a preview/infix operator and JSON does not have these!
- Cons
  - Unfamiliar use of newlines in expressions, see below.
    - However the addition of a `kind` option makes it easy for
      our users to catch unwanted forms like this.
  - A single "gotcha" - shown at the start of this text.
- Interesting
  - Allowing semicolons 'feels' modern.

This option allows for somewhat odd-looking expressions like:
```
x.f(
  y
  z
)
```

But also for fairly intuitive expressions such as:

```
list = [
  "a"
  "b"
  "c"
]
```

### Option 4
- Pros
  - Keeps `()` and `[]` content aligned with `{}`.
  - Aligns form-content with all delimiters content.
  - Consistent with role of braces as compound-statement.
  - JSON expressions work because the only ambiguous case is when there
    is a preview/infix operator and JSON does not have these!
- Cons
  - Unfamiliar use of newlines in expressions, see above.
  - Unfamiliar use of commas in forms, see below.
    - However the addition of a `kind` option makes it easy for
      our users to catch unwanted forms like this.
  - A single "gotcha" - shown at the start of this text.


This option allows for somewhat odd-looking statement sequences like:

```
while T do:
    x = f(x), T = x > 0
endwhile
```


## Outcome and Consequences

To my complete surprise I find myself persuaded by my own arguments that Option
4 is the most attractive. The idea of adding `kind` to forms and distinguishing
the kind of separator as `comma`, `semicolon`, `newline` or `unknown` makes a
lot of sense. Post-parse filtering is expected to handle this perfectly well.

As far as teachability goes, the basic rule is that a newline will terminate any
expression that can be terminated. Delimited expressions cannot be terminated
this way but any expression with operators can be. So if you want to continue
across multiple lines then make sure that each line finishes with an operator
(or use the end-of-line backslash).

## Additional Notes

Insisting on semi-colons to separate statements comes across as quite
old-fashioned to me. Languages such as Go and Swift and Lua have embraced
newlines, as well as scripting languages such as Python, Ruby and Lua. But this
is a subjective factor that I have given very little weight although it was
part of the original motivation.
