# Cosine addition formula

## Monogram

```txt
cos(A) * cos(B) - sin(A) * sin(B)

```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  node_0xc000122cc0["operator: -"]:::custom_operator;
  node_0xc000122960["operator: *"]:::custom_operator;
  node_0xc000122cc0 --> node_0xc000122960;
  node_0xc0001227b0["apply"]:::custom_apply;
  node_0xc000122960 --> node_0xc0001227b0;
  node_0xc0001226c0["identifier: cos"]:::custom_identifier;
  node_0xc0001227b0 --> node_0xc0001226c0;
  node_0xc000122750["arguments"]:::custom_arguments;
  node_0xc0001227b0 --> node_0xc000122750;
  node_0xc000122720["identifier: A"]:::custom_identifier;
  node_0xc000122750 --> node_0xc000122720;
  node_0xc000122900["apply"]:::custom_apply;
  node_0xc000122960 --> node_0xc000122900;
  node_0xc000122810["identifier: cos"]:::custom_identifier;
  node_0xc000122900 --> node_0xc000122810;
  node_0xc0001228a0["arguments"]:::custom_arguments;
  node_0xc000122900 --> node_0xc0001228a0;
  node_0xc000122870["identifier: B"]:::custom_identifier;
  node_0xc0001228a0 --> node_0xc000122870;
  node_0xc000122c60["operator: *"]:::custom_operator;
  node_0xc000122cc0 --> node_0xc000122c60;
  node_0xc000122ab0["apply"]:::custom_apply;
  node_0xc000122c60 --> node_0xc000122ab0;
  node_0xc0001229c0["identifier: sin"]:::custom_identifier;
  node_0xc000122ab0 --> node_0xc0001229c0;
  node_0xc000122a50["arguments"]:::custom_arguments;
  node_0xc000122ab0 --> node_0xc000122a50;
  node_0xc000122a20["identifier: A"]:::custom_identifier;
  node_0xc000122a50 --> node_0xc000122a20;
  node_0xc000122c00["apply"]:::custom_apply;
  node_0xc000122c60 --> node_0xc000122c00;
  node_0xc000122b10["identifier: sin"]:::custom_identifier;
  node_0xc000122c00 --> node_0xc000122b10;
  node_0xc000122ba0["arguments"]:::custom_arguments;
  node_0xc000122c00 --> node_0xc000122ba0;
  node_0xc000122b70["identifier: B"]:::custom_identifier;
  node_0xc000122ba0 --> node_0xc000122b70;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;
classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
```

## XML

```xml
<operator name="-" syntax="infix">
  <operator name="*" syntax="infix">
    <apply kind="parentheses" separator="undefined">
      <identifier name="cos" />
      <arguments>
        <identifier name="A" />
      </arguments>
    </apply>
    <apply kind="parentheses" separator="undefined">
      <identifier name="cos" />
      <arguments>
        <identifier name="B" />
      </arguments>
    </apply>
  </operator>
  <operator name="*" syntax="infix">
    <apply kind="parentheses" separator="undefined">
      <identifier name="sin" />
      <arguments>
        <identifier name="A" />
      </arguments>
    </apply>
    <apply kind="parentheses" separator="undefined">
      <identifier name="sin" />
      <arguments>
        <identifier name="B" />
      </arguments>
    </apply>
  </operator>
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
  "name": "-",
  "syntax": "infix",
  "children": [
    {
      "role": "operator",
      "syntax": "infix",
      "name": "*",
      "children": [
        {
          "role": "apply",
          "separator": "undefined",
          "kind": "parentheses",
          "children": [
            {
              "role": "identifier",
              "name": "cos"
            },
            {
              "role": "arguments",
              "children": [
                {
                  "role": "identifier",
                  "name": "A"
                }
              ]
            }
          ]
        },
        {
          "role": "apply",
          "kind": "parentheses",
          "separator": "undefined",
          "children": [
            {
              "role": "identifier",
              "name": "cos"
            },
            {
              "role": "arguments",
              "children": [
                {
                  "role": "identifier",
                  "name": "B"
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
      "name": "*",
      "children": [
        {
          "role": "apply",
          "separator": "undefined",
          "kind": "parentheses",
          "children": [
            {
              "role": "identifier",
              "name": "sin"
            },
            {
              "role": "arguments",
              "children": [
                {
                  "role": "identifier",
                  "name": "A"
                }
              ]
            }
          ]
        },
        {
          "role": "apply",
          "kind": "parentheses",
          "separator": "undefined",
          "children": [
            {
              "role": "identifier",
              "name": "sin"
            },
            {
              "role": "arguments",
              "children": [
                {
                  "role": "identifier",
                  "name": "B"
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
name: "-"
syntax: infix
children:
- role: operator
  name: "*"
  syntax: infix
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: cos
    - role: arguments
      children:
      - role: identifier
        name: A
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: cos
    - role: arguments
      children:
      - role: identifier
        name: B
- role: operator
  name: "*"
  syntax: infix
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: sin
    - role: arguments
      children:
      - role: identifier
        name: A
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: sin
    - role: arguments
      children:
      - role: identifier
        name: B
```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc0000aecc0" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000ae960" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000aecc0" -> "node_0xc0000ae960";
  "node_0xc0000ae7b0" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000ae960" -> "node_0xc0000ae7b0";
  "node_0xc0000ae6c0" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae7b0" -> "node_0xc0000ae6c0";
  "node_0xc0000ae750" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000ae7b0" -> "node_0xc0000ae750";
  "node_0xc0000ae720" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae750" -> "node_0xc0000ae720";
  "node_0xc0000ae900" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000ae960" -> "node_0xc0000ae900";
  "node_0xc0000ae810" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae900" -> "node_0xc0000ae810";
  "node_0xc0000ae8a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000ae900" -> "node_0xc0000ae8a0";
  "node_0xc0000ae870" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae8a0" -> "node_0xc0000ae870";
  "node_0xc0000aec60" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000aecc0" -> "node_0xc0000aec60";
  "node_0xc0000aeab0" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000aec60" -> "node_0xc0000aeab0";
  "node_0xc0000ae9c0" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aeab0" -> "node_0xc0000ae9c0";
  "node_0xc0000aea50" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000aeab0" -> "node_0xc0000aea50";
  "node_0xc0000aea20" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aea50" -> "node_0xc0000aea20";
  "node_0xc0000aec00" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000aec60" -> "node_0xc0000aec00";
  "node_0xc0000aeb10" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aec00" -> "node_0xc0000aeb10";
  "node_0xc0000aeba0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000aec00" -> "node_0xc0000aeba0";
  "node_0xc0000aeb70" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aeba0" -> "node_0xc0000aeb70";
}
```

![Image generated for example](images/cosine_addition.png)
