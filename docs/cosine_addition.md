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
  137412731831856["operator: -"]:::custom_operator;
  137412731831936["operator: *"]:::custom_operator;
  137412731831856 --> 137412731831936;
  137412731832016["apply"]:::custom_apply;
  137412731831936 --> 137412731832016;
  137412731832096["identifier: cos"]:::custom_identifier;
  137412731832016 --> 137412731832096;
  137412731832176["arguments"]:::custom_arguments;
  137412731832016 --> 137412731832176;
  137412731832256["identifier: A"]:::custom_identifier;
  137412731832176 --> 137412731832256;
  137412731832336["apply"]:::custom_apply;
  137412731831936 --> 137412731832336;
  137412731832416["identifier: cos"]:::custom_identifier;
  137412731832336 --> 137412731832416;
  137412731832496["arguments"]:::custom_arguments;
  137412731832336 --> 137412731832496;
  137412731832576["identifier: B"]:::custom_identifier;
  137412731832496 --> 137412731832576;
  137412731832656["operator: *"]:::custom_operator;
  137412731831856 --> 137412731832656;
  137412731832736["apply"]:::custom_apply;
  137412731832656 --> 137412731832736;
  137412731832816["identifier: sin"]:::custom_identifier;
  137412731832736 --> 137412731832816;
  137412731832896["arguments"]:::custom_arguments;
  137412731832736 --> 137412731832896;
  137412731832976["identifier: A"]:::custom_identifier;
  137412731832896 --> 137412731832976;
  137412731833056["apply"]:::custom_apply;
  137412731832656 --> 137412731833056;
  137412731833136["identifier: sin"]:::custom_identifier;
  137412731833056 --> 137412731833136;
  137412731833216["arguments"]:::custom_arguments;
  137412731833056 --> 137412731833216;
  137412731833296["identifier: B"]:::custom_identifier;
  137412731833216 --> 137412731833296;

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
name: '-'
syntax: infix
children:
- role: operator
  name: '*'
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
  name: '*'
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
  "node_0xc00007e870" [label="unit: cosine_addition.mg", shape="box", fillcolor="lightgray"];
  "node_0xc00007e810" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e870" -> "node_0xc00007e810";
  "node_0xc00007e4b0" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e810" -> "node_0xc00007e4b0";
  "node_0xc00007e300" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e4b0" -> "node_0xc00007e300";
  "node_0xc00007e210" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e300" -> "node_0xc00007e210";
  "node_0xc00007e2a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e300" -> "node_0xc00007e2a0";
  "node_0xc00007e270" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e2a0" -> "node_0xc00007e270";
  "node_0xc00007e450" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e4b0" -> "node_0xc00007e450";
  "node_0xc00007e360" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e450" -> "node_0xc00007e360";
  "node_0xc00007e3f0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e450" -> "node_0xc00007e3f0";
  "node_0xc00007e3c0" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e3f0" -> "node_0xc00007e3c0";
  "node_0xc00007e7b0" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e810" -> "node_0xc00007e7b0";
  "node_0xc00007e600" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e7b0" -> "node_0xc00007e600";
  "node_0xc00007e510" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e600" -> "node_0xc00007e510";
  "node_0xc00007e5a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e600" -> "node_0xc00007e5a0";
  "node_0xc00007e570" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e5a0" -> "node_0xc00007e570";
  "node_0xc00007e750" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e7b0" -> "node_0xc00007e750";
  "node_0xc00007e660" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e750" -> "node_0xc00007e660";
  "node_0xc00007e6f0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e750" -> "node_0xc00007e6f0";
  "node_0xc00007e6c0" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e6f0" -> "node_0xc00007e6c0";
}
```

![Image generated for example](images/cosine_addition.png)
