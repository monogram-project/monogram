#  Tokenisation in Monogram

In common with most programming languages, Mongram separates analysing the input
into two levels: tokenisation and then parsing the token stream. On this page 
we describe the different kinds of tokens that Monogram recognises and try
to explain the thinking behind some of the less familiar-looking tokens.

## Comments

Monogram supports `#` as the end of line comment character in the same way
as Python, Ruby and Bash. These are simply discarded by a Monogram processor.


## The different types of tokens

- Numbers: positive and negative integers and floats in decimal notation.
  - It is also possible to write these in other bases apart from base 10. 
  - Hex, binary and octal can be written with the generally familiar `0x`, `0b`
    and `0o` prefixes. 
  - Alternatively the `r` for radix indicator can be used to express any
    base from 2 to 36. 
    - To illustrate this, binary numbers can be written with `2r` e.g. 10 can be 
      written either as `0b1010` or `2r1010`.
    - Hex value `0xFFFF` can also be written `16rFFFF`.
    - Bases greater than 16 simply use more upper-case letters of the alphabet.
      e.g. 36rHELLO = 29234652.
  - It is possible to write floating point numbers in these non-decimal bases.
    - For example `0x1.1e2` = 272.0. 
    - Note that both the radix part and the exponent part are written in
      decimal notation.
  - A special case is [balanced
    ternary](https://en.wikipedia.org/wiki/Balanced_ternary) notation. This is
    introduced with the prefix `0t` and uses the digits 0, 1 and T. In this 
    notation T stands for -1 and the radix remains 3 as expected. Like other
    numeric literals it supports both floating point and exponents. A longer
    explanation can be found [here](balanced_ternary.md).
  - Lastly, non-finite values of `∞`, `-∞`, `⦰` are allowed. These can be
    interpreted as equivalent to 1/0, -1/0, and 0/0. In IEEE 754 these would be
    translated as positive and negative infinity and NaN. In [trans-real
    arithmetic](https://doi.org/10.36285/tm.91) these would be interpreted as
    positive and negative infinity and nullity. In standard mathematics these
    would be invalid numeric literals.


- Strings: single, double and back quoted strings are all supported. These all 
  support string interpolation and uniformly combine with raw and multiline 
  variants.

- Symbols: these include the single character delimiters `(`, `)`, `[`, `]`, 
  `{`, `}`and also punctuation such as `,` and `;`. These
  do not 'stick' to any other character, although may appear inside strings. 
  Delimiters play the dual roles of nesting expressions and
  supporting function and method calls. 
  e.g. `(x + y) * z` vs `f[x, y]` and `x.m(p, q)`.

- Signs: these are runs of sign-characters such as `+`, `**`, `-->` and so
  on. These primarily play the role of infix operators e.g. `x := y`. In
  some places they act as supporting 'punctuation'.

- Identifiers: the usual rule for identifiers is followed - they start with
  an alphabetical character or an underscore and continue with those, plus digits.
  However Monogram also supports the use of escape sequences in identifiers.

## Numbers in detail

These are in the same format as Python's decimal numbers. However we intend to
extend them to include different radixes.

## Strings in detail

### String quotes

Monogram allows all three string quotes to be utilized, so that all three of
these tokens are valid:

- `"This is a string"`
- `'This is also a string'`
- ``` `And so is this` ```

Strings start and end with the quote-delimiters that is used to start the string.
Inside the string the quote-delimiters must be escaped using `\`.

### Escape Sequences

All strings support escape sequences. The escape character is `\` and JSON
escape sequences are supported.

In addition, Monogram supports the empty escape sequence `\_`. This escape
sequence turns into no additional characters! It is used to add visual breaks 
in strings of digits and similar use-cases.

## String Interpolation

All strings support string interpolation. These are embedded using `\` followed
by a bracket delimiter e.g `hello, \(name)` or `The value of x is \[x]`.
Interpolated strings are automatically expanded by monogram into a tree
structure, so in some sense they are not normal tokens. Our example 
`hello, \(name)` would expand as follows:

```xml
<concatenate>
  <literal type="string" value="hello, "/>
  <interpolated delimiter="parentheses">
    <identifier name="name"/>
  </interpolated>
</concatenate>
```

In short, it is expanded into a "concatenation" with literal strings and
interpolated expressions, each of which are tagged with the delimiter used.
Exactly how this is processed is not part of monogram but the downstream
processing.

### Raw strings

Escaping inside a string with `\` can be disabled by prefixing a string token 
with `\`. This is useful when typing regular expressions. For example: 

- `\"\n"` is a two character string consisting of a backslash followed by the
  letter `n`.

This of course works regardless of which quote delimiter is used.

### Multiline strings

Finally multiline strings are introduced with triple-quote delimiters. This 
has been made familiar in languages such as Python, Java, Kotlin, Scala and 
Julia. Monogram uses the same rule as Java: the opening and closing triple
quotes must be on separate lines from the rest of the string content and common
indentation is removed.

```
# Valid Monogram
print("""
This is a valid multiline string
in Monogram.
""")
```

```
# Invalid Monogram
print("""
But this will cause a syntax error
because the closing quotes are not
on a separate line from the content.""")
```


## Symbols in detail

Symbols are single-character tokens that do not glue to any other character.
The list is as follows:

- Delimiter-pairs
  - `(` and `)`
  - `[` and `]`
  - `{` and `}`

- Punctuation
  - `;`, semicolon
  - `,`, comma

Note that `:`, `<` and `>` are signs and not not symbols.

## Signs in detail

Signs are runs of non-alphabetic, printing characters that have been historically 
used as part of infix/prefix operators in mainstream programming languages, esp.
those descended from `C`. The list includes:

- `!`, `@`, `$`, `%`, `^`, `&`, `*`, `-`, `=`, `+`, `<`, `>`, `.`, `/`, `?`, `|`

These can occur as single character tokens bur readily glue together to
form multi-character tokens. For example, here are some tokens and what they
have been used for in programming languages.

- `..<` often used to define half-open ranges
- `+/-` an operator that returns two results
- `++` concatenation
- `:=` assignment
- `=>` lambda expression

Signs are used as both prefix and infix operators. In fact the same sign can
be used in both ways inside the same expression! e.g. `(x - y ) * -z`

The precedence of signs is determined by two factors:

1. The first character and its position in this list: `.({[*/%+-<~!&|?:=`. 
   - The order is designed to broadly correspond to that established in C
     and endless copied after that.
2. Whether or not the first character is repeated.

If the position in the list is N then the precedence is: 
`if isrepeated then: 10 * N + 9 else: 10 * N + 10 endif`

Lower precedences bind tighter and equal precedences are right-associative. 
This means that `a * b * c` will bind like `a * ( b * c )`.

## Identifiers in detail

Basic identifiers are familiar from a host of other programming languages.
They start with a letter or underscore and continue with letters, digits and
underscores. Basic identifier such as `def` or `let` can be used as reserved
words if there is a matching form-end (`enddef` or `endlet`) or they are used
in prefix form (`def!` or `let!`).

However, extended identifiers allow the use of escape sequences e.g. `Open\sSesame`.
These may not be used as reserved words. Hence `\_endure` is identifer that is
spelled `endure`, since `\_` is the empty escape sequence.

Note that escape sequences for identifiers do not include expression 
interpolation, so `\(` just stands for a 1-character identifier.

## Token rules as a railroad diagram

![Railroad diagram](images/token_rules.png)