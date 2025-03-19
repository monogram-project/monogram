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
  124855315957056["operator: -"]:::custom_operator;
  124855315957136["operator: +"]:::custom_operator;
  124855315957056 --> 124855315957136;
  124855315957216["operator: *"]:::custom_operator;
  124855315957136 --> 124855315957216;
  124855315957296["number: 2"]:::custom_number;
  124855315957216 --> 124855315957296;
  124855315957376["operator: *"]:::custom_operator;
  124855315957216 --> 124855315957376;
  124855315957456["number: 100"]:::custom_number;
  124855315957376 --> 124855315957456;
  124855315957536["number: 100"]:::custom_number;
  124855315957376 --> 124855315957536;
  124855315957616["number: 100"]:::custom_number;
  124855315957136 --> 124855315957616;
  124855315957696["number: 1"]:::custom_number;
  124855315957056 --> 124855315957696;

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
name: '-'
syntax: infix
children:
- role: operator
  name: +
  syntax: infix
  children:
  - role: operator
    name: '*'
    syntax: infix
    children:
    - role: number
      value: 2
    - role: operator
      name: '*'
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
  "node_0xc00007e570" [label="unit: arithmetic.mg", shape="box", fillcolor="lightgray"];
  "node_0xc00007e510" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e570" -> "node_0xc00007e510";
  "node_0xc00007e450" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e510" -> "node_0xc00007e450";
  "node_0xc00007e390" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e450" -> "node_0xc00007e390";
  "node_0xc00007e210" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e390" -> "node_0xc00007e210";
  "node_0xc00007e330" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e390" -> "node_0xc00007e330";
  "node_0xc00007e270" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e330" -> "node_0xc00007e270";
  "node_0xc00007e2d0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e330" -> "node_0xc00007e2d0";
  "node_0xc00007e3f0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e450" -> "node_0xc00007e3f0";
  "node_0xc00007e4b0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e510" -> "node_0xc00007e4b0";
}
```

![Image generated for example](images/arithmetic.png)
