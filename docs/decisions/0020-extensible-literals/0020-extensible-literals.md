# 0020 - Extensible Literals, 2025-05-20

## Issue

To support extensible literals of arbitrary type, such as regular expressions,
dates or IP addresses, we need some kind of innovation. For example it would be
helpful to extend strings to indicate the intended type. We already have
something very similar with the `specifier` attribute on multi-line strings.

In this article we explore the options for supporting a specifier on strings
and possibly other data types.

## Factors

- Syntactic elegance
- Syntactic footprint
- Mnemonic quality
- Consistency with existing design
- Escaping: whatever we use to capture the literal text must support the 
  same kind of quoting as strings.
- Impact on the parser
- ASCII pairing if unicode is used.

## Options

In these examples we `date` as the specifier and a random date in ISO 8601 
format as the working example. 

- Option 1a: `date'2025-05-17'` an identifier immediately followed by string,
  which is currently an impossible combination.
- Option 1b: `@date'2025-05-17'` as (1a) but introduced wth `@`.
- Option 1c: `$date'2025-05-17'` as (1a) but introduced wth `$`.
- Option 1d: `#date'2025-05-17'` as (1a) but introduced wth `#`.
- Option 2: `date«2025-05-17»`, with variants (a)-(d)
- Option 3: `@date[2025-05-17]`, with variants (b)-(d)
- Option 4: Adopt `@` as raw-string signifier and « » as alternative string quotes
- Option 5: Do not have extensible literals.

## Pros and Cons of Options

### Option 1a: Identifier followed by a string

- Pros
    - Terse, with zero syntactic footprint.
    - Includes string notation, so all the single-line string variants apply.
    - Escaping is completely solved.
    - Trivial parser-only implementation, extending readIdentifier.

- Cons
    - Quite clunky looking.
    - Vulnerable to single-character typo problems e.g. `data|'foo'`.


### Option 1b: With @ prefix

- Pros
    - Visually distinctive.
    - Mnemonic, as the `@` symbol is often used to indicate attributes.
    - And in C# `@` is used to introduce raw-strings.
    - Elegant implementation possible.

- Cons
    - Still a bit clunky-looking.
    - Big syntactic footprint, using the `@` symbol for a niche role.

- Interesting
    - Can use `@` as an alternative way to denote raw-ness. `So @'\n\n'` 
      would denote a 4-character string.
    - With this interpretation, we would want a way to denote non-raw
      extended literals. Simply fallback to Option 1

### Option 1c: With $ prefix

- Pros
    - Visually distinctive
    - Elegant implementation possible.

- Cons
    - Not at all mnemonic.
    - Still a bit clunky-looking.
    - Big syntactic footprint, using the `$` symbol for a niche role.

- Interesting
    - This option is dominated by (1b)

### Option 1d: With # prefix

- Pros
    - Visually distinctive
    - Elegant implementation possible BUT requires changing comment convention.

- Cons
    - Clashes with comment syntax
    - Otherwise it has a minute syntactic footprint
    - Some mnemonic quality as languages such as Common Lisp use `#` in a
      similar way.
    - Still a bit clunky-looking.

- Interesting
    - Using `#` for nothing but comments is very clear but I do think it 
      is a waste of a flexible symbol.
    - The change would be to require `#` comments to be followed by a space,
      newline, `!`, or two futher hashes. 
    - Other sequences would cause an error.
    - This would have the incidental,  beneficial effect of outlawing comments
      that are glued to the leading `#`, which is visually cluttered.

### Option 2a: Allow guillemets as string quotes

This option requires that we can use « and » as string quotes. Because these
fall outside of the ASCII set we would need them paired. The obvious pairing
is by falling-back to options (1a). 

- Pros
    - Very natural use of these speech marks.
    - Visually distinctive and clean
    - Simple implementation.

- Cons
    - The fallback option means we cannot rate this option above (1a).
    - Syntactic footprint is significant.

### Option 2b-2d: Allow guillemets as string quotes with @, $ or # prefix

This option requires that we can use « and » as string quotes. Because these
fall outside of the ASCII set we would need them paired. The obvious pairing
is by falling-back to options (1b)-(1d) respectively. 

- Pros
    - Very natural use of these speech marks.
    - Visually distinctive and clean
    - Simple implementation.

- Cons
    - Syntactic footprint is very high when teamed with options (1b) and (1c)
      and high when teamed with (1d).

- Interesting
    - Since (1b) dominates (1c) the choice is between `@` and `#` which
      is a straight trade-off between mnemonic quality vs altering 
      end-of-line comments.

### Option 3b-3d: `@date[2025-05-17]`

- Cons
    - Asking people to read `[...]` as string quotes is a horrible conflation
      of roles. Showstopper.

### Option 4L Option 1a but Adopt @ to signify raw-strings and « » as alternative string quotes

This option emerged as I worked on this decision and seemed to have a lot 
going for it. The basic idea is that:

1. We use `@` to signify raw-string literals. This is a good idea anyway as
   reusing `\` is clever but visually confusing. We also retire the leading
   escape. Precedent here is C#.
2. We adopt `« »` as string quotes (using `chevron` as the quote attribute value).
   Again this is a good idea as these quotes are visually distinctive.
3. We adopt the full string syntax as an optional `@` followed by an 
   optional identifier (specifier) followed by a string.

Making raw strings easier to write is key, since regular expressions in 
particular typically need to be written raw.

- Pros
  - Straightforward design with reusable elements
  - The use of chevrons make the separation from the specifier-identifier nice
    and clear.
  - The `@` symbol as a raw-string indicator fixes an existing issue.

- Cons
  - The non-raw fallback to ASCII is clunky: `date"2025-05-25"`. In
    general it will look neater with raw strings `@date"2025-05-25"`.
  - A heavy syntactic footprint which is significantly ameliorated by the
    fact `@` is reusable.

### Option 5: Don't have extensible literals

- Pros
    - Simple to learn
- Cons
    - We force programmers to represent dates, IP addresses, colours as 
      strings and to use function calls to convert them. 
      - But Monogram has no built-in function calls, which pushes the 
        effort back to the developer.


## Outcome and Consequences

Option 4 is my pick of the bunch. Adding small features to accomplish the
job is the way to go, raw-tagged strings look good, an easy implementation and
escaping a fully solved problem.

- `@date«2025-05-25»`

This would be translated into:

```xml
<string quote="chevron" specifier="date" value="2025-05-25" />
```


## Additional Notes

Do extended literals make any sense in combination with string interpolation 
and multi-lines?

The applicability to string interpolation is unclear, since an interpolated
string is something of a compound literal. My view here is that if it is good
for strings it must be potentially good for other things - but `join` must be
sensitive to the target type.

Here I propose that: ``txt«This is my \(thing)»` becomes:

```xml
<join quote="double" specifier="txt">
    <string quote="double" value="This is my " specifier="" />
    <interpolation kind="parentheses">
        <identifier name="thing" />
    </interpolation>
</join>
```

In other words the specifier is floated to the outer level and the interior
strings are simply neutral.

Multiline strings already, in fact, handle specifiers, so the only change would
be to add the `specifier` attribute to the individual lines.
