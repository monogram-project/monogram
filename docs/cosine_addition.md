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
  129657822365232["operator: -"]:::custom_operator;
  129657822365312["operator: *"]:::custom_operator;
  129657822365232 --> 129657822365312;
  129657822365392["apply"]:::custom_apply;
  129657822365312 --> 129657822365392;
  129657822365472["identifier: cos"]:::custom_identifier;
  129657822365392 --> 129657822365472;
  129657822365552["arguments"]:::custom_arguments;
  129657822365392 --> 129657822365552;
  129657822365632["identifier: A"]:::custom_identifier;
  129657822365552 --> 129657822365632;
  129657822365712["apply"]:::custom_apply;
  129657822365312 --> 129657822365712;
  129657822365792["identifier: cos"]:::custom_identifier;
  129657822365712 --> 129657822365792;
  129657822365872["arguments"]:::custom_arguments;
  129657822365712 --> 129657822365872;
  129657822365952["identifier: B"]:::custom_identifier;
  129657822365872 --> 129657822365952;
  129657822366032["operator: *"]:::custom_operator;
  129657822365232 --> 129657822366032;
  129657822366112["apply"]:::custom_apply;
  129657822366032 --> 129657822366112;
  129657822366192["identifier: sin"]:::custom_identifier;
  129657822366112 --> 129657822366192;
  129657822366272["arguments"]:::custom_arguments;
  129657822366112 --> 129657822366272;
  129657822366352["identifier: A"]:::custom_identifier;
  129657822366272 --> 129657822366352;
  129657822366432["apply"]:::custom_apply;
  129657822366032 --> 129657822366432;
  129657822366512["identifier: sin"]:::custom_identifier;
  129657822366432 --> 129657822366512;
  129657822366592["arguments"]:::custom_arguments;
  129657822366432 --> 129657822366592;
  129657822366672["identifier: B"]:::custom_identifier;
  129657822366592 --> 129657822366672;

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
  <operator syntax="infix" name="*">
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
  <operator syntax="infix" name="*">
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
  "role": "unit",
  "src": "cosine_addition.mg",
  "children": [
    {
      "role": "operator",
      "syntax": "infix",
      "name": "-",
      "children": [
        {
          "role": "operator",
          "syntax": "infix",
          "name": "*",
          "children": [
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
  name: '*'
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
  syntax: infix
  name: '*'
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
  "node_0xc0000a8870" [label="unit: cosine_addition.mg", shape="box", fillcolor="lightgray"];
  "node_0xc0000a8810" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8870" -> "node_0xc0000a8810";
  "node_0xc0000a84b0" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8810" -> "node_0xc0000a84b0";
  "node_0xc0000a8300" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a84b0" -> "node_0xc0000a8300";
  "node_0xc0000a8210" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8300" -> "node_0xc0000a8210";
  "node_0xc0000a82a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8300" -> "node_0xc0000a82a0";
  "node_0xc0000a8270" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a82a0" -> "node_0xc0000a8270";
  "node_0xc0000a8450" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a84b0" -> "node_0xc0000a8450";
  "node_0xc0000a8360" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8450" -> "node_0xc0000a8360";
  "node_0xc0000a83f0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8450" -> "node_0xc0000a83f0";
  "node_0xc0000a83c0" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a83f0" -> "node_0xc0000a83c0";
  "node_0xc0000a87b0" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8810" -> "node_0xc0000a87b0";
  "node_0xc0000a8600" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a87b0" -> "node_0xc0000a8600";
  "node_0xc0000a8510" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8600" -> "node_0xc0000a8510";
  "node_0xc0000a85a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8600" -> "node_0xc0000a85a0";
  "node_0xc0000a8570" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a85a0" -> "node_0xc0000a8570";
  "node_0xc0000a8750" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a87b0" -> "node_0xc0000a8750";
  "node_0xc0000a8660" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8750" -> "node_0xc0000a8660";
  "node_0xc0000a86f0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8750" -> "node_0xc0000a86f0";
  "node_0xc0000a86c0" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a86f0" -> "node_0xc0000a86c0";
}
```

![Image generated for example](images/cosine_addition.png)
