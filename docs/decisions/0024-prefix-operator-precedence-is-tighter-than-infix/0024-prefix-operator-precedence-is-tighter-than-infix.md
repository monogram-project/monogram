# 0024 - Prefix operator precedence is tighter than infix, 2025-07-22

## Issue

The existing reference implementation assigned all prefix operators the
same tightest possible precedence (1). This is significantly inconsistent
with C and violates expectations. e.g.

```
+foo.bar
```

Is parsed as if it was `(+foo).bar`:

```xml
<unit>
  <get name="bar">
    <operator name="+" syntax="prefix">
      <identifier name="foo" />
    </operator>
  </get>
</unit>
```

However, we want it to bind like this:

```xml
<unit>
  <operator name="+" syntax="prefix">
    <get name="bar">
      <identifier name="foo" />
    </get>
  </operator>
</unit>
```

This note describes the way prefix and infix operator precedences should 
work to minimise the violation of expectations.

## Factors

- To mimic, reasonable closely in most contexts, the expectations of programmers
  coming from C/C++/C#/Java/JavaScript etc.
- To have relative precedences of two prefix operators be consistent when they
  are both in infix roles.


### Priorities

- Precedence of unary `+` and `-` should be the same
- `+foo()` should bind as `+(foo())` i.e. `(` < `+`
- `+foo[x]` should bind as `+(foo[x])` i.e. `[` < `+`
- `+foo.bar` should bind as `+(foo.bar)` i.e. `.` < `+`
- Otherwise prefix operators should bind tighter than infix operators.

## Outcome

The following precedence order, based on the initial character of the operator.

- infix `.`
- infix `(`
- infix `[`

- prefix `.`

- main sequence
    - prefix `*`
    - prefix `/`
    - prefix `%`
    - prefix `-`
    - prefix `+`
    - prefix `<`
    - prefix `>`
    - prefix `~`
    - prefix `!`
    - prefix `&`
    - prefix `^`
    - prefix `|`
    - prefix `?`
    - prefix `:`
    - prefix `=`

- repeat main sequence for infix versions
    - infix `*`
    - infix `/`
    - infix `%`
    - infix `-`
    - infix `+`
    - infix `<`
    - infix `>`
    - infix `~`
    - infix `!`
    - infix `&`
    - infix `^`
    - infix `|`
    - infix `?`
    - infix `:`
    - infix `=`


