# C-style conditional

## Monogram

```txt
if! (x) {
  a += 1
} else-if (y) {
  a += 2
} else: {
  a -= 1
}```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  node_0xc00018c630["form: prefix"]:::custom_form;
  node_0xc00018c690["part: if"]:::custom_part;
  node_0xc00018c630 --> node_0xc00018c690;
  node_0xc00018c750["delimited"]:::custom_delimited;
  node_0xc00018c690 --> node_0xc00018c750;
  node_0xc00018c6f0["identifier: x"]:::custom_identifier;
  node_0xc00018c750 --> node_0xc00018c6f0;
  node_0xc00018c930["part: _"]:::custom_part;
  node_0xc00018c630 --> node_0xc00018c930;
  node_0xc00018c8d0["delimited"]:::custom_delimited;
  node_0xc00018c930 --> node_0xc00018c8d0;
  node_0xc00018c870["operator: +="]:::custom_operator;
  node_0xc00018c8d0 --> node_0xc00018c870;
  node_0xc00018c7b0["identifier: a"]:::custom_identifier;
  node_0xc00018c870 --> node_0xc00018c7b0;
  node_0xc00018c810["number: 1"]:::custom_number;
  node_0xc00018c870 --> node_0xc00018c810;
  node_0xc00018c990["part: else-if"]:::custom_part;
  node_0xc00018c630 --> node_0xc00018c990;
  node_0xc00018ca50["delimited"]:::custom_delimited;
  node_0xc00018c990 --> node_0xc00018ca50;
  node_0xc00018c9f0["identifier: y"]:::custom_identifier;
  node_0xc00018ca50 --> node_0xc00018c9f0;
  node_0xc00018cc30["part: _"]:::custom_part;
  node_0xc00018c630 --> node_0xc00018cc30;
  node_0xc00018cbd0["delimited"]:::custom_delimited;
  node_0xc00018cc30 --> node_0xc00018cbd0;
  node_0xc00018cb70["operator: +="]:::custom_operator;
  node_0xc00018cbd0 --> node_0xc00018cb70;
  node_0xc00018cab0["identifier: a"]:::custom_identifier;
  node_0xc00018cb70 --> node_0xc00018cab0;
  node_0xc00018cb10["number: 2"]:::custom_number;
  node_0xc00018cb70 --> node_0xc00018cb10;
  node_0xc00018cc90["part: else"]:::custom_part;
  node_0xc00018c630 --> node_0xc00018cc90;
  node_0xc00018ce10["delimited"]:::custom_delimited;
  node_0xc00018cc90 --> node_0xc00018ce10;
  node_0xc00018cdb0["operator: -="]:::custom_operator;
  node_0xc00018ce10 --> node_0xc00018cdb0;
  node_0xc00018ccf0["identifier: a"]:::custom_identifier;
  node_0xc00018cdb0 --> node_0xc00018ccf0;
  node_0xc00018cd50["number: 1"]:::custom_number;
  node_0xc00018cdb0 --> node_0xc00018cd50;
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
<form syntax="prefix">
  <part keyword="if">
    <delimited kind="parentheses" separator="undefined">
      <identifier name="x" />
    </delimited>
  </part>
  <part keyword="_">
    <delimited kind="braces" separator="undefined">
      <operator name="+=" syntax="infix">
        <identifier name="a" />
        <number value="1" />
      </operator>
    </delimited>
  </part>
  <part keyword="else-if">
    <delimited kind="parentheses" separator="undefined">
      <identifier name="y" />
    </delimited>
  </part>
  <part keyword="_">
    <delimited kind="braces" separator="undefined">
      <operator name="+=" syntax="infix">
        <identifier name="a" />
        <number value="2" />
      </operator>
    </delimited>
  </part>
  <part keyword="else">
    <delimited kind="braces" separator="undefined">
      <operator name="-=" syntax="infix">
        <identifier name="a" />
        <number value="1" />
      </operator>
    </delimited>
  </part>
</form>
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
  "role": "form",
  "syntax": "prefix",
  "children": [
    {
      "role": "part",
      "keyword": "if",
      "children": [
        {
          "role": "delimited",
          "kind": "parentheses",
          "separator": "undefined",
          "children": [
            {
              "role": "identifier",
              "name": "x"
            }
          ]
        }
      ]
    },
    {
      "role": "part",
      "keyword": "_",
      "children": [
        {
          "role": "delimited",
          "kind": "braces",
          "separator": "undefined",
          "children": [
            {
              "role": "operator",
              "syntax": "infix",
              "name": "+=",
              "children": [
                {
                  "role": "identifier",
                  "name": "a"
                },
                {
                  "role": "number",
                  "value": "1"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "role": "part",
      "keyword": "else-if",
      "children": [
        {
          "role": "delimited",
          "kind": "parentheses",
          "separator": "undefined",
          "children": [
            {
              "role": "identifier",
              "name": "y"
            }
          ]
        }
      ]
    },
    {
      "role": "part",
      "keyword": "_",
      "children": [
        {
          "role": "delimited",
          "kind": "braces",
          "separator": "undefined",
          "children": [
            {
              "role": "operator",
              "syntax": "infix",
              "name": "+=",
              "children": [
                {
                  "role": "identifier",
                  "name": "a"
                },
                {
                  "role": "number",
                  "value": "2"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "role": "part",
      "keyword": "else",
      "children": [
        {
          "role": "delimited",
          "kind": "braces",
          "separator": "undefined",
          "children": [
            {
              "role": "operator",
              "syntax": "infix",
              "name": "-=",
              "children": [
                {
                  "role": "identifier",
                  "name": "a"
                },
                {
                  "role": "number",
                  "value": "1"
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
syntax: prefix
children:
- role: part
  keyword: if
  children:
  - role: delimited
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: x
- role: part
  keyword: _
  children:
  - role: delimited
    kind: braces
    separator: undefined
    children:
    - role: operator
      name: +=
      syntax: infix
      children:
      - role: identifier
        name: a
      - role: number
        value: 1
- role: part
  keyword: "else-if"
  children:
  - role: delimited
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: y
- role: part
  keyword: _
  children:
  - role: delimited
    kind: braces
    separator: undefined
    children:
    - role: operator
      name: +=
      syntax: infix
      children:
      - role: identifier
        name: a
      - role: number
        value: 2
- role: part
  keyword: else
  children:
  - role: delimited
    kind: braces
    separator: undefined
    children:
    - role: operator
      name: "-="
      syntax: infix
      children:
      - role: identifier
        name: a
      - role: number
        value: 1
```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc0000ee630" [label="form: prefix", shape="box", fillcolor="lightpink"];
  "node_0xc0000ee690" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000ee630" -> "node_0xc0000ee690";
  "node_0xc0000ee750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000ee690" -> "node_0xc0000ee750";
  "node_0xc0000ee6f0" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ee750" -> "node_0xc0000ee6f0";
  "node_0xc0000ee930" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000ee630" -> "node_0xc0000ee930";
  "node_0xc0000ee8d0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000ee930" -> "node_0xc0000ee8d0";
  "node_0xc0000ee870" [label="operator: +=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000ee8d0" -> "node_0xc0000ee870";
  "node_0xc0000ee7b0" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ee870" -> "node_0xc0000ee7b0";
  "node_0xc0000ee810" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000ee870" -> "node_0xc0000ee810";
  "node_0xc0000ee990" [label="part: else-if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000ee630" -> "node_0xc0000ee990";
  "node_0xc0000eea50" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000ee990" -> "node_0xc0000eea50";
  "node_0xc0000ee9f0" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc0000eea50" -> "node_0xc0000ee9f0";
  "node_0xc0000eec30" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000ee630" -> "node_0xc0000eec30";
  "node_0xc0000eebd0" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000eec30" -> "node_0xc0000eebd0";
  "node_0xc0000eeb70" [label="operator: +=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000eebd0" -> "node_0xc0000eeb70";
  "node_0xc0000eeab0" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc0000eeb70" -> "node_0xc0000eeab0";
  "node_0xc0000eeb10" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000eeb70" -> "node_0xc0000eeb10";
  "node_0xc0000eec90" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000ee630" -> "node_0xc0000eec90";
  "node_0xc0000eee10" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000eec90" -> "node_0xc0000eee10";
  "node_0xc0000eedb0" [label="operator: -=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000eee10" -> "node_0xc0000eedb0";
  "node_0xc0000eecf0" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc0000eedb0" -> "node_0xc0000eecf0";
  "node_0xc0000eed50" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000eedb0" -> "node_0xc0000eed50";
}
```

![Image generated for example](images/cstyle_if.png)
