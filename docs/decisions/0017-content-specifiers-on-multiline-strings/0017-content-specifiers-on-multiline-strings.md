# 0017 - Content specifiers on multiline strings, 2025-05-02

## Issue

It occurred to me that because a multiline string can start with three
backticks, it was possible to extend Monogram so that it imitated code-fences
in markdown. The extension was simply allow an identifier to be placed 
immediately after the opening triple-quotes. e.g.

``````
example := 
    ```py
    def increment(n):
        return n + 1
    ```
``````

Because this is fundamentally an opportunistic extension of the language it 
might inhibit a more principled extension in the future. So this is a tradeoff
between a cute extension of unknown value versus an impossible to quantify 
future.


## Factors

- Is there a real use-case for this feature?
- Can it be integrated into the AST with minimal impact?


## Options

Option 1: Do not implement.
Option 2: Implement this feature.

## Pros and Cons of Options

### Option 1

Do not implement
- Pros
  - Reserves options for the future
- Cons
  - Missed opportunity


### Option 2

Implement this feature:
- Pros
    - AST is expanded by one attribute `specifier` on `JoinLines`, which is the
      minimal impact possible.
    - The obvious use-case is the inclusion of values from an external syntax.
    - And for generating Markdown
    - The syntax will feel very familiar to any programmers who know markdown.
- Cons
    - It might inhibit future expansion
    - Mitigate against this by requiring the specifier to be an identifier.


## Outcome and Consequences

**Option 2** - the opportunity is good and the risk to futuree expansion
real but modest.

