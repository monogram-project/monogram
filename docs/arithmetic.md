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
  125413982569792["operator: -"]:::custom_operator;
  125413982569872["operator: +"]:::custom_operator;
  125413982569792 --> 125413982569872;
  125413982569952["operator: *"]:::custom_operator;
  125413982569872 --> 125413982569952;
  125413982570032["number: 2"]:::custom_number;
  125413982569952 --> 125413982570032;
  125413982570112["operator: *"]:::custom_operator;
  125413982569952 --> 125413982570112;
  125413982570192["number: 100"]:::custom_number;
  125413982570112 --> 125413982570192;
  125413982570272["number: 100"]:::custom_number;
  125413982570112 --> 125413982570272;
  125413982570352["number: 100"]:::custom_number;
  125413982569872 --> 125413982570352;
  125413982570432["number: 1"]:::custom_number;
  125413982569792 --> 125413982570432;

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
  "node_0xc00007c510" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c450" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c510" -> "node_0xc00007c450";
  "node_0xc00007c390" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c450" -> "node_0xc00007c390";
  "node_0xc00007c210" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c390" -> "node_0xc00007c210";
  "node_0xc00007c330" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c390" -> "node_0xc00007c330";
  "node_0xc00007c270" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c330" -> "node_0xc00007c270";
  "node_0xc00007c2d0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c330" -> "node_0xc00007c2d0";
  "node_0xc00007c3f0" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c450" -> "node_0xc00007c3f0";
  "node_0xc00007c4b0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c510" -> "node_0xc00007c4b0";
}
```

![Image generated for example](images/arithmetic.png)
