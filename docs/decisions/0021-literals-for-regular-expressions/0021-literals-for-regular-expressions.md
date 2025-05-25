# 0021 - Literals for Regular Expressions, 2025-05-24

## Issue

Regular expressions are an important category that deserve their own
literal notation. In this decision record we make a specific proposal
to use the unicode symbol `⫽`. Hence `⫽the (cat|dog) sat on the mat\.⫽ 
would turn into:

```xml
<literal type="regex" value="the (cat|dog) sat on the mat\." >
```

Note that the contents between the opening and closing quotes are
_raw_ in the sense that no string escapes are enabled.

## Factors

- Must be mnemonic
- Must have an ASCII fallback option
- Must align with the new literal extension syntax


## Pros and Cons

- Pros
    - The `⫽` symbol is visually similar to the `/` syntax of Perl, which
      is often used to denote such literals.
- Cons
    - Copilot warns that not all fonts support this character. 
      - Our constituency is programmers working in plain text editors,
        typically using programming fonts. So this is likely an acceptable
        situation.
- Interesting
    - The existing literal extension syntax works as an ASCII fallback 
      `@regex\«the (cat|dog) sat on the mat\.»`
    - We only need to ensure the generated AST is identical in both cases.



## Outcome and Consequences

Adopting this syntax does mean that Monogram programmers are increasingly
reliant on being able to type these characters _or_ for the editor to 
perform automatic substitution.
