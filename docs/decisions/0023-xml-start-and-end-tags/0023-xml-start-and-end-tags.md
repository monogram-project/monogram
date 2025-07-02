# 0023 - XML Start and End Tags, 2025-06-30

## Issue

The concept is to support XML templating in Monogram with a focus on
data-centric XML documents. These occur in systems such as `SOAP`, `MathML`,
`Chemical Markup`, `Geography Markup` and commonly in configuration files such
as `ant`, `maven` and `Visual Studio`. This is the kind of application that we
hope Monogram will support nicely.

The following example, which consists of start, end and standalone tags is to
be valid Monogram syntax.
```xml
<person>
  <name first="John" last="Doe"/>
  <age value="30"/>
</person>
```

## Factors

### Special Syntax

To make this work we need to reserve some special syntax, unfortunately:

- `<` and `</` in prefix position have their meanings pinned to XML-tags.
- signs `><` `/><` `></` and `/></` need to be split between `>` and `<`

### Weak Support for Document-Centric XML

Where document-centric XML was needed, we would support character data using
Monogram's rich string syntax. This is more fussy than XML for document-centric
text and doesn't provide good support for `span`-style inserts. 

For example in normal XML you could write:

```xml
<letter>
  <to>Jane</to>
  <from>John</from>
  <body>Hello Jane, I hope you're well.</body>
</letter>
```

But in Monogram we would need to use string-quotes to represent this as:

```xml
<letter>
  <to>"Jane"</to>
  <from>"John"</from>
  <body>"Hello Jane, I hope you're well."</body>
</letter>
```

### Interpolating values into tags

We need to be able to embed arbitrary Monogram expressions inside tags. For
example:

```xml
<person>
  <name first="John" last="Doe"/>
  <age value=(age_of("John", "Doe"))/>
</person>
```

We need to decide which parts of a tag can be replaced with a general
expression, what the syntax looks like and what the generated tree is like.
Also, if we allow the replacement of tag-names then there is a question
over how we match it with an end-tag.

```txt
<(EXPR)>
    ...
</ ??? >
```


## Options

I have grouped the complex set of options around topics, labelled each
topic A-F.

- Option A1: `>` does not stick to `<` in signs e.g. `++><++` becomes two
  tokens `++>` and `<++`.

- Option A2: The tokens `><` `/><` `></` and `/></` are specially split.

- Option A3: Require programmers to insert whitespace. e.g. `<foo><bar/></foo>`
  has to be written `<foo> <bar/> </foo>`

- Option B1: Provide no additional support for document-centric XML.

- Option B2: Arrange that text inside a start/end tag is assumed to be character
  data until some embedding syntax is encountered e.g. `<< ... >>`.

- Option C1: Embedded expressions must be placed within brackets `()`, `[]` or
  `{}`. 

- Option C2: Embedded expressions are parsed at a precedence level tighter than
  `>` or `/`.

- Option D1: The tag name can be substituted with an embedded expression.
  e.g. `< (choice) />`

- Option D2: The attribute name can be substituted with an embedded expression.
  e.g. `<process (option)="true">`

- Option D3: The attribute value can be substituted with an embedded expression.
  e.g. `<foo data=(fetch_from_db())>`

- Option D4: Attribute name=value expressions can be substituted with an 
  embedded expression. `<foo (attribute_pair) />`

- Option E1: Attribute values can be any Monogram literal.

- Option E2: Attribute values must be a Monogram string and may use monogram
  escaping and string-tags.

- Option E3: Attribute values must be a single/doubled quoted string to match
  XML and use XML escape syntax.

- Option F1: If an attribute name is embeddable then it can be closed using
  a wild card e.g. `< (choice) > details </_>`

- Option F2: It an attribute name is embeddable then it can be closed using
  an identical value. e.g. `< (choice) > details </ (choice) >`


## Pros and Cons of Options

### Option A

Outcome: Option A1 is preferred because of the simplicity of the syntactic rule.

#### Option A1

Option A1: `>` does not stick to `<` in signs e.g. `++><++` becomes two tokens `++>` and `<++`.

- Pros
  - The syntax is simple and easy to understand.
  - It is simple to implement in the tokeniser.
- Cons
  - It has a bigger syntactic footprint than Option A2,

#### Option A2

Option A2: The tokens `><` `/><` `></` and `/></` are specially split.

- Pros
  - Syntactic footprint is as small as possible.
- Cons
  - It has 4 special cases for people to remember.

### Option B

Outcome: Option B1 is preferred because B2 has a showstopper issue with
whitespace.

#### Option B1

Provide no additional support for document-centric XML.

- Pros
  - Nothing to remember, no special syntax.
- Cons
  -  It is not as user-friendly for document-centric XML.

#### Option B2

Arrange that text inside a start/end tag is assumed to be character
data until some embedding syntax is encountered e.g. `<< ... >>`.

- Pros
  - It is more user-friendly for document-centric XML.
  - It allows for more complex text inside tags.
- Cons
  - It repeats the same dire mistake of XML, which is that it it not
    indentation-friendly because whitespace is significant.

### Option C

Outcome: C1 because C2's grammatical ambiguty is worse than it looks and 
it is difficult to learn the precedence levels.

#### Option C1

Embedded expressions must be placed within brackets `()`, `[]` or `{}`. 

- Pros
  - It is 100% unambiguous.
  - It is easy to teach and remember.
- Cons
  - It is a bit more verbose than Option C2.

#### Option C2

Embedded expressions are parsed at a precedence level tighter than  `>` or `/`.

- Pros
  - It is terser than Option C1.
- Cons
  - It is not 100% unambiguous, as it can be confused with a
    comparison expression.


### Option D

Outcome: D1, D2, D3, D4. I am not aware of any reason for limiting any
of these.

- Option D1: The tag name can be substituted with an embedded expression.
  e.g. `< (choice) />`
- Option D2: The attribute name can be substituted with an embedded expression.
  e.g. `<process (option)="true">`
- Option D3: The attribute value can be substituted with an embedded expression.
  e.g. `<foo data=(fetch_from_db())>`
- Option D4: Attribute name=value expressions can be substituted with an 
  embedded expression. `<foo (attribute_pair) />`


### Option E

Outcome: E1 because limiting attribute values to strings has no compelling
reason. If restrictions need to be applied, this is not Monogram's job.

#### Option E1

Attribute values can be any Monogram literal.

- Pros
  - Simple to teach
  - Consistent with the language.
  - Supports numerical and date syntax.
- Cons
  - Not valid attribute values if the target is actually XML.


#### Option E2

Attribute values must be a Monogram string and may use monogram escaping and
string-tags.

- Pros
  - Simple to teach
  - Consistent with the language.
  - Directly corresponds with XML.
- Cons
  - Requires string based encoding of number and date types.


#### Option E3

Attribute values must be a single/doubled quoted string to match XML and use XML
escape syntax.

- Pros
  - Directly corresponds with XML.
- Cons
  - Requires string based encoding of number and date types.
  - These strings are never going to be consistent with XML syntax and 
    that is a massive in teaching load.


### Option F

Outcome: F1, since F2 is not a real alternative.

#### Option F1

If an attribute name is embeddable then it can be closed using a wild card e.g.
`< (choice) > details </_>`

- Pros
  - Trivial to teach
  - Trivial to implement
  - No overhead of deferred check
- Cons
  - None

#### Option F2

It an attribute name is embeddable then it can be closed using an identical
value. e.g. `< (choice) > details </ (choice) >`

- Pros
  - None
- Cons
  - Gives the impression of making a useful check - but Monogram has no
    evaluation semantics

    
## Outcome and Consequences

In summary:

- Option A1: `>` does not stick to `<` in signs e.g. `++><++` becomes two
  tokens `++>` and `<++`.
- Option B1: Provide no additional support for document-centric XML.
- Option C1: Embedded expressions must be placed within brackets `()`, `[]` or
  `{}`. 
- Option D1: The tag name can be substituted with an embedded expression.
  e.g. `< (choice) />`
- Option D2: The attribute name can be substituted with an embedded expression.
  e.g. `<process (option)="true">`
- Option D3: The attribute value can be substituted with an embedded expression.
  e.g. `<foo data=(fetch_from_db())>`
- Option D4: Attribute name=value expressions can be substituted with an 
  embedded expression. `<foo (attribute_pair) />`
- Option E1: Attribute values can be any Monogram literal.
- Option F1: If an attribute name is embeddable then it can be closed using
  a wild card e.g. `< (choice) > details </_>`
