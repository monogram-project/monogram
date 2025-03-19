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
  134131508954432["delimited"]:::custom_delimited;
  134131508954512["operator: :"]:::custom_operator;
  134131508954432 --> 134131508954512;
  134131508954592["string: person"]:::custom_string;
  134131508954512 --> 134131508954592;
  134131508954672["delimited"]:::custom_delimited;
  134131508954512 --> 134131508954672;
  134131508954752["operator: :"]:::custom_operator;
  134131508954672 --> 134131508954752;
  134131508954832["string: name"]:::custom_string;
  134131508954752 --> 134131508954832;
  134131508954912["string: Alice"]:::custom_string;
  134131508954752 --> 134131508954912;
  134131508954992["operator: :"]:::custom_operator;
  134131508954672 --> 134131508954992;
  134131508955072["string: age"]:::custom_string;
  134131508954992 --> 134131508955072;
  134131508955152["number: 25"]:::custom_number;
  134131508954992 --> 134131508955152;
  134131508955232["operator: :"]:::custom_operator;
  134131508954672 --> 134131508955232;
  134131508955312["string: isStudent"]:::custom_string;
  134131508955232 --> 134131508955312;
  134131508955392["identifier: true"]:::custom_identifier;
  134131508955232 --> 134131508955392;
  134131508955472["operator: :"]:::custom_operator;
  134131508954672 --> 134131508955472;
  134131508955552["string: skills"]:::custom_string;
  134131508955472 --> 134131508955552;
  134131508955632["delimited"]:::custom_delimited;
  134131508955472 --> 134131508955632;
  134131508955712["string: Python"]:::custom_string;
  134131508955632 --> 134131508955712;
  134131508955792["string: JavaScript"]:::custom_string;
  134131508955632 --> 134131508955792;
  134131508955872["string: SQL"]:::custom_string;
  134131508955632 --> 134131508955872;
  134131508955952["operator: :"]:::custom_operator;
  134131508954672 --> 134131508955952;
  134131508956032["string: address"]:::custom_string;
  134131508955952 --> 134131508956032;
  134131508956112["delimited"]:::custom_delimited;
  134131508955952 --> 134131508956112;
  134131508956192["operator: :"]:::custom_operator;
  134131508956112 --> 134131508956192;
  134131508956272["string: street"]:::custom_string;
  134131508956192 --> 134131508956272;
  134131508956352["string:<br/>123 Maple Street"]:::custom_string;
  134131508956192 --> 134131508956352;
  134131508956432["operator: :"]:::custom_operator;
  134131508956112 --> 134131508956432;
  134131508956512["string: city"]:::custom_string;
  134131508956432 --> 134131508956512;
  134131508956592["string: Exampleville"]:::custom_string;
  134131508956432 --> 134131508956592;
  134131508956672["operator: :"]:::custom_operator;
  134131508956112 --> 134131508956672;
  134131508956752["string: country"]:::custom_string;
  134131508956672 --> 134131508956752;
  134131508956832["string: Neverland"]:::custom_string;
  134131508956672 --> 134131508956832;
  134131508956912["operator: :"]:::custom_operator;
  134131508954672 --> 134131508956912;
  134131508956992["string:<br/>favoriteBooks"]:::custom_string;
  134131508956912 --> 134131508956992;
  134131508957072["delimited"]:::custom_delimited;
  134131508956912 --> 134131508957072;
  134131508957152["delimited"]:::custom_delimited;
  134131508957072 --> 134131508957152;
  134131508957232["operator: :"]:::custom_operator;
  134131508957152 --> 134131508957232;
  134131508957312["string: title"]:::custom_string;
  134131508957232 --> 134131508957312;
  134131508957392["string:<br/>To Kill a Mockingbird"]:::custom_string;
  134131508957232 --> 134131508957392;
  134131508957472["operator: :"]:::custom_operator;
  134131508957152 --> 134131508957472;
  134131508957552["string: author"]:::custom_string;
  134131508957472 --> 134131508957552;
  134131508957632["string: Harper Lee"]:::custom_string;
  134131508957472 --> 134131508957632;
  134131508957712["operator: :"]:::custom_operator;
  134131508957152 --> 134131508957712;
  134131508957792["string:<br/>yearPublished"]:::custom_string;
  134131508957712 --> 134131508957792;
  134131508957872["number: 1960"]:::custom_number;
  134131508957712 --> 134131508957872;
  134131508957952["delimited"]:::custom_delimited;
  134131508957072 --> 134131508957952;
  134131508958032["operator: :"]:::custom_operator;
  134131508957952 --> 134131508958032;
  134131508958112["string: title"]:::custom_string;
  134131508958032 --> 134131508958112;
  134131508958192["string: 1984"]:::custom_string;
  134131508958032 --> 134131508958192;
  134131508958272["operator: :"]:::custom_operator;
  134131508957952 --> 134131508958272;
  134131508958352["string: author"]:::custom_string;
  134131508958272 --> 134131508958352;
  134131508958432["string:<br/>George Orwell"]:::custom_string;
  134131508958272 --> 134131508958432;
  134131508958512["operator: :"]:::custom_operator;
  134131508957952 --> 134131508958512;
  134131508958592["string:<br/>yearPublished"]:::custom_string;
  134131508958512 --> 134131508958592;
  134131508958672["number: 1949"]:::custom_number;
  134131508958512 --> 134131508958672;

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
  <operator name=":" syntax="infix">
    <string quote="double" value="person" />
    <delimited kind="braces" separator="comma">
      <operator name=":" syntax="infix">
        <string quote="double" value="name" />
        <string quote="double" value="Alice" />
      </operator>
      <operator name=":" syntax="infix">
        <string quote="double" value="age" />
        <number value="25" />
      </operator>
      <operator name=":" syntax="infix">
        <string quote="double" value="isStudent" />
        <identifier name="true" />
      </operator>
      <operator name=":" syntax="infix">
        <string quote="double" value="skills" />
        <delimited kind="brackets" separator="comma">
          <string quote="double" value="Python" />
          <string quote="double" value="JavaScript" />
          <string quote="double" value="SQL" />
        </delimited>
      </operator>
      <operator name=":" syntax="infix">
        <string quote="double" value="address" />
        <delimited kind="braces" separator="comma">
          <operator name=":" syntax="infix">
            <string quote="double" value="street" />
            <string quote="double" value="123 Maple Street" />
          </operator>
          <operator name=":" syntax="infix">
            <string quote="double" value="city" />
            <string quote="double" value="Exampleville" />
          </operator>
          <operator name=":" syntax="infix">
            <string quote="double" value="country" />
            <string quote="double" value="Neverland" />
          </operator>
        </delimited>
      </operator>
      <operator name=":" syntax="infix">
        <string quote="double" value="favoriteBooks" />
        <delimited kind="brackets" separator="comma">
          <delimited kind="braces" separator="comma">
            <operator name=":" syntax="infix">
              <string quote="double" value="title" />
              <string quote="double" value="To Kill a Mockingbird" />
            </operator>
            <operator name=":" syntax="infix">
              <string quote="double" value="author" />
              <string quote="double" value="Harper Lee" />
            </operator>
            <operator name=":" syntax="infix">
              <string quote="double" value="yearPublished" />
              <number value="1960" />
            </operator>
          </delimited>
          <delimited kind="braces" separator="comma">
            <operator name=":" syntax="infix">
              <string quote="double" value="title" />
              <string quote="double" value="1984" />
            </operator>
            <operator name=":" syntax="infix">
              <string quote="double" value="author" />
              <string quote="double" value="George Orwell" />
            </operator>
            <operator name=":" syntax="infix">
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
                  "name": ":",
                  "syntax": "infix",
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
                      "quote": "double",
                      "value": "skills"
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
                      "separator": "comma",
                      "kind": "braces",
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
                          "name": ":",
                          "syntax": "infix",
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
                          "syntax": "infix",
                          "name": ":",
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
                                  "quote": "double",
                                  "value": "title"
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
  name: ':'
  syntax: infix
  children:
  - role: string
    quote: double
    value: person
  - role: delimited
    kind: braces
    separator: comma
    children:
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: double
        value: name
      - role: string
        quote: double
        value: Alice
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: double
        value: age
      - role: number
        value: 25
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: double
        value: isStudent
      - role: identifier
        name: 'true'
    - role: operator
      name: ':'
      syntax: infix
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
          quote: double
          value: JavaScript
        - role: string
          quote: double
          value: SQL
    - role: operator
      name: ':'
      syntax: infix
      children:
      - role: string
        quote: double
        value: address
      - role: delimited
        kind: braces
        separator: comma
        children:
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: double
            value: street
          - role: string
            quote: double
            value: 123 Maple Street
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: double
            value: city
          - role: string
            quote: double
            value: Exampleville
        - role: operator
          name: ':'
          syntax: infix
          children:
          - role: string
            quote: double
            value: country
          - role: string
            quote: double
            value: Neverland
    - role: operator
      name: ':'
      syntax: infix
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
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: double
              value: title
            - role: string
              quote: double
              value: To Kill a Mockingbird
          - role: operator
            name: ':'
            syntax: infix
            children:
            - role: string
              quote: double
              value: author
            - role: string
              quote: double
              value: Harper Lee
          - role: operator
            name: ':'
            syntax: infix
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
            name: ':'
            syntax: infix
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
            name: ':'
            syntax: infix
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
  "node_0xc0000a9650" [label="unit: json_blob.mg", shape="box", fillcolor="lightgray"];
  "node_0xc0000a95f0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9650" -> "node_0xc0000a95f0";
  "node_0xc0000a9590" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a95f0" -> "node_0xc0000a9590";
  "node_0xc0000a8210" [label="string: person", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9590" -> "node_0xc0000a8210";
  "node_0xc0000a9530" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9590" -> "node_0xc0000a9530";
  "node_0xc0000a8330" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a8330";
  "node_0xc0000a8270" [label="string: name", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8330" -> "node_0xc0000a8270";
  "node_0xc0000a82d0" [label="string: Alice", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8330" -> "node_0xc0000a82d0";
  "node_0xc0000a8450" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a8450";
  "node_0xc0000a8390" [label="string: age", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8450" -> "node_0xc0000a8390";
  "node_0xc0000a83f0" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8450" -> "node_0xc0000a83f0";
  "node_0xc0000a8570" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a8570";
  "node_0xc0000a84b0" [label="string: isStudent", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8570" -> "node_0xc0000a84b0";
  "node_0xc0000a8510" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8570" -> "node_0xc0000a8510";
  "node_0xc0000a87b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a87b0";
  "node_0xc0000a85d0" [label="string: skills", shape="box", fillcolor="lightgray"];
  "node_0xc0000a87b0" -> "node_0xc0000a85d0";
  "node_0xc0000a8750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a87b0" -> "node_0xc0000a8750";
  "node_0xc0000a8630" [label="string: Python", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8750" -> "node_0xc0000a8630";
  "node_0xc0000a8690" [label="string: JavaScript", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8750" -> "node_0xc0000a8690";
  "node_0xc0000a86f0" [label="string: SQL", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8750" -> "node_0xc0000a86f0";
  "node_0xc0000a8c30" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a8c30";
  "node_0xc0000a8810" [label="string: address", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8c30" -> "node_0xc0000a8810";
  "node_0xc0000a8bd0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8c30" -> "node_0xc0000a8bd0";
  "node_0xc0000a8930" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8bd0" -> "node_0xc0000a8930";
  "node_0xc0000a8870" [label="string: street", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8930" -> "node_0xc0000a8870";
  "node_0xc0000a88d0" [label="string: 123 Maple Street", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8930" -> "node_0xc0000a88d0";
  "node_0xc0000a8a50" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8bd0" -> "node_0xc0000a8a50";
  "node_0xc0000a8990" [label="string: city", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8a50" -> "node_0xc0000a8990";
  "node_0xc0000a89f0" [label="string: Exampleville", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8a50" -> "node_0xc0000a89f0";
  "node_0xc0000a8b70" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8bd0" -> "node_0xc0000a8b70";
  "node_0xc0000a8ab0" [label="string: country", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8b70" -> "node_0xc0000a8ab0";
  "node_0xc0000a8b10" [label="string: Neverland", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8b70" -> "node_0xc0000a8b10";
  "node_0xc0000a94d0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9530" -> "node_0xc0000a94d0";
  "node_0xc0000a8c90" [label="string: favoriteBooks", shape="box", fillcolor="lightgray"];
  "node_0xc0000a94d0" -> "node_0xc0000a8c90";
  "node_0xc0000a9470" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a94d0" -> "node_0xc0000a9470";
  "node_0xc0000a9050" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9470" -> "node_0xc0000a9050";
  "node_0xc0000a8db0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9050" -> "node_0xc0000a8db0";
  "node_0xc0000a8cf0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8db0" -> "node_0xc0000a8cf0";
  "node_0xc0000a8d50" [label="string: To Kill a Mockingbird", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8db0" -> "node_0xc0000a8d50";
  "node_0xc0000a8ed0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9050" -> "node_0xc0000a8ed0";
  "node_0xc0000a8e10" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8ed0" -> "node_0xc0000a8e10";
  "node_0xc0000a8e70" [label="string: Harper Lee", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8ed0" -> "node_0xc0000a8e70";
  "node_0xc0000a8ff0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9050" -> "node_0xc0000a8ff0";
  "node_0xc0000a8f30" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8ff0" -> "node_0xc0000a8f30";
  "node_0xc0000a8f90" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8ff0" -> "node_0xc0000a8f90";
  "node_0xc0000a9410" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9470" -> "node_0xc0000a9410";
  "node_0xc0000a9170" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9410" -> "node_0xc0000a9170";
  "node_0xc0000a90b0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9170" -> "node_0xc0000a90b0";
  "node_0xc0000a9110" [label="string: 1984", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9170" -> "node_0xc0000a9110";
  "node_0xc0000a9290" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9410" -> "node_0xc0000a9290";
  "node_0xc0000a91d0" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9290" -> "node_0xc0000a91d0";
  "node_0xc0000a9230" [label="string: George Orwell", shape="box", fillcolor="lightgray"];
  "node_0xc0000a9290" -> "node_0xc0000a9230";
  "node_0xc0000a93b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a9410" -> "node_0xc0000a93b0";
  "node_0xc0000a92f0" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc0000a93b0" -> "node_0xc0000a92f0";
  "node_0xc0000a9350" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a93b0" -> "node_0xc0000a9350";
}
```

![Image generated for example](images/json_blob.png)
