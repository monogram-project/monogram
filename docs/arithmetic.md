# Simple arithmetic expression

## Monogram

```txt
2 * 100 * 100 + 100 - 1
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  node_0xc0001229c0["operator: -"]:::custom_operator;
  node_0xc000122900["operator: +"]:::custom_operator;
  node_0xc0001229c0 --> node_0xc000122900;
  node_0xc000122840["operator: *"]:::custom_operator;
  node_0xc000122900 --> node_0xc000122840;
  node_0xc0001226c0["number: 2"]:::custom_number;
  node_0xc000122840 --> node_0xc0001226c0;
  node_0xc0001227e0["operator: *"]:::custom_operator;
  node_0xc000122840 --> node_0xc0001227e0;
  node_0xc000122720["number: 100"]:::custom_number;
  node_0xc0001227e0 --> node_0xc000122720;
  node_0xc000122780["number: 100"]:::custom_number;
  node_0xc0001227e0 --> node_0xc000122780;
  node_0xc0001228a0["number: 100"]:::custom_number;
  node_0xc000122900 --> node_0xc0001228a0;
  node_0xc000122960["number: 1"]:::custom_number;
  node_0xc0001229c0 --> node_0xc000122960;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;
classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
```

## XML

```xml
<operator name="-" syntax="infix">
  <operator name="+" syntax="infix">
    <operator name="*" syntax="infix">
      <number value="2" />
      <operator name="*" syntax="infix">
        <number value="100" />
        <number value="100" />
      </operator>
    </operator>
    <number value="100" />
  </operator>
  <number value="1" />
</operator>
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
  "role": "operator",
  "syntax": "infix",
  "name": "-",
  "children": [
    {
      "role": "operator",
      "syntax": "infix",
      "name": "+",
      "children": [
        {
          "role": "operator",
          "syntax": "infix",
          "name": "*",
          "children": [
            {
              "role": "number",
              "value": "2"
            },
            {
              "role": "operator",
              "syntax": "infix",
              "name": "*",
              "children": [
                {
                  "role": "number",
                  "value": "100"
                },
                {
                  "role": "number",
                  "value": "100"
                }
              ]
            }
          ]
        },
        {
          "role": "number",
          "value": "100"
        }
      ]
    },
    {
      "role": "number",
      "value": "1"
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
name: "-"
syntax: infix
children:
- role: operator
  name: +
  syntax: infix
  children:
  - role: operator
    name: "*"
    syntax: infix
    children:
    - role: number
      value: 2
    - role: operator
      name: "*"
      syntax: infix
      children:
      - role: number
        value: 100
      - role: number
        value: 100
  - role: number
    value: 100
- role: number
  value: 1
```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc0001229c0" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc000122900" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0001229c0" -> "node_0xc000122900";
  "node_0xc000122840" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc000122900" -> "node_0xc000122840";
  "node_0xc0001226c0" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc000122840" -> "node_0xc0001226c0";
  "node_0xc0001227e0" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc000122840" -> "node_0xc0001227e0";
  "node_0xc000122720" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0001227e0" -> "node_0xc000122720";
  "node_0xc000122780" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0001227e0" -> "node_0xc000122780";
  "node_0xc0001228a0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc000122900" -> "node_0xc0001228a0";
  "node_0xc000122960" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0001229c0" -> "node_0xc000122960";
}
```

![Image generated for example](images/arithmetic.png)
