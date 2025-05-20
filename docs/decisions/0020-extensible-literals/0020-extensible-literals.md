# 0020 - Extensible Literals, 2025-05-20

## Issue

To support extensible literals of arbitrary type, such as dates or IP addresses,
it would be helpful to have another pair of brackets - but ones that were not
already in wide use. This would allow a syntax such as: `TYPE(VALUE-TEXT)`. 
This really means borrowing some alternative Unicode brackets that lie
outside of the ASCII range. However, ASCII alternatives must be paired with
these (and always available).


Possible candidates would be:

| Title                 | Open        | Close       | Example            | ASCII                |
| --------------------- | ----------- | ----------- | ------------------ | -------------------- |
| Angle Brackets        | 〈 (U+2329) | 〉 (U+232A) | date〈2025-05-20〉 | date<_2025-05-20_> |
| Double Angle Brackets | « (U+00AB)  | » (U+00BB)  | date«2025-05-20»   | date<<2025-05-20>>   |
| White Square Brackets | ⟦ (U+27E6)  | ⟧ (U+27E7)  | date⟦2025-05-20⟧   | date[\|2025-05-20\|] |
| White Parentheses | ⦅ (U+2985)  | ⦆ (U+2986)  | date⦅2025-05-20⦆   | date(\|2025-05-20\|) |

The specific proposal is to write literals in the following two styles:

- `date⟦ 2025-05-20 ⟧`, the type of literal being explicitly stated.
- `⟦ 192.168.0.1 ⟧`, the type not being stated and decision as to what the 
  literal denotes is deferred.

The first of these would be rendered in XML as:

```xml
<literal type="date" value="2025-05-20" />
```

and the second as follows:

```xml
<literal type="_" value="192.168.0.1" />
```

In this second example the default type would be `_` by default but could be
overidden by the `--default-type=TYPENAME` option or the more elaborate option:

```
--match-type=TYPENAME,REGEX
```

If the `value` matches the REGEX of a given rule then the TYPENAME is inferred.
Matches would be performed in the order of the option and the first match wins.


## Factors

- Impact on the parser
- ASCII pairing
- Visibility
- Mnemonic role

## Options

- Option 1: Angle Brackets
- Option 2: Double Angle Brackets
- Option 3: White Square Brackets
- Option 4: White Parentheses
- Option 5: Don't have extensible literals, the vast majority of mainstream
  programming languages get by with just procedure calls.

## Pros and Cons of Options

### Option 1: Angle Brackets

- Cons
    - Visually indistinct from `<` and `>`
    - Conflicts with the planned XML tag syntax
    - No memorable ASCII pairing
    - No mnemonic relevance



### Option 2: Double Angle Brackets
- Pros
    - Visually distinctive
    - Mnemonic - these are speech marks in several European languages
    - Has a strong ASCII pairing
- Cons
    - `<<` already has well-defined prefix and infix roles
        - Showstopper.

### Option 3: White Square Brackets

- Pros
    - Visually distinctive
    - Good ASCII pairing
- Cons
    - Possible to implement in the parser without backtracking. But not efficient.
- Interesting
    - Weakly mnemonic, used to denote feature in linguistics and closed intervals
      in maths e.g. `⟦0, 1⟧`. Both of these are suitable uses for extended
      literals.



### Option 4: White Parentheses
- Pros
    - Good ASCII pairing
- Cons
    - Existing fonts make it too confusable with standard parentheses.
    - No efficient algorithm possible. Although backtracking can be avoided.
    - No mnemonic aspect.

### Option 5: Don't have extensible literals

- Pros
    - Simple to learn
- Cons
    - We force programmers to represent dates, IP addresses, colours as 
      strings and to use function calls to convert them. 
      - But Monogram has no built-in function calls, which pushes the 
        effort back to the developer.


## Outcome and Consequences

Option 5 is viable but timid and fails to embrace the purely syntactic
nature of Monogram. Option 2 would be the clear winner if `<<` and `>>`
could be repurposed. 

Option 3 is selected as it is the best option without giving up.


## Additional Notes

### Distinctiveness of alternative Unicode bracket characters

It is worth noting that the lack of distinction between (say) angle brackets
and the `<` and `>` symbols might be addressed by a wise font choice. But
making Monogram succeed on the basis of selecting custom fonts is not an
option we should entertain.


### The difficulty in parsing

The issue we have is that, using recursive descent, we cannot resolve whether
`|` is part of the bracket or a prefix operator `|` in the below sequence of
tokens.

```
f[|long_expression ....
```

Obviously we can try one and then revert to the other if it fails. But that 
leads texts that take an exponential amount of time to parse.

### An approach to parsing

However, if we simply scan ahead linking paired open/close brackets together (and
linking to the following/prior token) then we limit the impact to a
single scan of the input. 

In our reference implementation we perform a full input scan anyway, so this
has a negligible overhead. We simply augment the scan with an FSA with pushdown
stack.

### Future expansion

However, we intend to support the ability to pre-compile a limited grammar. And
the tool uses that limited grammer to allow a bit more readability and a lot
more robustness. In this scenario there would be no need to support arbitrary 
lookahead - apart from this requirement.

On the other hand, in this scenario, it will be relatively rare to support
`|` as a prefix operator. It does occur, for example in Verilog it appears as a 
multi-argument bitwise-or. But such use-cases are not frequent. As a consequence 
the lookahead is only triggered when:

- The compiled grammar includes prefix `|`.
- And includes the use of extended literals.
- And `[|` occurs in the text, with no intervening whitespace.

This is not going to be a mainstream scenario, so I think the concern is 
quite modest.

And it is worth adding that table-drive parsers might be able to cope with the
ambiguity without much difficulty anyway.







