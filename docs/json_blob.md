# JSON Expression

## Monogram

Monograph supports JSON expressions. Since Monogram is a good deal more 
complicated than JSON, it is not too surprising that translating a JSON
expression into a Monogram-tree looks quite different from plain JSON!

```txt
{
  "person": {
    "name": "Alice",
    "age": 25,
    "isStudent": true,
    "skills": ["Python", "JavaScript", "SQL"],
    "address": {
      "street": "123 Maple Street",
      "city": "Exampleville",
      "country": "Neverland"
    },
    "favoriteBooks": [
      {
        "title": "To Kill a Mockingbird",
        "author": "Harper Lee",
        "yearPublished": 1960
      },
      {
        "title": "1984",
        "author": "George Orwell",
        "yearPublished": 1949
      }
    ]
  }
}
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  129990958089536["delimited"]:::custom_delimited;
  129990958089616["operator: :"]:::custom_operator;
  129990958089536 --> 129990958089616;
  129990958089696["string: person"]:::custom_string;
  129990958089616 --> 129990958089696;
  129990958089776["delimited"]:::custom_delimited;
  129990958089616 --> 129990958089776;
  129990958089856["operator: :"]:::custom_operator;
  129990958089776 --> 129990958089856;
  129990958089936["string: name"]:::custom_string;
  129990958089856 --> 129990958089936;
  129990958090016["string: Alice"]:::custom_string;
  129990958089856 --> 129990958090016;
  129990958090096["operator: :"]:::custom_operator;
  129990958089776 --> 129990958090096;
  129990958090176["string: age"]:::custom_string;
  129990958090096 --> 129990958090176;
  129990958090256["number: 25"]:::custom_number;
  129990958090096 --> 129990958090256;
  129990958090336["operator: :"]:::custom_operator;
  129990958089776 --> 129990958090336;
  129990958090416["string: isStudent"]:::custom_string;
  129990958090336 --> 129990958090416;
  129990958090496["identifier: true"]:::custom_identifier;
  129990958090336 --> 129990958090496;
  129990958090576["operator: :"]:::custom_operator;
  129990958089776 --> 129990958090576;
  129990958090656["string: skills"]:::custom_string;
  129990958090576 --> 129990958090656;
  129990958090736["delimited"]:::custom_delimited;
  129990958090576 --> 129990958090736;
  129990958090816["string: Python"]:::custom_string;
  129990958090736 --> 129990958090816;
  129990958090896["string: JavaScript"]:::custom_string;
  129990958090736 --> 129990958090896;
  129990958090976["string: SQL"]:::custom_string;
  129990958090736 --> 129990958090976;
  129990958091056["operator: :"]:::custom_operator;
  129990958089776 --> 129990958091056;
  129990958091136["string: address"]:::custom_string;
  129990958091056 --> 129990958091136;
  129990958091216["delimited"]:::custom_delimited;
  129990958091056 --> 129990958091216;
  129990958091296["operator: :"]:::custom_operator;
  129990958091216 --> 129990958091296;
  129990958091376["string: street"]:::custom_string;
  129990958091296 --> 129990958091376;
  129990958091456["string:<br/>123 Maple Street"]:::custom_string;
  129990958091296 --> 129990958091456;
  129990958091536["operator: :"]:::custom_operator;
  129990958091216 --> 129990958091536;
  129990958091616["string: city"]:::custom_string;
  129990958091536 --> 129990958091616;
  129990958091696["string: Exampleville"]:::custom_string;
  129990958091536 --> 129990958091696;
  129990958091776["operator: :"]:::custom_operator;
  129990958091216 --> 129990958091776;
  129990958091856["string: country"]:::custom_string;
  129990958091776 --> 129990958091856;
  129990958091936["string: Neverland"]:::custom_string;
  129990958091776 --> 129990958091936;
  129990958092016["operator: :"]:::custom_operator;
  129990958089776 --> 129990958092016;
  129990958092096["string:<br/>favoriteBooks"]:::custom_string;
  129990958092016 --> 129990958092096;
  129990958092176["delimited"]:::custom_delimited;
  129990958092016 --> 129990958092176;
  129990958092256["delimited"]:::custom_delimited;
  129990958092176 --> 129990958092256;
  129990958092336["operator: :"]:::custom_operator;
  129990958092256 --> 129990958092336;
  129990958092416["string: title"]:::custom_string;
  129990958092336 --> 129990958092416;
  129990958092496["string:<br/>To Kill a Mockingbird"]:::custom_string;
  129990958092336 --> 129990958092496;
  129990958092576["operator: :"]:::custom_operator;
  129990958092256 --> 129990958092576;
  129990958092656["string: author"]:::custom_string;
  129990958092576 --> 129990958092656;
  129990958092736["string: Harper Lee"]:::custom_string;
  129990958092576 --> 129990958092736;
  129990958092816["operator: :"]:::custom_operator;
  129990958092256 --> 129990958092816;
  129990958092896["string:<br/>yearPublished"]:::custom_string;
  129990958092816 --> 129990958092896;
  129990958092976["number: 1960"]:::custom_number;
  129990958092816 --> 129990958092976;
  129990958093056["delimited"]:::custom_delimited;
  129990958092176 --> 129990958093056;
  129990958093136["operator: :"]:::custom_operator;
  129990958093056 --> 129990958093136;
  129990958093216["string: title"]:::custom_string;
  129990958093136 --> 129990958093216;
  129990958093296["string: 1984"]:::custom_string;
  129990958093136 --> 129990958093296;
  129990958093376["operator: :"]:::custom_operator;
  129990958093056 --> 129990958093376;
  129990958093456["string: author"]:::custom_string;
  129990958093376 --> 129990958093456;
  129990958093536["string:<br/>George Orwell"]:::custom_string;
  129990958093376 --> 129990958093536;
  129990958093616["operator: :"]:::custom_operator;
  129990958093056 --> 129990958093616;
  129990958093696["string:<br/>yearPublished"]:::custom_string;
  129990958093616 --> 129990958093696;
  129990958093776["number: 1949"]:::custom_number;
  129990958093616 --> 129990958093776;

classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;

```

## XML

```xml
<delimited kind="braces" separator="undefined">
  <operator syntax="infix" name=":">
    <string value="person" quote="double" />
    <delimited separator="comma" kind="braces">
      <operator syntax="infix" name=":">
        <string quote="double" value="name" />
        <string quote="double" value="Alice" />
      </operator>
      <operator syntax="infix" name=":">
        <string value="age" quote="double" />
        <number value="25" />
      </operator>
      <operator syntax="infix" name=":">
        <string quote="double" value="isStudent" />
        <identifier name="true" />
      </operator>
      <operator syntax="infix" name=":">
        <string quote="double" value="skills" />
        <delimited kind="brackets" separator="comma">
          <string quote="double" value="Python" />
          <string value="JavaScript" quote="double" />
          <string quote="double" value="SQL" />
        </delimited>
      </operator>
      <operator syntax="infix" name=":">
        <string quote="double" value="address" />
        <delimited kind="braces" separator="comma">
          <operator syntax="infix" name=":">
            <string quote="double" value="street" />
            <string quote="double" value="123 Maple Street" />
          </operator>
          <operator syntax="infix" name=":">
            <string quote="double" value="city" />
            <string quote="double" value="Exampleville" />
          </operator>
          <operator syntax="infix" name=":">
            <string quote="double" value="country" />
            <string quote="double" value="Neverland" />
          </operator>
        </delimited>
      </operator>
      <operator syntax="infix" name=":">
        <string quote="double" value="favoriteBooks" />
        <delimited kind="brackets" separator="comma">
          <delimited kind="braces" separator="comma">
            <operator syntax="infix" name=":">
              <string quote="double" value="title" />
              <string quote="double" value="To Kill a Mockingbird" />
            </operator>
            <operator syntax="infix" name=":">
              <string quote="double" value="author" />
              <string quote="double" value="Harper Lee" />
            </operator>
            <operator syntax="infix" name=":">
              <string quote="double" value="yearPublished" />
              <number value="1960" />
            </operator>
          </delimited>
          <delimited kind="braces" separator="comma">
            <operator syntax="infix" name=":">
              <string quote="double" value="title" />
              <string quote="double" value="1984" />
            </operator>
            <operator name=":" syntax="infix">
              <string quote="double" value="author" />
              <string quote="double" value="George Orwell" />
            </operator>
            <operator syntax="infix" name=":">
              <string quote="double" value="yearPublished" />
              <number value="1949" />
            </operator>
          </delimited>
        </delimited>
      </operator>
    </delimited>
  </operator>
</delimited>
```

## JSON

We can target JSON as an output format. The format of each node is
a bit verbose but straightforward:

```json
{
    "role": "{{NODE NAME}}",
    "ATTRIBUTE_1": "VALUE_1", 
    ... 
    "ATTRIBUTE_N": "VALUE_N",
    "children": [ 
        ...
    ]
}
```

And this is what it expands into:

```json
{
  "role": "unit",
  "src": "json_blob.mg",
  "children": [
    {
      "role": "delimited",
      "kind": "braces",
      "separator": "undefined",
      "children": [
        {
          "role": "operator",
          "syntax": "infix",
          "name": ":",
          "children": [
            {
              "role": "string",
              "quote": "double",
              "value": "person"
            },
            {
              "role": "delimited",
              "kind": "braces",
              "separator": "comma",
              "children": [
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": ":",
                  "children": [
                    {
                      "role": "string",
                      "quote": "double",
                      "value": "name"
                    },
                    {
                      "role": "string",
                      "value": "Alice",
                      "quote": "double"
                    }
                  ]
                },
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": ":",
                  "children": [
                    {
                      "role": "string",
                      "quote": "double",
                      "value": "age"
                    },
                    {
                      "role": "number",
                      "value": "25"
                    }
                  ]
                },
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": ":",
                  "children": [
                    {
                      "role": "string",
                      "quote": "double",
                      "value": "isStudent"
                    },
                    {
                      "role": "identifier",
                      "name": "true"
                    }
                  ]
                },
                {
                  "role": "operator",
                  "name": ":",
                  "syntax": "infix",
                  "children": [
                    {
                      "role": "string",
                      "value": "skills",
                      "quote": "double"
                    },
                    {
                      "role": "delimited",
                      "kind": "brackets",
                      "separator": "comma",
                      "children": [
                        {
                          "role": "string",
                          "quote": "double",
                          "value": "Python"
                        },
                        {
                          "role": "string",
                          "quote": "double",
                          "value": "JavaScript"
                        },
                        {
                          "role": "string",
                          "quote": "double",
                          "value": "SQL"
                        }
                      ]
                    }
                  ]
                },
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": ":",
                  "children": [
                    {
                      "role": "string",
                      "quote": "double",
                      "value": "address"
                    },
                    {
                      "role": "delimited",
                      "kind": "braces",
                      "separator": "comma",
                      "children": [
                        {
                          "role": "operator",
                          "syntax": "infix",
                          "name": ":",
                          "children": [
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "street"
                            },
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "123 Maple Street"
                            }
                          ]
                        },
                        {
                          "role": "operator",
                          "syntax": "infix",
                          "name": ":",
                          "children": [
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "city"
                            },
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "Exampleville"
                            }
                          ]
                        },
                        {
                          "role": "operator",
                          "name": ":",
                          "syntax": "infix",
                          "children": [
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "country"
                            },
                            {
                              "role": "string",
                              "quote": "double",
                              "value": "Neverland"
                            }
                          ]
                        }
                      ]
                    }
                  ]
                },
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": ":",
                  "children": [
                    {
                      "role": "string",
                      "quote": "double",
                      "value": "favoriteBooks"
                    },
                    {
                      "role": "delimited",
                      "kind": "brackets",
                      "separator": "comma",
                      "children": [
                        {
                          "role": "delimited",
                          "kind": "braces",
                          "separator": "comma",
                          "children": [
                            {
                              "role": "operator",
                              "syntax": "infix",
                              "name": ":",
                              "children": [
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "title"
                                },
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "To Kill a Mockingbird"
                                }
                              ]
                            },
                            {
                              "role": "operator",
                              "syntax": "infix",
                              "name": ":",
                              "children": [
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "author"
                                },
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "Harper Lee"
                                }
                              ]
                            },
                            {
                              "role": "operator",
                              "syntax": "infix",
                              "name": ":",
                              "children": [
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "yearPublished"
                                },
                                {
                                  "role": "number",
                                  "value": "1960"
                                }
                              ]
                            }
                          ]
                        },
                        {
                          "role": "delimited",
                          "kind": "braces",
                          "separator": "comma",
                          "children": [
                            {
                              "role": "operator",
                              "syntax": "infix",
                              "name": ":",
                              "children": [
                                {
                                  "role": "string",
                                  "value": "title",
                                  "quote": "double"
                                },
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "1984"
                                }
                              ]
                            },
                            {
                              "role": "operator",
                              "name": ":",
                              "syntax": "infix",
                              "children": [
                                {
                                  "role": "string",
                                  "value": "author",
                                  "quote": "double"
                                },
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "George Orwell"
                                }
                              ]
                            },
                            {
                              "role": "operator",
                              "syntax": "infix",
                              "name": ":",
                              "children": [
                                {
                                  "role": "string",
                                  "quote": "double",
                                  "value": "yearPublished"
                                },
                                {
                                  "role": "number",
                                  "value": "1949"
                                }
                              ]
                            }
                          ]
                        }
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}```

## YAML

We can target YAML as an output format. The format of each node is
a bit verbose but easy to understand.

```yaml
role": "{{NODE NAME}}",
ATTRIBUTE_1: VALUE_1,
...
children": 
- ...
- ...
```

And this is what it expands into:



```yaml
role: delimited
kind: braces
separator: undefined
children:
- role: operator
  syntax: infix
  name: ':'
  children:
  - role: string
    value: person
    quote: double
  - role: delimited
    separator: comma
    kind: braces
    children:
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        quote: double
        value: name
      - role: string
        quote: double
        value: Alice
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        value: age
        quote: double
      - role: number
        value: 25
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        quote: double
        value: isStudent
      - role: identifier
        name: 'true'
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        quote: double
        value: skills
      - role: delimited
        kind: brackets
        separator: comma
        children:
        - role: string
          quote: double
          value: Python
        - role: string
          value: JavaScript
          quote: double
        - role: string
          quote: double
          value: SQL
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        quote: double
        value: address
      - role: delimited
        kind: braces
        separator: comma
        children:
        - role: operator
          syntax: infix
          name: ':'
          children:
          - role: string
            quote: double
            value: street
          - role: string
            quote: double
            value: 123 Maple Street
        - role: operator
          syntax: infix
          name: ':'
          children:
          - role: string
            quote: double
            value: city
          - role: string
            quote: double
            value: Exampleville
        - role: operator
          syntax: infix
          name: ':'
          children:
          - role: string
            quote: double
            value: country
          - role: string
            quote: double
            value: Neverland
    - role: operator
      syntax: infix
      name: ':'
      children:
      - role: string
        quote: double
        value: favoriteBooks
      - role: delimited
        kind: brackets
        separator: comma
        children:
        - role: delimited
          kind: braces
          separator: comma
          children:
          - role: operator
            syntax: infix
            name: ':'
            children:
            - role: string
              quote: double
              value: title
            - role: string
              quote: double
              value: To Kill a Mockingbird
          - role: operator
            syntax: infix
            name: ':'
            children:
            - role: string
              quote: double
              value: author
            - role: string
              quote: double
              value: Harper Lee
          - role: operator
            syntax: infix
            name: ':'
            children:
            - role: string
              quote: double
              value: yearPublished
            - role: number
              value: 1960
        - role: delimited
          kind: braces
          separator: comma
          children:
          - role: operator
            syntax: infix
            name: ':'
            children:
            - role: string
              quote: double
              value: title
            - role: string
              quote: double
              value: '1984'
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: double
              value: author
            - role: string
              quote: double
              value: George Orwell
          - role: operator
            syntax: infix
            name: ':'
            children:
            - role: string
              quote: double
              value: yearPublished
            - role: number
              value: 1949

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc00007f650" [label="unit: json_blob.mg", shape="box", fillcolor="lightgray"];
  "node_0xc00007f5f0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007f650" -> "node_0xc00007f5f0";
  "node_0xc00007f590" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f5f0" -> "node_0xc00007f590";
  "node_0xc00007e210" [label="string: person", shape="box", fillcolor="lightgray"];
  "node_0xc00007f590" -> "node_0xc00007e210";
  "node_0xc00007f530" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007f590" -> "node_0xc00007f530";
  "node_0xc00007e330" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007e330";
  "node_0xc00007e270" [label="string: name", shape="box", fillcolor="lightgray"];
  "node_0xc00007e330" -> "node_0xc00007e270";
  "node_0xc00007e2d0" [label="string: Alice", shape="box", fillcolor="lightgray"];
  "node_0xc00007e330" -> "node_0xc00007e2d0";
  "node_0xc00007e450" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007e450";
  "node_0xc00007e390" [label="string: age", shape="box", fillcolor="lightgray"];
  "node_0xc00007e450" -> "node_0xc00007e390";
  "node_0xc00007e3f0" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e450" -> "node_0xc00007e3f0";
  "node_0xc00007e570" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007e570";
  "node_0xc00007e4b0" [label="string: isStudent", shape="box", fillcolor="lightgray"];
  "node_0xc00007e570" -> "node_0xc00007e4b0";
  "node_0xc00007e510" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e570" -> "node_0xc00007e510";
  "node_0xc00007e7b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007e7b0";
  "node_0xc00007e5d0" [label="string: skills", shape="box", fillcolor="lightgray"];
  "node_0xc00007e7b0" -> "node_0xc00007e5d0";
  "node_0xc00007e750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007e7b0" -> "node_0xc00007e750";
  "node_0xc00007e630" [label="string: Python", shape="box", fillcolor="lightgray"];
  "node_0xc00007e750" -> "node_0xc00007e630";
  "node_0xc00007e690" [label="string: JavaScript", shape="box", fillcolor="lightgray"];
  "node_0xc00007e750" -> "node_0xc00007e690";
  "node_0xc00007e6f0" [label="string: SQL", shape="box", fillcolor="lightgray"];
  "node_0xc00007e750" -> "node_0xc00007e6f0";
  "node_0xc00007ec30" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007ec30";
  "node_0xc00007e810" [label="string: address", shape="box", fillcolor="lightgray"];
  "node_0xc00007ec30" -> "node_0xc00007e810";
  "node_0xc00007ebd0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007ec30" -> "node_0xc00007ebd0";
  "node_0xc00007e930" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007ebd0" -> "node_0xc00007e930";
  "node_0xc00007e870" [label="string: street", shape="box", fillcolor="lightgray"];
  "node_0xc00007e930" -> "node_0xc00007e870";
  "node_0xc00007e8d0" [label="string: 123 Maple Street", shape="box", fillcolor="lightgray"];
  "node_0xc00007e930" -> "node_0xc00007e8d0";
  "node_0xc00007ea50" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007ebd0" -> "node_0xc00007ea50";
  "node_0xc00007e990" [label="string: city", shape="box", fillcolor="lightgray"];
  "node_0xc00007ea50" -> "node_0xc00007e990";
  "node_0xc00007e9f0" [label="string: Exampleville", shape="box", fillcolor="lightgray"];
  "node_0xc00007ea50" -> "node_0xc00007e9f0";
  "node_0xc00007eb70" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007ebd0" -> "node_0xc00007eb70";
  "node_0xc00007eab0" [label="string: country", shape="box", fillcolor="lightgray"];
  "node_0xc00007eb70" -> "node_0xc00007eab0";
  "node_0xc00007eb10" [label="string: Neverland", shape="box", fillcolor="lightgray"];
  "node_0xc00007eb70" -> "node_0xc00007eb10";
  "node_0xc00007f4d0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f530" -> "node_0xc00007f4d0";
  "node_0xc00007ec90" [label="string: favoriteBooks", shape="box", fillcolor="lightgray"];
  "node_0xc00007f4d0" -> "node_0xc00007ec90";
  "node_0xc00007f470" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007f4d0" -> "node_0xc00007f470";
  "node_0xc00007f050" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007f470" -> "node_0xc00007f050";
  "node_0xc00007edb0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f050" -> "node_0xc00007edb0";
  "node_0xc00007ecf0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00007edb0" -> "node_0xc00007ecf0";
  "node_0xc00007ed50" [label="string: To Kill a Mockingbird", shape="box", fillcolor="lightgray"];
  "node_0xc00007edb0" -> "node_0xc00007ed50";
  "node_0xc00007eed0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f050" -> "node_0xc00007eed0";
  "node_0xc00007ee10" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00007eed0" -> "node_0xc00007ee10";
  "node_0xc00007ee70" [label="string: Harper Lee", shape="box", fillcolor="lightgray"];
  "node_0xc00007eed0" -> "node_0xc00007ee70";
  "node_0xc00007eff0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f050" -> "node_0xc00007eff0";
  "node_0xc00007ef30" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00007eff0" -> "node_0xc00007ef30";
  "node_0xc00007ef90" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007eff0" -> "node_0xc00007ef90";
  "node_0xc00007f410" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007f470" -> "node_0xc00007f410";
  "node_0xc00007f170" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f410" -> "node_0xc00007f170";
  "node_0xc00007f0b0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00007f170" -> "node_0xc00007f0b0";
  "node_0xc00007f110" [label="string: 1984", shape="box", fillcolor="lightgray"];
  "node_0xc00007f170" -> "node_0xc00007f110";
  "node_0xc00007f290" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f410" -> "node_0xc00007f290";
  "node_0xc00007f1d0" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00007f290" -> "node_0xc00007f1d0";
  "node_0xc00007f230" [label="string: George Orwell", shape="box", fillcolor="lightgray"];
  "node_0xc00007f290" -> "node_0xc00007f230";
  "node_0xc00007f3b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007f410" -> "node_0xc00007f3b0";
  "node_0xc00007f2f0" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00007f3b0" -> "node_0xc00007f2f0";
  "node_0xc00007f350" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007f3b0" -> "node_0xc00007f350";
}
```

![Image generated for example](images/json_blob.png)
