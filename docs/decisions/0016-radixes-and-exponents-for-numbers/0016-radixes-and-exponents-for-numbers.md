# 0016 - Radixes and exponents for numbers, 2025-04-24

## Issue

Today's programming notations typically support integers in hex, octal and
binary as well as decimal e.g. Python, Java, C/C++, Go, Rust, Ruby, Swift, PHP. 

However both Common Lisp and Pop-11 support integers in any base from 2 to 36.
Striking similar notation is used in both. For example:

- `#36rZ` is 35 in base 36 in Common Lisp
- `36rZ` is 35 in base 36 in Pop-11

And uniquely Pop-11 went further by allowing fractional parts in non-decimal
bases. Here's 12.5 in binary in Pop11: 

- `2r1100.1`

And it even supported exponents, although similar to the way the base is always
written in decimal, the exponent is too.

- `2r1e20`, which is 2 to the power of 20 or 1048576 (aka a binary 'mega')

Having gone this far, it is unsurprising that Pop-11 allowed these to be
combined in literals such as:

- `16r1.1e2`, which is 272.0 in decimal notation.

So the question is which of these options should Monogram support? Only the
fixed bases of 2, 8, 10 and 16 for integers, any base from 2-36 for integers,
any base with a fractional point, or any base with fractional point and
exponent notation?

## Factors

From a practical viewpoint there are solid cases for integer values in
binary and hex. Octal is a lot less significant these days but is established
in C and Unix shells. Good use cases for other bases are much less common e.g. 
base 36 for URL shortening or, in the case of 3GPP, for versioning documents.

But there are only rare niche use-cases for non-decimal floating point. And to
make use of Monogram, devs will need to translate number-text into the native
number format of their chosen language - typically 64-bits ints or 64-bit
"double precision" floating point. But normal programming languages do not come
equipped with the ability to translate non-decimal floating points.

If we translate Monogram into XML, the processing will not be done via a
Monogram API but some other tool that knows nothing about these exotic formats. 

From a purist point of view, the concepts of decimal numbers with decimal point
and exponent work every bit as well in binary, hex or duodecimal. Limiting these
advantages to decimal is just baking in convention. And although the niche
applications are rare, it is very frustrating to recreate the mechanisms that
are already implemented but in a different base.

## Options

- Option 1: Only support binary, octal and hex for integers. e.g. 0xFF
- Option 2: Support bases 2-36 for integers but provide a --decimal
  option that adds an additional `decimal` attribute to numbers. e.g. 36rZZ
- Option 3: Support bases 2-35 for integers and floating points with a
  --with-decimal option that adds an additional `decimal` attribute to numbers.
  e.g. 0xFF.8

### The --decimal option

This works by adding an additional attribute `decimal` to number elements
which has the text of the same number but in decimal notation. For example:
`36rZZ` would become the following in XML:

```xml
<number decimal="1295" value="36rZZ" />
```

The idea is that programs that consume XML or JSON do not need to write special
decoding routines but can ask the monogram tool to do the conversion for them.

## Pros and Cons of Options

### Option 1: Binary, octal and hex integers

- Pros
    - Equivalent or better than numeric literal support in all mainstream
      languages.
    - Conversion to native formats needs no extra support
        -  With the caveat that the numbers are in range - a range that 
           constantly changes with advancing  
    - No teaching overhead.

- Cons
    - Does not cover Common Lisp, Racket, Ada, Prolog, Smalltalk or Pop-11.
    - Limitation lacks any strong reason apart from convention.

- Interesting
    - There's no teaching overhead for programmers because it is following
      convention but it is a bit baffling to mathematically able people.
    - Most mainstream languages only offer 64-bit integers, so the argument
      that no special features are using is only true for short numeric
      literals! (Python is an exception since it supports arbitrary sized
      integers.)

### Option 2: Integers with Bases 2-36, with --decimal option

- Pros
    - Coverage equivalent to virtually all current programming languages.
    - The `--decimal` option eliminates the need for non-standard
      library routines.

- Cons
    - The `--decimal` option makes the AST bulkier
    - Still weaker than Pop-11, which represents a gold standard in this area.
    - The failure to cover fractional point is still merely convention.

### Option 3: Mantissa of all numbers with Bases 2-36, with --decimal option

- Pros
    - Coverage of finite real numbers is as good as any programming language
      I am aware of.
    - The `--decimal` option eliminates the need for non-standard
      library routines.

- Cons
    - The `--decimal` option makes the AST bulkier
    - The radix (the base part) and exponential parts are always in decimal
      notation and this has to be taught as almost no programmers have prior
      experience of base N != 10 with floating point and/or exponentials.

- Interesting
    - The `decimal` attribute needs to preserve the integer vs floating point
      quality of the numeric literal - and exponentials automatically count
      as floating pont (to keep the rule simple).
    - This leaves open the question of rationals and complex numeric literals.


## Outcome and Consequences

**Option 3**: The mantissa of numbers can be written in any base from 2-36.

As far as I can tell there is simply no argument that the `--decimal` option
does not fully counter in principle, although it does bulk up the AST. If we
provide a `ConvertToDecimal` function as part of the API even that overhead is
removed.

## Additional Notes

In previous drafts I used mathematical terms that are not in common use. I 
have replaced them with slightly less precise but more accessible language:

- The technically accurate term "radix" was replaced by "base" throughout.
- The term "denary" is replaced by "decimal". I replaced it on the basis that
  "decimal" is in much greater usage and can just mean "base 10" and does not
  imply a decimal point.

### The ConvertToDecimal function

In addition a `mg.ConvertToDecimal: string -> string` function will be provided
in the Go reference implementation, so that programmers using the library can
perform the conversion on demand. This is useful for using the same notation in
other contexts, for example.
