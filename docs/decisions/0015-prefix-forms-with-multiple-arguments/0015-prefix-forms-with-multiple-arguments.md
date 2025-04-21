# 0015 - Prefix forms with multiple arguments, 2025-04-20

## Issue

We would like prefix forms to cover the following scenarios:

- `return! EXPR` as in C, C#, Java, Python, Go, ...
- `return! x, y` as in Python, Go
- `while! (EXPR) {EXPR}` as in C, C#, Java...
- `while! EXPR {EXPR}` as in Go

However the decision to make `EXPR{ ARGS }` an example of a function-apply, so
that the three types of delimiters are treated symmetrically, interferes with
these scenarios.

## Factors

- Simplicity of learning
- Familiarity to programmers

## Options

- Option 1: Leave as-is, with all three delimiters being identical, and require
  devs to untangle the AST
- Option 2: Disallow `EXPR{ARGS}`, breaking symmetry between delimiters, and
  allow prefix forms to handle multiple arguments
  - We want different ASTs built for comma-separated arguments and
    whitespace separated arguments.
  - Resolve this by making comma-separated arguments yield a single part with
    multiple children and whitespace separated arguments yield multiple parts
    with single children.

## Pros and Cons of Options

### Option 1: Leave as-is

- Pros
  - Complete symmetry is easy to teach and understand
- Cons
  - The idea that `f{x, y}` is a legal form is unfamiliar to programmers
    approaching Monogram
  - Makes it very awkward to use `{ ... }` to imitate lexical blocks

### Option 2: f{...} no longer supported

- Pros
  - Familiar to programmers coming to Mongram
  - Straightforward to use in the scenarios listed at the start
- Cons
  - The asymmetry needs to be learnt
- Interesting
  - There is a risk of runaway consumption of input when arguments are
    simply justaposed.
  - This is addressed by requiring _multiple_ exclamation marks, each mark
    allows one more arguments to be consumed.
  - We would like `while E: S endwhile` and `while! (E) {S}` to build the 
    same tree. And `return! x, y` to be the same as `return x; y endreturn`.

## Outcome and Consequences

Outcome **Option 2**: in this case we think that the ability to represent
some C-style forms is a significant advantage and the loss of symmetry is
overall out-weighed by this. 

By allowing prefix operators to incorporate simple and compound breakers
we can also more-or-less support C-style `if`:

### Monogram
```txt
if! (x) {
  a += 1
} else-if (y) {
  a += 2
} else: {
  a -= 1
}
```

### XML
```xml
<unit>
  <form syntax="prefix">
    <part keyword="if">
      <delimited kind="parentheses" separator="undefined">
        <identifier name="x" />
      </delimited>
    </part>
    <part keyword="_">
      <delimited kind="braces" separator="undefined">
        <operator name="+=" syntax="infix">
          <identifier name="a" />
          <number value="1" />
        </operator>
      </delimited>
    </part>
    <part keyword="else-if">
      <delimited kind="parentheses" separator="undefined">
        <identifier name="y" />
      </delimited>
    </part>
    <part keyword="_">
      <delimited kind="braces" separator="undefined">
        <operator name="+=" syntax="infix">
          <identifier name="a" />
          <number value="2" />
        </operator>
      </delimited>
    </part>
    <part keyword="else">
      <delimited kind="braces" separator="undefined">
        <operator name="-=" syntax="infix">
          <identifier name="a" />
          <number value="1" />
        </operator>
      </delimited>
    </part>
  </form>
</unit>
```

