# JSON Expression

## Monogram

Monograph supports JSON expressions:

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
  139731259196736["delimited"]:::custom_delimited;
  139731259196816["operator: :"]:::custom_operator;
  139731259196736 --> 139731259196816;
  139731259196896["string: person"]:::custom_string;
  139731259196816 --> 139731259196896;
  139731259196976["delimited"]:::custom_delimited;
  139731259196816 --> 139731259196976;
  139731259197056["operator: :"]:::custom_operator;
  139731259196976 --> 139731259197056;
  139731259197136["string: name"]:::custom_string;
  139731259197056 --> 139731259197136;
  139731259197216["string: Alice"]:::custom_string;
  139731259197056 --> 139731259197216;
  139731259197296["operator: :"]:::custom_operator;
  139731259196976 --> 139731259197296;
  139731259197376["string: age"]:::custom_string;
  139731259197296 --> 139731259197376;
  139731259197456["number: 25"]:::custom_number;
  139731259197296 --> 139731259197456;
  139731259197536["operator: :"]:::custom_operator;
  139731259196976 --> 139731259197536;
  139731259197616["string: isStudent"]:::custom_string;
  139731259197536 --> 139731259197616;
  139731259197696["identifier: true"]:::custom_identifier;
  139731259197536 --> 139731259197696;
  139731259197776["operator: :"]:::custom_operator;
  139731259196976 --> 139731259197776;
  139731259197856["string: skills"]:::custom_string;
  139731259197776 --> 139731259197856;
  139731259197936["delimited"]:::custom_delimited;
  139731259197776 --> 139731259197936;
  139731259198016["string: Python"]:::custom_string;
  139731259197936 --> 139731259198016;
  139731259198096["string: JavaScript"]:::custom_string;
  139731259197936 --> 139731259198096;
  139731259198176["string: SQL"]:::custom_string;
  139731259197936 --> 139731259198176;
  139731259198256["operator: :"]:::custom_operator;
  139731259196976 --> 139731259198256;
  139731259198336["string: address"]:::custom_string;
  139731259198256 --> 139731259198336;
  139731259198416["delimited"]:::custom_delimited;
  139731259198256 --> 139731259198416;
  139731259198496["operator: :"]:::custom_operator;
  139731259198416 --> 139731259198496;
  139731259198656["string: street"]:::custom_string;
  139731259198496 --> 139731259198656;
  139731259198816["string: 123 Maple Street"]:::custom_string;
  139731259198496 --> 139731259198816;
  139731259198896["operator: :"]:::custom_operator;
  139731259198416 --> 139731259198896;
  139731259199056["string: city"]:::custom_string;
  139731259198896 --> 139731259199056;
  139731259199216["string: Exampleville"]:::custom_string;
  139731259198896 --> 139731259199216;
  139731259199296["operator: :"]:::custom_operator;
  139731259198416 --> 139731259199296;
  139731259199456["string: country"]:::custom_string;
  139731259199296 --> 139731259199456;
  139731259199616["string: Neverland"]:::custom_string;
  139731259199296 --> 139731259199616;
  139731259199696["operator: :"]:::custom_operator;
  139731259196976 --> 139731259199696;
  139731259199776["string: favoriteBooks"]:::custom_string;
  139731259199696 --> 139731259199776;
  139731259199856["delimited"]:::custom_delimited;
  139731259199696 --> 139731259199856;
  139731259199936["delimited"]:::custom_delimited;
  139731259199856 --> 139731259199936;
  139731259200096["operator: :"]:::custom_operator;
  139731259199936 --> 139731259200096;
  139731259200256["string: title"]:::custom_string;
  139731259200096 --> 139731259200256;
  139731259200416["string: To Kill a Mockingbird"]:::custom_string;
  139731259200096 --> 139731259200416;
  139731259200656["operator: :"]:::custom_operator;
  139731259199936 --> 139731259200656;
  139731259200816["string: author"]:::custom_string;
  139731259200656 --> 139731259200816;
  139731259200976["string: Harper Lee"]:::custom_string;
  139731259200656 --> 139731259200976;
  139731259201216["operator: :"]:::custom_operator;
  139731259199936 --> 139731259201216;
  139731259201376["string: yearPublished"]:::custom_string;
  139731259201216 --> 139731259201376;
  139731259201536["number: 1960"]:::custom_number;
  139731259201216 --> 139731259201536;
  139731259201696["delimited"]:::custom_delimited;
  139731259199856 --> 139731259201696;
  139731259201856["operator: :"]:::custom_operator;
  139731259201696 --> 139731259201856;
  139731259202016["string: title"]:::custom_string;
  139731259201856 --> 139731259202016;
  139731259202176["string: 1984"]:::custom_string;
  139731259201856 --> 139731259202176;
  139731259202416["operator: :"]:::custom_operator;
  139731259201696 --> 139731259202416;
  139731259202576["string: author"]:::custom_string;
  139731259202416 --> 139731259202576;
  139731259202736["string: George Orwell"]:::custom_string;
  139731259202416 --> 139731259202736;
  139731259202976["operator: :"]:::custom_operator;
  139731259201696 --> 139731259202976;
  139731259203136["string: yearPublished"]:::custom_string;
  139731259202976 --> 139731259203136;
  139731259203296["number: 1949"]:::custom_number;
  139731259202976 --> 139731259203296;

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
    <operator name=":">
        <string quote=":" value="person"/>
        <delimited kind="braces" separator="comma">
            <operator name=":">
                <string quote=":" value="name"/>
                <string quote="," value="Alice"/>
            </operator>
            <operator name=":">
                <string quote=":" value="age"/>
                <number value="25"/>
            </operator>
            <operator name=":">
                <string quote=":" value="isStudent"/>
                <identifier name="true"/>
            </operator>
            <operator name=":">
                <string quote=":" value="skills"/>
                <delimited kind="brackets" separator="comma">
                    <string quote="," value="Python"/>
                    <string quote="," value="JavaScript"/>
                    <string quote="]" value="SQL"/>
                </delimited>
            </operator>
            <operator name=":">
                <string quote=":" value="address"/>
                <delimited kind="braces" separator="comma">
                    <operator name=":">
                        <string quote=":" value="street"/>
                        <string quote="," value="123 Maple Street"/>
                    </operator>
                    <operator name=":">
                        <string quote=":" value="city"/>
                        <string quote="," value="Exampleville"/>
                    </operator>
                    <operator name=":">
                        <string quote=":" value="country"/>
                        <string quote="&quot;" value="Neverland"/>
                    </operator>
                </delimited>
            </operator>
            <operator name=":">
                <string quote=":" value="favoriteBooks"/>
                <delimited kind="brackets" separator="comma">
                    <delimited kind="braces" separator="comma">
                        <operator name=":">
                            <string quote=":" value="title"/>
                            <string quote="," value="To Kill a Mockingbird"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="author"/>
                            <string quote="," value="Harper Lee"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="yearPublished"/>
                            <number value="1960"/>
                        </operator>
                    </delimited>
                    <delimited kind="braces" separator="comma">
                        <operator name=":">
                            <string quote=":" value="title"/>
                            <string quote="," value="1984"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="author"/>
                            <string quote="," value="George Orwell"/>
                        </operator>
                        <operator name=":">
                            <string quote=":" value="yearPublished"/>
                            <number value="1949"/>
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
            "children": [
                {
                    "role": "string",
                    "quote": ":",
                    "value": "person"
                },
                {
                    "role": "delimited",
                    "kind": "braces",
                    "separator": "comma",
                    "children": [
                        {
                            "role": "operator",
                            "name": ":",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "name"
                                },
                                {
                                    "role": "string",
                                    "quote": ",",
                                    "value": "Alice"
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "age"
                                },
                                {
                                    "role": "number",
                                    "value": 25
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
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
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "skills"
                                },
                                {
                                    "role": "delimited",
                                    "kind": "brackets",
                                    "separator": "comma",
                                    "children": [
                                        {
                                            "role": "string",
                                            "quote": ",",
                                            "value": "Python"
                                        },
                                        {
                                            "role": "string",
                                            "quote": ",",
                                            "value": "JavaScript"
                                        },
                                        {
                                            "role": "string",
                                            "quote": "]",
                                            "value": "SQL"
                                        }
                                    ]
                                }
                            ]
                        },
                        {
                            "role": "operator",
                            "name": ":",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
                                    "value": "address"
                                },
                                {
                                    "role": "delimited",
                                    "kind": "braces",
                                    "separator": "comma",
                                    "children": [
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "street"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": ",",
                                                    "value": "123 Maple Street"
                                                }
                                            ]
                                        },
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "city"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": ",",
                                                    "value": "Exampleville"
                                                }
                                            ]
                                        },
                                        {
                                            "role": "operator",
                                            "name": ":",
                                            "children": [
                                                {
                                                    "role": "string",
                                                    "quote": ":",
                                                    "value": "country"
                                                },
                                                {
                                                    "role": "string",
                                                    "quote": "\"",
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
                            "name": ":",
                            "children": [
                                {
                                    "role": "string",
                                    "quote": ":",
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
                                                    "name": ":",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "title"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "To Kill a Mockingbird"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "author"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "Harper Lee"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "yearPublished"
                                                        },
                                                        {
                                                            "role": "number",
                                                            "value": 1960
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
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "title"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "1984"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "author"
                                                        },
                                                        {
                                                            "role": "string",
                                                            "quote": ",",
                                                            "value": "George Orwell"
                                                        }
                                                    ]
                                                },
                                                {
                                                    "role": "operator",
                                                    "name": ":",
                                                    "children": [
                                                        {
                                                            "role": "string",
                                                            "quote": ":",
                                                            "value": "yearPublished"
                                                        },
                                                        {
                                                            "role": "number",
                                                            "value": 1949
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
```

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
  children:
  - role: string
    quote: ':'
    value: person
  - role: delimited
    kind: braces
    separator: comma
    children:
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: name
      - role: string
        quote: ','
        value: Alice
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: age
      - role: number
        value: 25
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: isStudent
      - role: identifier
        name: 'true'
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: skills
      - role: delimited
        kind: brackets
        separator: comma
        children:
        - role: string
          quote: ','
          value: Python
        - role: string
          quote: ','
          value: JavaScript
        - role: string
          quote: ']'
          value: SQL
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
        value: address
      - role: delimited
        kind: braces
        separator: comma
        children:
        - role: operator
          name: ':'
          children:
          - role: string
            quote: ':'
            value: street
          - role: string
            quote: ','
            value: 123 Maple Street
        - role: operator
          name: ':'
          children:
          - role: string
            quote: ':'
            value: city
          - role: string
            quote: ','
            value: Exampleville
        - role: operator
          name: ':'
          children:
          - role: string
            quote: ':'
            value: country
          - role: string
            quote: '"'
            value: Neverland
    - role: operator
      name: ':'
      children:
      - role: string
        quote: ':'
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
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: To Kill a Mockingbird
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: Harper Lee
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: yearPublished
            - role: number
              value: 1960
        - role: delimited
          kind: braces
          separator: comma
          children:
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: title
            - role: string
              quote: ','
              value: '1984'
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: author
            - role: string
              quote: ','
              value: George Orwell
          - role: operator
            name: ':'
            children:
            - role: string
              quote: ':'
              value: yearPublished
            - role: number
              value: 1949

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "136543728665840" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728665920" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728665840" -> "136543728665920";
  "136543728666000" [label="string", shape="box", fillcolor="lightgray"];
  "136543728665920" -> "136543728666000";
  "136543728666080" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728665920" -> "136543728666080";
  "136543728666160" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728666160";
  "136543728666240" [label="string", shape="box", fillcolor="lightgray"];
  "136543728666160" -> "136543728666240";
  "136543728666320" [label="string", shape="box", fillcolor="lightgray"];
  "136543728666160" -> "136543728666320";
  "136543728666400" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728666400";
  "136543728666480" [label="string", shape="box", fillcolor="lightgray"];
  "136543728666400" -> "136543728666480";
  "136543728666560" [label="number: 25", shape="box", fillcolor="lightgoldenrodyellow"];
  "136543728666400" -> "136543728666560";
  "136543728666640" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728666640";
  "136543728666720" [label="string", shape="box", fillcolor="lightgray"];
  "136543728666640" -> "136543728666720";
  "136543728666800" [label="identifier: true", shape="box", fillcolor="Honeydew"];
  "136543728666640" -> "136543728666800";
  "136543728666880" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728666880";
  "136543728666960" [label="string", shape="box", fillcolor="lightgray"];
  "136543728666880" -> "136543728666960";
  "136543728667040" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728666880" -> "136543728667040";
  "136543728667120" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667040" -> "136543728667120";
  "136543728667200" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667040" -> "136543728667200";
  "136543728667280" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667040" -> "136543728667280";
  "136543728667360" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728667360";
  "136543728667440" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667360" -> "136543728667440";
  "136543728667520" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728667360" -> "136543728667520";
  "136543728667600" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728667520" -> "136543728667600";
  "136543728667760" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667600" -> "136543728667760";
  "136543728667920" [label="string", shape="box", fillcolor="lightgray"];
  "136543728667600" -> "136543728667920";
  "136543728668000" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728667520" -> "136543728668000";
  "136543728668160" [label="string", shape="box", fillcolor="lightgray"];
  "136543728668000" -> "136543728668160";
  "136543728668320" [label="string", shape="box", fillcolor="lightgray"];
  "136543728668000" -> "136543728668320";
  "136543728668400" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728667520" -> "136543728668400";
  "136543728668560" [label="string", shape="box", fillcolor="lightgray"];
  "136543728668400" -> "136543728668560";
  "136543728668720" [label="string", shape="box", fillcolor="lightgray"];
  "136543728668400" -> "136543728668720";
  "136543728668800" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728666080" -> "136543728668800";
  "136543728668880" [label="string", shape="box", fillcolor="lightgray"];
  "136543728668800" -> "136543728668880";
  "136543728668960" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728668800" -> "136543728668960";
  "136543728669040" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728668960" -> "136543728669040";
  "136543728669200" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728669040" -> "136543728669200";
  "136543728669360" [label="string", shape="box", fillcolor="lightgray"];
  "136543728669200" -> "136543728669360";
  "136543728669520" [label="string", shape="box", fillcolor="lightgray"];
  "136543728669200" -> "136543728669520";
  "136543728669760" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728669040" -> "136543728669760";
  "136543728669920" [label="string", shape="box", fillcolor="lightgray"];
  "136543728669760" -> "136543728669920";
  "136543728670080" [label="string", shape="box", fillcolor="lightgray"];
  "136543728669760" -> "136543728670080";
  "136543728670320" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728669040" -> "136543728670320";
  "136543728670480" [label="string", shape="box", fillcolor="lightgray"];
  "136543728670320" -> "136543728670480";
  "136543728670640" [label="number: 1960", shape="box", fillcolor="lightgoldenrodyellow"];
  "136543728670320" -> "136543728670640";
  "136543728670800" [label="delimited", shape="box", fillcolor="lightgray"];
  "136543728668960" -> "136543728670800";
  "136543728670960" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728670800" -> "136543728670960";
  "136543728671120" [label="string", shape="box", fillcolor="lightgray"];
  "136543728670960" -> "136543728671120";
  "136543728671280" [label="string", shape="box", fillcolor="lightgray"];
  "136543728670960" -> "136543728671280";
  "136543728671520" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728670800" -> "136543728671520";
  "136543728671680" [label="string", shape="box", fillcolor="lightgray"];
  "136543728671520" -> "136543728671680";
  "136543728671840" [label="string", shape="box", fillcolor="lightgray"];
  "136543728671520" -> "136543728671840";
  "136543728672080" [label="operator: :", shape="box", fillcolor="#C0FFC0"];
  "136543728670800" -> "136543728672080";
  "136543728672240" [label="string", shape="box", fillcolor="lightgray"];
  "136543728672080" -> "136543728672240";
  "136543728672400" [label="number: 1949", shape="box", fillcolor="lightgoldenrodyellow"];
  "136543728672080" -> "136543728672400";
}
```

![Image generated for example](images/factorial.png)
