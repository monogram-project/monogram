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
  132956878374208["delimited"]:::custom_delimited;
  132956878374288["operator: :"]:::custom_operator;
  132956878374208 --> 132956878374288;
  132956878374368["string: person"]:::custom_string;
  132956878374288 --> 132956878374368;
  132956878374448["delimited"]:::custom_delimited;
  132956878374288 --> 132956878374448;
  132956878374528["operator: :"]:::custom_operator;
  132956878374448 --> 132956878374528;
  132956878374608["string: name"]:::custom_string;
  132956878374528 --> 132956878374608;
  132956878374688["string: Alice"]:::custom_string;
  132956878374528 --> 132956878374688;
  132956878374768["operator: :"]:::custom_operator;
  132956878374448 --> 132956878374768;
  132956878374848["string: age"]:::custom_string;
  132956878374768 --> 132956878374848;
  132956878374928["number: 25"]:::custom_number;
  132956878374768 --> 132956878374928;
  132956878375008["operator: :"]:::custom_operator;
  132956878374448 --> 132956878375008;
  132956878375088["string: isStudent"]:::custom_string;
  132956878375008 --> 132956878375088;
  132956878375168["identifier: true"]:::custom_identifier;
  132956878375008 --> 132956878375168;
  132956878375248["operator: :"]:::custom_operator;
  132956878374448 --> 132956878375248;
  132956878375328["string: skills"]:::custom_string;
  132956878375248 --> 132956878375328;
  132956878375408["delimited"]:::custom_delimited;
  132956878375248 --> 132956878375408;
  132956878375488["string: Python"]:::custom_string;
  132956878375408 --> 132956878375488;
  132956878375568["string: JavaScript"]:::custom_string;
  132956878375408 --> 132956878375568;
  132956878375648["string: SQL"]:::custom_string;
  132956878375408 --> 132956878375648;
  132956878375728["operator: :"]:::custom_operator;
  132956878374448 --> 132956878375728;
  132956878375808["string: address"]:::custom_string;
  132956878375728 --> 132956878375808;
  132956878375888["delimited"]:::custom_delimited;
  132956878375728 --> 132956878375888;
  132956878375968["operator: :"]:::custom_operator;
  132956878375888 --> 132956878375968;
  132956878376048["string: street"]:::custom_string;
  132956878375968 --> 132956878376048;
  132956878376128["string:<br/>123 Maple Street"]:::custom_string;
  132956878375968 --> 132956878376128;
  132956878376208["operator: :"]:::custom_operator;
  132956878375888 --> 132956878376208;
  132956878376288["string: city"]:::custom_string;
  132956878376208 --> 132956878376288;
  132956878376368["string: Exampleville"]:::custom_string;
  132956878376208 --> 132956878376368;
  132956878376448["operator: :"]:::custom_operator;
  132956878375888 --> 132956878376448;
  132956878376528["string: country"]:::custom_string;
  132956878376448 --> 132956878376528;
  132956878376608["string: Neverland"]:::custom_string;
  132956878376448 --> 132956878376608;
  132956878376688["operator: :"]:::custom_operator;
  132956878374448 --> 132956878376688;
  132956878376768["string:<br/>favoriteBooks"]:::custom_string;
  132956878376688 --> 132956878376768;
  132956878376848["delimited"]:::custom_delimited;
  132956878376688 --> 132956878376848;
  132956878376928["delimited"]:::custom_delimited;
  132956878376848 --> 132956878376928;
  132956878377008["operator: :"]:::custom_operator;
  132956878376928 --> 132956878377008;
  132956878377088["string: title"]:::custom_string;
  132956878377008 --> 132956878377088;
  132956878377168["string:<br/>To Kill a Mockingbird"]:::custom_string;
  132956878377008 --> 132956878377168;
  132956878377248["operator: :"]:::custom_operator;
  132956878376928 --> 132956878377248;
  132956878377328["string: author"]:::custom_string;
  132956878377248 --> 132956878377328;
  132956878377408["string: Harper Lee"]:::custom_string;
  132956878377248 --> 132956878377408;
  132956878377488["operator: :"]:::custom_operator;
  132956878376928 --> 132956878377488;
  132956878377568["string:<br/>yearPublished"]:::custom_string;
  132956878377488 --> 132956878377568;
  132956878377648["number: 1960"]:::custom_number;
  132956878377488 --> 132956878377648;
  132956878377728["delimited"]:::custom_delimited;
  132956878376848 --> 132956878377728;
  132956878377808["operator: :"]:::custom_operator;
  132956878377728 --> 132956878377808;
  132956878377888["string: title"]:::custom_string;
  132956878377808 --> 132956878377888;
  132956878377968["string: 1984"]:::custom_string;
  132956878377808 --> 132956878377968;
  132956878378048["operator: :"]:::custom_operator;
  132956878377728 --> 132956878378048;
  132956878378128["string: author"]:::custom_string;
  132956878378048 --> 132956878378128;
  132956878378208["string:<br/>George Orwell"]:::custom_string;
  132956878378048 --> 132956878378208;
  132956878378288["operator: :"]:::custom_operator;
  132956878377728 --> 132956878378288;
  132956878378368["string:<br/>yearPublished"]:::custom_string;
  132956878378288 --> 132956878378368;
  132956878378448["number: 1949"]:::custom_number;
  132956878378288 --> 132956878378448;

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
  "role": "delimited",
  "kind": "braces",
  "separator": "undefined",
  "children": [
    {
      "role": "operator",
      "name": ":",
      "syntax": "infix",
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
                  "value": "name",
                  "quote": "double"
                },
                {
                  "role": "string",
                  "quote": "double",
                  "value": "Alice"
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
                  "value": "isStudent",
                  "quote": "double"
                },
                {
                  "role": "identifier",
                  "name": "true"
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
                      "value": "SQL",
                      "quote": "double"
                    }
                  ]
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
                          "value": "Exampleville",
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
                  "value": "favoriteBooks",
                  "quote": "double"
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
                              "value": "title",
                              "quote": "double"
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
                          "name": ":",
                          "syntax": "infix",
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
  "node_0xc00007d5f0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007d590" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d5f0" -> "node_0xc00007d590";
  "node_0xc00007c210" [label="string: person", shape="box", fillcolor="lightgray"];
  "node_0xc00007d590" -> "node_0xc00007c210";
  "node_0xc00007d530" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007d590" -> "node_0xc00007d530";
  "node_0xc00007c330" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007c330";
  "node_0xc00007c270" [label="string: name", shape="box", fillcolor="lightgray"];
  "node_0xc00007c330" -> "node_0xc00007c270";
  "node_0xc00007c2d0" [label="string: Alice", shape="box", fillcolor="lightgray"];
  "node_0xc00007c330" -> "node_0xc00007c2d0";
  "node_0xc00007c450" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007c450";
  "node_0xc00007c390" [label="string: age", shape="box", fillcolor="lightgray"];
  "node_0xc00007c450" -> "node_0xc00007c390";
  "node_0xc00007c3f0" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c450" -> "node_0xc00007c3f0";
  "node_0xc00007c570" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007c570";
  "node_0xc00007c4b0" [label="string: isStudent", shape="box", fillcolor="lightgray"];
  "node_0xc00007c570" -> "node_0xc00007c4b0";
  "node_0xc00007c510" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c570" -> "node_0xc00007c510";
  "node_0xc00007c7b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007c7b0";
  "node_0xc00007c5d0" [label="string: skills", shape="box", fillcolor="lightgray"];
  "node_0xc00007c7b0" -> "node_0xc00007c5d0";
  "node_0xc00007c750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007c7b0" -> "node_0xc00007c750";
  "node_0xc00007c630" [label="string: Python", shape="box", fillcolor="lightgray"];
  "node_0xc00007c750" -> "node_0xc00007c630";
  "node_0xc00007c690" [label="string: JavaScript", shape="box", fillcolor="lightgray"];
  "node_0xc00007c750" -> "node_0xc00007c690";
  "node_0xc00007c6f0" [label="string: SQL", shape="box", fillcolor="lightgray"];
  "node_0xc00007c750" -> "node_0xc00007c6f0";
  "node_0xc00007cc30" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007cc30";
  "node_0xc00007c810" [label="string: address", shape="box", fillcolor="lightgray"];
  "node_0xc00007cc30" -> "node_0xc00007c810";
  "node_0xc00007cbd0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007cc30" -> "node_0xc00007cbd0";
  "node_0xc00007c930" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007cbd0" -> "node_0xc00007c930";
  "node_0xc00007c870" [label="string: street", shape="box", fillcolor="lightgray"];
  "node_0xc00007c930" -> "node_0xc00007c870";
  "node_0xc00007c8d0" [label="string: 123 Maple Street", shape="box", fillcolor="lightgray"];
  "node_0xc00007c930" -> "node_0xc00007c8d0";
  "node_0xc00007ca50" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007cbd0" -> "node_0xc00007ca50";
  "node_0xc00007c990" [label="string: city", shape="box", fillcolor="lightgray"];
  "node_0xc00007ca50" -> "node_0xc00007c990";
  "node_0xc00007c9f0" [label="string: Exampleville", shape="box", fillcolor="lightgray"];
  "node_0xc00007ca50" -> "node_0xc00007c9f0";
  "node_0xc00007cb70" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007cbd0" -> "node_0xc00007cb70";
  "node_0xc00007cab0" [label="string: country", shape="box", fillcolor="lightgray"];
  "node_0xc00007cb70" -> "node_0xc00007cab0";
  "node_0xc00007cb10" [label="string: Neverland", shape="box", fillcolor="lightgray"];
  "node_0xc00007cb70" -> "node_0xc00007cb10";
  "node_0xc00007d4d0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d530" -> "node_0xc00007d4d0";
  "node_0xc00007cc90" [label="string: favoriteBooks", shape="box", fillcolor="lightgray"];
  "node_0xc00007d4d0" -> "node_0xc00007cc90";
  "node_0xc00007d470" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007d4d0" -> "node_0xc00007d470";
  "node_0xc00007d050" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007d470" -> "node_0xc00007d050";
  "node_0xc00007cdb0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d050" -> "node_0xc00007cdb0";
  "node_0xc00007ccf0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00007cdb0" -> "node_0xc00007ccf0";
  "node_0xc00007cd50" [label="string: To Kill a Mockingbird", shape="box", fillcolor="lightgray"];
  "node_0xc00007cdb0" -> "node_0xc00007cd50";
  "node_0xc00007ced0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d050" -> "node_0xc00007ced0";
  "node_0xc00007ce10" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00007ced0" -> "node_0xc00007ce10";
  "node_0xc00007ce70" [label="string: Harper Lee", shape="box", fillcolor="lightgray"];
  "node_0xc00007ced0" -> "node_0xc00007ce70";
  "node_0xc00007cff0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d050" -> "node_0xc00007cff0";
  "node_0xc00007cf30" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00007cff0" -> "node_0xc00007cf30";
  "node_0xc00007cf90" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007cff0" -> "node_0xc00007cf90";
  "node_0xc00007d410" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00007d470" -> "node_0xc00007d410";
  "node_0xc00007d170" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d410" -> "node_0xc00007d170";
  "node_0xc00007d0b0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00007d170" -> "node_0xc00007d0b0";
  "node_0xc00007d110" [label="string: 1984", shape="box", fillcolor="lightgray"];
  "node_0xc00007d170" -> "node_0xc00007d110";
  "node_0xc00007d290" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d410" -> "node_0xc00007d290";
  "node_0xc00007d1d0" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00007d290" -> "node_0xc00007d1d0";
  "node_0xc00007d230" [label="string: George Orwell", shape="box", fillcolor="lightgray"];
  "node_0xc00007d290" -> "node_0xc00007d230";
  "node_0xc00007d3b0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007d410" -> "node_0xc00007d3b0";
  "node_0xc00007d2f0" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00007d3b0" -> "node_0xc00007d2f0";
  "node_0xc00007d350" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007d3b0" -> "node_0xc00007d350";
}
```

![Image generated for example](images/json_blob.png)
