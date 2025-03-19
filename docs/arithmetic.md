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
  137880009803072["operator: -"]:::custom_operator;
  137880009803152["operator: +"]:::custom_operator;
  137880009803072 --> 137880009803152;
  137880009803232["operator: *"]:::custom_operator;
  137880009803152 --> 137880009803232;
  137880009803312["number: 2"]:::custom_number;
  137880009803232 --> 137880009803312;
  137880009803392["operator: *"]:::custom_operator;
  137880009803232 --> 137880009803392;
  137880009803472["number: 100"]:::custom_number;
  137880009803392 --> 137880009803472;
  137880009803552["number: 100"]:::custom_number;
  137880009803392 --> 137880009803552;
  137880009803632["number: 100"]:::custom_number;
  137880009803152 --> 137880009803632;
  137880009803712["number: 1"]:::custom_number;
  137880009803072 --> 137880009803712;

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
<operator syntax="infix" name="-">
  <operator syntax="infix" name="+">
    <operator syntax="infix" name="*">
      <number value="2" />
      <operator syntax="infix" name="*">
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
  "role": "unit",
  "src": "arithmetic.mg",
  "children": [
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
role: operator
syntax: infix
name: '-'
children:
- role: operator
  syntax: infix
  name: +
  children:
  - role: operator
    syntax: infix
    name: '*'
    children:
    - role: number
      value: 2
    - role: operator
      syntax: infix
      name: '*'
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
  "node_0xc0000a8570" [label="unit: arithmetic.mg", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8510" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8570" -> "node_0xc0000a8510";
  "node_0xc0000a8450" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8510" -> "node_0xc0000a8450";
  "node_0xc0000a8390" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8450" -> "node_0xc0000a8390";
  "node_0xc0000a8210" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8390" -> "node_0xc0000a8210";
  "node_0xc0000a8330" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8390" -> "node_0xc0000a8330";
  "node_0xc0000a8270" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8330" -> "node_0xc0000a8270";
  "node_0xc0000a82d0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8330" -> "node_0xc0000a82d0";
  "node_0xc0000a83f0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8450" -> "node_0xc0000a83f0";
  "node_0xc0000a84b0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8510" -> "node_0xc0000a84b0";
}
```

![Image generated for example](images/arithmetic.png)
