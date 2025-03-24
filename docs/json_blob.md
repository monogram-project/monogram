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
  node_0xc0000afaa0["delimited"]:::custom_delimited;
  node_0xc0000afa40["operator: :"]:::custom_operator;
  node_0xc0000afaa0 --> node_0xc0000afa40;
  node_0xc0000ae6c0["string: person"]:::custom_string;
  node_0xc0000afa40 --> node_0xc0000ae6c0;
  node_0xc0000af9e0["delimited"]:::custom_delimited;
  node_0xc0000afa40 --> node_0xc0000af9e0;
  node_0xc0000ae7e0["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000ae7e0;
  node_0xc0000ae720["string: name"]:::custom_string;
  node_0xc0000ae7e0 --> node_0xc0000ae720;
  node_0xc0000ae780["string: Alice"]:::custom_string;
  node_0xc0000ae7e0 --> node_0xc0000ae780;
  node_0xc0000ae900["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000ae900;
  node_0xc0000ae840["string: age"]:::custom_string;
  node_0xc0000ae900 --> node_0xc0000ae840;
  node_0xc0000ae8a0["number: 25"]:::custom_number;
  node_0xc0000ae900 --> node_0xc0000ae8a0;
  node_0xc0000aea20["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000aea20;
  node_0xc0000ae960["string: isStudent"]:::custom_string;
  node_0xc0000aea20 --> node_0xc0000ae960;
  node_0xc0000ae9c0["identifier: true"]:::custom_identifier;
  node_0xc0000aea20 --> node_0xc0000ae9c0;
  node_0xc0000aec60["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000aec60;
  node_0xc0000aea80["string: skills"]:::custom_string;
  node_0xc0000aec60 --> node_0xc0000aea80;
  node_0xc0000aec00["delimited"]:::custom_delimited;
  node_0xc0000aec60 --> node_0xc0000aec00;
  node_0xc0000aeae0["string: Python"]:::custom_string;
  node_0xc0000aec00 --> node_0xc0000aeae0;
  node_0xc0000aeb40["string: JavaScript"]:::custom_string;
  node_0xc0000aec00 --> node_0xc0000aeb40;
  node_0xc0000aeba0["string: SQL"]:::custom_string;
  node_0xc0000aec00 --> node_0xc0000aeba0;
  node_0xc0000af0e0["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000af0e0;
  node_0xc0000aecc0["string: address"]:::custom_string;
  node_0xc0000af0e0 --> node_0xc0000aecc0;
  node_0xc0000af080["delimited"]:::custom_delimited;
  node_0xc0000af0e0 --> node_0xc0000af080;
  node_0xc0000aede0["operator: :"]:::custom_operator;
  node_0xc0000af080 --> node_0xc0000aede0;
  node_0xc0000aed20["string: street"]:::custom_string;
  node_0xc0000aede0 --> node_0xc0000aed20;
  node_0xc0000aed80["string:<br/>123 Maple Street"]:::custom_string;
  node_0xc0000aede0 --> node_0xc0000aed80;
  node_0xc0000aef00["operator: :"]:::custom_operator;
  node_0xc0000af080 --> node_0xc0000aef00;
  node_0xc0000aee40["string: city"]:::custom_string;
  node_0xc0000aef00 --> node_0xc0000aee40;
  node_0xc0000aeea0["string: Exampleville"]:::custom_string;
  node_0xc0000aef00 --> node_0xc0000aeea0;
  node_0xc0000af020["operator: :"]:::custom_operator;
  node_0xc0000af080 --> node_0xc0000af020;
  node_0xc0000aef60["string: country"]:::custom_string;
  node_0xc0000af020 --> node_0xc0000aef60;
  node_0xc0000aefc0["string: Neverland"]:::custom_string;
  node_0xc0000af020 --> node_0xc0000aefc0;
  node_0xc0000af980["operator: :"]:::custom_operator;
  node_0xc0000af9e0 --> node_0xc0000af980;
  node_0xc0000af140["string:<br/>favoriteBooks"]:::custom_string;
  node_0xc0000af980 --> node_0xc0000af140;
  node_0xc0000af920["delimited"]:::custom_delimited;
  node_0xc0000af980 --> node_0xc0000af920;
  node_0xc0000af500["delimited"]:::custom_delimited;
  node_0xc0000af920 --> node_0xc0000af500;
  node_0xc0000af260["operator: :"]:::custom_operator;
  node_0xc0000af500 --> node_0xc0000af260;
  node_0xc0000af1a0["string: title"]:::custom_string;
  node_0xc0000af260 --> node_0xc0000af1a0;
  node_0xc0000af200["string:<br/>To Kill a Mockingbird"]:::custom_string;
  node_0xc0000af260 --> node_0xc0000af200;
  node_0xc0000af380["operator: :"]:::custom_operator;
  node_0xc0000af500 --> node_0xc0000af380;
  node_0xc0000af2c0["string: author"]:::custom_string;
  node_0xc0000af380 --> node_0xc0000af2c0;
  node_0xc0000af320["string: Harper Lee"]:::custom_string;
  node_0xc0000af380 --> node_0xc0000af320;
  node_0xc0000af4a0["operator: :"]:::custom_operator;
  node_0xc0000af500 --> node_0xc0000af4a0;
  node_0xc0000af3e0["string:<br/>yearPublished"]:::custom_string;
  node_0xc0000af4a0 --> node_0xc0000af3e0;
  node_0xc0000af440["number: 1960"]:::custom_number;
  node_0xc0000af4a0 --> node_0xc0000af440;
  node_0xc0000af8c0["delimited"]:::custom_delimited;
  node_0xc0000af920 --> node_0xc0000af8c0;
  node_0xc0000af620["operator: :"]:::custom_operator;
  node_0xc0000af8c0 --> node_0xc0000af620;
  node_0xc0000af560["string: title"]:::custom_string;
  node_0xc0000af620 --> node_0xc0000af560;
  node_0xc0000af5c0["string: 1984"]:::custom_string;
  node_0xc0000af620 --> node_0xc0000af5c0;
  node_0xc0000af740["operator: :"]:::custom_operator;
  node_0xc0000af8c0 --> node_0xc0000af740;
  node_0xc0000af680["string: author"]:::custom_string;
  node_0xc0000af740 --> node_0xc0000af680;
  node_0xc0000af6e0["string:<br/>George Orwell"]:::custom_string;
  node_0xc0000af740 --> node_0xc0000af6e0;
  node_0xc0000af860["operator: :"]:::custom_operator;
  node_0xc0000af8c0 --> node_0xc0000af860;
  node_0xc0000af7a0["string:<br/>yearPublished"]:::custom_string;
  node_0xc0000af860 --> node_0xc0000af7a0;
  node_0xc0000af800["number: 1949"]:::custom_number;
  node_0xc0000af860 --> node_0xc0000af800;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;
classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
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
                  "separator": "comma",
                  "kind": "brackets",
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
                          "value": "country",
                          "quote": "double"
                        },
                        {
                          "role": "string",
                          "value": "Neverland",
                          "quote": "double"
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
                              "value": "Harper Lee",
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
                          "name": ":",
                          "syntax": "infix",
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
kind: braces
separator: undefined
children:
- role: operator
  name: ":"
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
      name: ":"
      syntax: infix
      children:
      - role: string
        quote: double
        value: name
      - role: string
        quote: double
        value: Alice
    - role: operator
      name: ":"
      syntax: infix
      children:
      - role: string
        quote: double
        value: age
      - role: number
        value: 25
    - role: operator
      name: ":"
      syntax: infix
      children:
      - role: string
        quote: double
        value: isStudent
      - role: identifier
        name: true
    - role: operator
      name: ":"
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
      name: ":"
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
          name: ":"
          syntax: infix
          children:
          - role: string
            quote: double
            value: street
          - role: string
            quote: double
            value: 123 Maple Street
        - role: operator
          name: ":"
          syntax: infix
          children:
          - role: string
            quote: double
            value: city
          - role: string
            quote: double
            value: Exampleville
        - role: operator
          name: ":"
          syntax: infix
          children:
          - role: string
            quote: double
            value: country
          - role: string
            quote: double
            value: Neverland
    - role: operator
      name: ":"
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
            name: ":"
            syntax: infix
            children:
            - role: string
              quote: double
              value: title
            - role: string
              quote: double
              value: To Kill a Mockingbird
          - role: operator
            name: ":"
            syntax: infix
            children:
            - role: string
              quote: double
              value: author
            - role: string
              quote: double
              value: Harper Lee
          - role: operator
            name: ":"
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
            name: ":"
            syntax: infix
            children:
            - role: string
              quote: double
              value: title
            - role: string
              quote: double
              value: 1984
          - role: operator
            name: ":"
            syntax: infix
            children:
            - role: string
              quote: double
              value: author
            - role: string
              quote: double
              value: George Orwell
          - role: operator
            name: ":"
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
  "node_0xc00011baa0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011ba40" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011baa0" -> "node_0xc00011ba40";
  "node_0xc00011a6c0" [label="string: person", shape="box", fillcolor="lightgray"];
  "node_0xc00011ba40" -> "node_0xc00011a6c0";
  "node_0xc00011b9e0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011ba40" -> "node_0xc00011b9e0";
  "node_0xc00011a7e0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011a7e0";
  "node_0xc00011a720" [label="string: name", shape="box", fillcolor="lightgray"];
  "node_0xc00011a7e0" -> "node_0xc00011a720";
  "node_0xc00011a780" [label="string: Alice", shape="box", fillcolor="lightgray"];
  "node_0xc00011a7e0" -> "node_0xc00011a780";
  "node_0xc00011a900" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011a900";
  "node_0xc00011a840" [label="string: age", shape="box", fillcolor="lightgray"];
  "node_0xc00011a900" -> "node_0xc00011a840";
  "node_0xc00011a8a0" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00011a900" -> "node_0xc00011a8a0";
  "node_0xc00011aa20" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011aa20";
  "node_0xc00011a960" [label="string: isStudent", shape="box", fillcolor="lightgray"];
  "node_0xc00011aa20" -> "node_0xc00011a960";
  "node_0xc00011a9c0" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "node_0xc00011aa20" -> "node_0xc00011a9c0";
  "node_0xc00011ac60" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011ac60";
  "node_0xc00011aa80" [label="string: skills", shape="box", fillcolor="lightgray"];
  "node_0xc00011ac60" -> "node_0xc00011aa80";
  "node_0xc00011ac00" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011ac60" -> "node_0xc00011ac00";
  "node_0xc00011aae0" [label="string: Python", shape="box", fillcolor="lightgray"];
  "node_0xc00011ac00" -> "node_0xc00011aae0";
  "node_0xc00011ab40" [label="string: JavaScript", shape="box", fillcolor="lightgray"];
  "node_0xc00011ac00" -> "node_0xc00011ab40";
  "node_0xc00011aba0" [label="string: SQL", shape="box", fillcolor="lightgray"];
  "node_0xc00011ac00" -> "node_0xc00011aba0";
  "node_0xc00011b0e0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011b0e0";
  "node_0xc00011acc0" [label="string: address", shape="box", fillcolor="lightgray"];
  "node_0xc00011b0e0" -> "node_0xc00011acc0";
  "node_0xc00011b080" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011b0e0" -> "node_0xc00011b080";
  "node_0xc00011ade0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b080" -> "node_0xc00011ade0";
  "node_0xc00011ad20" [label="string: street", shape="box", fillcolor="lightgray"];
  "node_0xc00011ade0" -> "node_0xc00011ad20";
  "node_0xc00011ad80" [label="string: 123 Maple Street", shape="box", fillcolor="lightgray"];
  "node_0xc00011ade0" -> "node_0xc00011ad80";
  "node_0xc00011af00" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b080" -> "node_0xc00011af00";
  "node_0xc00011ae40" [label="string: city", shape="box", fillcolor="lightgray"];
  "node_0xc00011af00" -> "node_0xc00011ae40";
  "node_0xc00011aea0" [label="string: Exampleville", shape="box", fillcolor="lightgray"];
  "node_0xc00011af00" -> "node_0xc00011aea0";
  "node_0xc00011b020" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b080" -> "node_0xc00011b020";
  "node_0xc00011af60" [label="string: country", shape="box", fillcolor="lightgray"];
  "node_0xc00011b020" -> "node_0xc00011af60";
  "node_0xc00011afc0" [label="string: Neverland", shape="box", fillcolor="lightgray"];
  "node_0xc00011b020" -> "node_0xc00011afc0";
  "node_0xc00011b980" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b9e0" -> "node_0xc00011b980";
  "node_0xc00011b140" [label="string: favoriteBooks", shape="box", fillcolor="lightgray"];
  "node_0xc00011b980" -> "node_0xc00011b140";
  "node_0xc00011b920" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011b980" -> "node_0xc00011b920";
  "node_0xc00011b500" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011b920" -> "node_0xc00011b500";
  "node_0xc00011b260" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b500" -> "node_0xc00011b260";
  "node_0xc00011b1a0" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00011b260" -> "node_0xc00011b1a0";
  "node_0xc00011b200" [label="string: To Kill a Mockingbird", shape="box", fillcolor="lightgray"];
  "node_0xc00011b260" -> "node_0xc00011b200";
  "node_0xc00011b380" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b500" -> "node_0xc00011b380";
  "node_0xc00011b2c0" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00011b380" -> "node_0xc00011b2c0";
  "node_0xc00011b320" [label="string: Harper Lee", shape="box", fillcolor="lightgray"];
  "node_0xc00011b380" -> "node_0xc00011b320";
  "node_0xc00011b4a0" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b500" -> "node_0xc00011b4a0";
  "node_0xc00011b3e0" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00011b4a0" -> "node_0xc00011b3e0";
  "node_0xc00011b440" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00011b4a0" -> "node_0xc00011b440";
  "node_0xc00011b8c0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00011b920" -> "node_0xc00011b8c0";
  "node_0xc00011b620" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b8c0" -> "node_0xc00011b620";
  "node_0xc00011b560" [label="string: title", shape="box", fillcolor="lightgray"];
  "node_0xc00011b620" -> "node_0xc00011b560";
  "node_0xc00011b5c0" [label="string: 1984", shape="box", fillcolor="lightgray"];
  "node_0xc00011b620" -> "node_0xc00011b5c0";
  "node_0xc00011b740" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b8c0" -> "node_0xc00011b740";
  "node_0xc00011b680" [label="string: author", shape="box", fillcolor="lightgray"];
  "node_0xc00011b740" -> "node_0xc00011b680";
  "node_0xc00011b6e0" [label="string: George Orwell", shape="box", fillcolor="lightgray"];
  "node_0xc00011b740" -> "node_0xc00011b6e0";
  "node_0xc00011b860" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00011b8c0" -> "node_0xc00011b860";
  "node_0xc00011b7a0" [label="string: yearPublished", shape="box", fillcolor="lightgray"];
  "node_0xc00011b860" -> "node_0xc00011b7a0";
  "node_0xc00011b800" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00011b860" -> "node_0xc00011b800";
}
```

![Image generated for example](images/json_blob.png)
