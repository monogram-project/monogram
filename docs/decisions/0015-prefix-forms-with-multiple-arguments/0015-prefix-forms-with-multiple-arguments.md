# 0015 - Prefix forms with multiple arguments, 2025-04-20

## Issue

We would like prefix forms to cover the following scenarios:

- `return! EXPR` as in C, C#, Java, Python, Go, ...
- `return! x, y` as in Python, Go
- `while! (EXPR) {EXPR}` as in C, C#, Java...
- `while! EXPR {EXPR}` as in Swift

However the decision to make `EXPR{ ARGS }` an example of a function-apply,
which interferes with these scenarios. Removing this interpretation of
function-apply makes prefix form much easier to use at the expense of the
symmetrical treatment of the three types of delimiters.

To illustrate this interaction, here is the current XML parse of `while! (x) {y}`:
```xml
<unit>
  <form syntax="prefix">
    <part keyword="while">
      <apply kind="braces" separator="undefined">
        <delimited kind="parentheses" separator="undefined">
          <identifier name="x" />
        </delimited>
        <arguments>
          <identifier name="y" />
        </arguments>
      </apply>
    </part>
  </form>
</unit>
```

Although this does parse, it is not the intended form. Of course it would 
be possible to disentangle this and transform it into the following.

```xml
<unit>
  <form syntax="prefix">
    <part keyword="while">
      <delimited kind="parentheses" separator="undefined">
        <identifier name="x" />
      </delimited>
    </part>
    <part keyword="_">
      <delimited kind="braces" separator="undefined">
        <identifier name="y" />
      </delimited>
    </part>
  </form>
</unit>
```

However, the need to disentangle forces an artifical distinction between prefix
forms and surround forms, complicates downstream programming and adds an
interpretative ambiguity the notation did not have before (is it a 1 or 2 part
form)?


## Factors

- Simplicity of learning
- Familiarity to programmers

## Options

- Option 1: Leave as-is, with all three delimiters being identical, and require
  devs to untangle the AST
- Option 2: Disallow `EXPR{ARGS}`, breaking symmetry between delimiters, and
  allow prefix forms to handle multiple arguments

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
  - Familiar to programmers coming to Monogram
  - Straightforward to use in the scenarios listed at the start
- Cons
  - The asymmetry needs to be learnt
- Interesting
  - There is a risk of runaway consumption of input when arguments are
    simply justaposed.
  - This is addressed by not reading past line-break terminated expressions.
  - We would like `while E: S endwhile` and `while! (E) {S}` to build the 
    same shaped tree - although the latter will include the extra brackets.

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

