# Monogram

🚨 **Alpha Status**: This project is currently in alpha. Features and functionality may change, and breaking changes may occur as development continues. Use at your own risk. 🚨

Monogram is a "no batteries" notation for writing domain-specific programs and
configuration files. It is easy for humans to read and write. It is easy for
machines to parse and generate. It deliberately borrows from many programming
languages but feels familiar to Python and Ruby programmers.

## _"It's source code, Jim. But not as we know it!"_

Here's an initial example to help explain what we mean by 'batteries not included'.
To experienced programmers, the following code looks a lot like the definition
of the factorial function:
```py
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

However, the twist is that Monogram has no idea what `def` or `if` might mean!
Nor does it have a clue about `*` or `-` either. And it definitely cannot
execute this program. 

And yet Monogram can easily translate this example into neatly structured XML
(shown below). Or it can translate to [JSON](docs/json.md) or [YAML](docs/yaml.md).
```xml
<form syntax="surround">
    <part keyword="def">
        <apply kind="parentheses" separator="undefined">
            <identifier name="f"/>
            <arguments>
                <identifier name="n"/>
            </arguments>
        </apply>
    </part>
    <part keyword="_">
        <form syntax="surround">
            <part keyword="if">
                <operator name="&lt;=" syntax="infix">
                    <identifier name="n"/>
                    <number value="1"/>
                </operator>
            </part>
            <part keyword="_">
                <number value="1"/>
            </part>
            <part keyword="else">
                <operator name="*" syntax="infix">
                    <identifier name="n"/>
                    <apply kind="parentheses" separator="undefined">
                        <identifier name="f"/>
                        <arguments>
                            <operator name="-" syntax="infix">
                                <identifier name="n"/>
                                <number value="1"/>
                            </operator>
                        </arguments>
                    </apply>
                </operator>
            </part>
        </form>
    </part>
</form>
```

Alternatively it can render the code as a diagram using Mermaid (below) or 
[Graphviz](docs/dot.md). Here's the same structure visualised as a graph.

```mermaid
graph TD
  137321037236880["form: surround"]:::custom_form;
  137321037236960["part: def"]:::custom_part;
  137321037236880 --> 137321037236960;
  137321037237040["apply"]:::custom_apply;
  137321037236960 --> 137321037237040;
  137321037237120["identifier: f"]:::custom_identifier;
  137321037237040 --> 137321037237120;
  137321037237200["arguments"]:::custom_arguments;
  137321037237040 --> 137321037237200;
  137321037237280["identifier: n"]:::custom_identifier;
  137321037237200 --> 137321037237280;
  137321037237360["part: _"]:::custom_part;
  137321037236880 --> 137321037237360;
  137321037237440["form: surround"]:::custom_form;
  137321037237360 --> 137321037237440;
  137321037237520["part: if"]:::custom_part;
  137321037237440 --> 137321037237520;
  137321037237600["operator: <="]:::custom_operator;
  137321037237520 --> 137321037237600;
  137321037237680["identifier: n"]:::custom_identifier;
  137321037237600 --> 137321037237680;
  137321037237760["number: 1"]:::custom_number;
  137321037237600 --> 137321037237760;
  137321037237840["part: _"]:::custom_part;
  137321037237440 --> 137321037237840;
  137321037237920["number: 1"]:::custom_number;
  137321037237840 --> 137321037237920;
  137321037238000["part: else"]:::custom_part;
  137321037237440 --> 137321037238000;
  137321037238080["operator: *"]:::custom_operator;
  137321037238000 --> 137321037238080;
  137321037238160["identifier: n"]:::custom_identifier;
  137321037238080 --> 137321037238160;
  137321037238240["apply"]:::custom_apply;
  137321037238080 --> 137321037238240;
  137321037238400["identifier: f"]:::custom_identifier;
  137321037238240 --> 137321037238400;
  137321037238560["arguments"]:::custom_arguments;
  137321037238240 --> 137321037238560;
  137321037238720["operator: -"]:::custom_operator;
  137321037238560 --> 137321037238720;
  137321037238880["identifier: n"]:::custom_identifier;
  137321037238720 --> 137321037238880;
  137321037239040["number: 1"]:::custom_number;
  137321037238720 --> 137321037239040;

classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;
```

In other words, Monogram is just a notation for writing program-like "code" but
comes without any built-in meanings. Although it is not infinitely flexible, it 
can often save you the effort of designing the syntax and implementing a parser
when you want an application/domain-specific language.

For more examples and more output formats (like JSON, YAML, PNG) see the 
[examples page](docs/examples.md).


## Monogram grammar

### Overview of tokens

The basic building blocks of a Monogram document are tokens - that is to say
numbers (`123`, `-0.12`), strings (`"hello, world"`), symbols (`{`, `}`), signs
(`:`, `++`) and various kinds of identifiers (`true`, `x`, `while`). These will
be largely familiar to anyone used to working with JSON or any mainstream
programming language.

Full details of tokenisation are given on [this page](docs/tokens.md) but
because these are generally so familiar to most programmers we highlight just a
few aspects that will be less familiar here:

- **Strings** support all three quote characters: single , double and back quotes.
    - All three are completely symmetrical in their design.
    - And support escape sequences, string interpolation, and raw and multiline
      versions.
    - `\_` is an escape sequence that expands into no characters. This helps
      with escaped identifiers and also inserting visual breaks into long
      strings e.g. phone numbers `"0765\_432\_1098"`

- **Symbols** include parentheses, brackets and braces as well as punctuation
  such as `,` and `;` (but not `.`)
    - The three different brackets are treated symmetrically
    - So these are all valid expressions, for instance: `m.f(x)`,  `m.f[x]`, `m.f{x}`.

- **Operators** are runs of sign-characters. In addition to familiar single-character
  operators such as `+`, `*`, `^`, Monogram allows for arbitrary combinations
  such as `:=`, `-->` or even `++^=!$$`. 
    - These primarily play the role of infix operators.
    - Operator precedence is decided on the first character of the sign and follows
      the precedence rules of the C-programming language. As a consequence,
      we can use sequences such as `s = x + y * z` and get expected results.
    - N.B. If the first character is repeated then the precedence is slightly 
      adjusted so it binds slightly more tightly. Which is why `p = a == b`
      binds the expected way. 

- **Identifiers** 
  - Support arbitrary identifiers via escape sequences e.g. `hello\,\sworld`
  - The empty-escape sequence is the neat way to handle reserved words e.g. `\_if`
  - Identifiers starting `end` are key to the way the grammar works as they
    mark reserved words.

### Overview of the grammar

In the next section we give the formal grammar in railroad diagram format. But
first we explain the main elements of it.

#### Operators

Firstly, Monogram's infix operators provide the basic **operator** precedence
syntax. This allows you to build up the familar alternating pattern of
expressions and operators. e.g. ```alpha + beta * gamma```. Any sequence
of 'sign' characters can be used as an infix operator. This will turn into
(say) XML that looks like:

```xml
<operator name="+" syntax="infix">
    <identifier name="alpha"/>
    <operator name="*" syntax="infix">
        <identifier name="beta"/>
        <identifier name="gamma"/>
    </operator>
</operator>
```

#### Brackets

Secondly, all three **brackets** `()`, `[]` and `{}` can be used to enclose a 
sequence of comma-or-semicolon separated expression. You can use either commas
or semicolons but not both. e.g. ```[alpha, beta, gamma]``` and ```(alpha; beta; gamma)```. These turn into the following XML respectively.

```xml
<delimited kind="brackets" separator="comma">
    <identifier name="alpha"/>
    <identifier name="beta"/>
    <identifier name="gamma"/>
</delimited>

<!-- and -->

<delimited kind="parentheses" separator="semicomma">
    <identifier name="alpha"/>
    <identifier name="beta"/>
    <identifier name="gamma"/>
</delimited>
```

### Function/Method calls

All three brackets also support **function and method call** syntax. These look
like ```table[key]``` and ```table.lookup(key)```. These respectively turn into 
these:

```xml
<apply kind="brackets" separator="undefined">
    <identifier name="table"/>
    <arguments>
        <identifier name="key"/>
    </arguments>
</apply>

<!-- and -->

<invoke kind="parentheses" name="lookup" separator="undefined">
    <identifier name="table"/>
    <arguments>
        <identifier name="key"/>
    </arguments>
</invoke>
```

#### Property/Field accesses

And since we have touched on method-like syntax, this is a good place to mention
**property-like** syntax e.g. ```table.length```. That turns into:

```xml
<get name="length">
    <identifier name="table"/>
</get>
```


#### Forms

Next we have **forms**, which are characterised by an enclosing
pair of distinctive identifiers, where the closing identifer is the same
as the opener but prefixed by "end". e.g. ```if ... endif``` or 
```whoop ... endwhoop```. Almost any identifier will do for the opening 
keyboard (although it may not start or end with an underscore).

The idea behind forms is that they allow us to mimick the multi-line syntax of
(say) `if`, `while` and `foreach` constructs from languages such as Javascript
or C#. As a consequence the expressions enclosed within a form are separated by
semi-colons and not commas. 

Forms typically have multiple interior sections, called "parts" which are
separated by "breakers". The basic type of breaker is an identifier followed by
a colon (`:`). The syntax is chosen to echo the look-and-feel of Python whilst
avoiding the need for any reserved words. e.g.

```
while test() do: 
    x += 1 
endwhile
```

The above example has two parts. The first part lies between `while` and `do:` 
and the second part is sandwiched between `do:` and `endwhile`. This example
would turn into:

```xml
<form syntax="surround">
    <part keyword="while">
        <apply kind="parentheses" separator="undefined">
            <identifier name="test"/>
            <arguments/>
        </apply>
    </part>
    <part keyword="do">
        <operator name="+=" syntax="infix">
            <identifier name="x"/>
            <number value="1"/>
        </operator>
    </part>
</form>
```

Note how the first part of the form takes the opening identifier as its
"keyword". The second part of the form takes the name-part of the breaker.

Todays programming languages have tended to veer away from using intermediate
keywords such as `then` or `do`. To help make Monogram feel more familiar, 
we have followed Python in allowing the breaker-name to be omitted 
immediately after the opening keyword. So we could have written this example
like this, very similar to Python's syntax if you can overlook the `endwhile` 
:smile:.

```
while test(): 
    x += 1 
endwhile
```

Furthermore many programming languages 'cascade' of conditions via an
intermediate keyword such as `elif`. Here it is in Python:

```py
# In Python syntax
if test():
    statements
elif other_test():      # Cascaded if
    other_statements
else:
    catch_all_statements
```

Monogram allows us to get quite close to this pattern of named and anonymous
sections by utilizing _compound breakers_. Compound breakers are a hypenated
pair of identifiers e.g. `else-if` or `and-while`, where the second identifier
reuses the enclosing form-start. And immediately after a compound breaker we are
allowed another anonymous breaker.

Here's the equivalent of the above Python snippet in Monogram.

```
if test():              # Anonymous breaker
    statements
else-if other_test():   # A second anonymous breaker
    other_statements
else:
    catch_all_statements
endif
```

Breakers give their names to the parts they introduce. But anonymous breakers do
not have a name. To handle this, any parts introduced by an anonymous breaker are
treated as if they were named `_` (this can be overridden).

Hence the above example would turn into this XML:

```xml
<form syntax="surround">
    <part keyword="if">
        <apply kind="parentheses" separator="undefined">
            <identifier name="test"/>
            <arguments/>
        </apply>
    </part>
    <part keyword="_">
        <identifier name="statements"/>
    </part>
    <part keyword="else-if">
        <apply kind="parentheses" separator="undefined">
            <identifier name="other_test"/>
            <arguments/>
        </apply>
    </part>
    <part keyword="_">
        <identifier name="other_statements"/>
    </part>
    <part keyword="else">
        <identifier name="catch_all_statements"/>
    </part>
</form>
```

#### Prefix forms

Last but not least we have **prefix forms**. Most programming
languages utilize simple prefix forms such as `return` or `pass`. Monogram
imitates these like this:

```
if t:
    return! 99
else:
    pass!
endif
```

By placing an `!` after an ordinary identifier, it is treated as a form that
takes a single optional expression. And this example turns into the following
XML, where the prefix form is treated as a form with an optional part:

```xml
<form syntax="surround">
    <part keyword="if">
        <identifier name="t"/>
    </part>
    <part keyword="_">
        <form syntax="prefix">
            <part keyword="return">
                <number value="99"/>
            </part>
        </form>
    </part>
    <part keyword="else">
        <form syntax="prefix">
            <part keyword="pass"/>
        </form>
    </part>
</form>
 ```


### Railroad diagrams

Here's the grammar for Monogram as a railroad diagram; also available in
[HTML](docs/grammar.html), [PDF](docs/images/grammar.pdf) and
[PNG](docs/images/grammar.png).

![Monogram Grammar PDF](docs/images/grammar.png) 
