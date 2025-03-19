# Let Expression


## Monogram

This example shows how you might simulate `let` expressions in Monogram:

```txt
let x = f(a)
    y = g(b)
in:
    (x, y)
endlet

```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  124595389695216["form: surround"]:::custom_form;
  124595389695296["part: let"]:::custom_part;
  124595389695216 --> 124595389695296;
  124595389695376["operator: ="]:::custom_operator;
  124595389695296 --> 124595389695376;
  124595389695456["identifier: x"]:::custom_identifier;
  124595389695376 --> 124595389695456;
  124595389695536["apply"]:::custom_apply;
  124595389695376 --> 124595389695536;
  124595389695616["identifier: f"]:::custom_identifier;
  124595389695536 --> 124595389695616;
  124595389695696["arguments"]:::custom_arguments;
  124595389695536 --> 124595389695696;
  124595389695776["identifier: a"]:::custom_identifier;
  124595389695696 --> 124595389695776;
  124595389695856["operator: ="]:::custom_operator;
  124595389695296 --> 124595389695856;
  124595389695936["identifier: y"]:::custom_identifier;
  124595389695856 --> 124595389695936;
  124595389696016["apply"]:::custom_apply;
  124595389695856 --> 124595389696016;
  124595389696096["identifier: g"]:::custom_identifier;
  124595389696016 --> 124595389696096;
  124595389696176["arguments"]:::custom_arguments;
  124595389696016 --> 124595389696176;
  124595389696256["identifier: b"]:::custom_identifier;
  124595389696176 --> 124595389696256;
  124595389696336["part: in"]:::custom_part;
  124595389695216 --> 124595389696336;
  124595389696416["delimited"]:::custom_delimited;
  124595389696336 --> 124595389696416;
  124595389696496["identifier: x"]:::custom_identifier;
  124595389696416 --> 124595389696496;
  124595389696576["identifier: y"]:::custom_identifier;
  124595389696416 --> 124595389696576;

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
<form syntax="surround">
  <part keyword="let">
    <operator syntax="infix" name="=">
      <identifier name="x" />
      <apply kind="parentheses" separator="undefined">
        <identifier name="f" />
        <arguments>
          <identifier name="a" />
        </arguments>
      </apply>
    </operator>
    <operator name="=" syntax="infix">
      <identifier name="y" />
      <apply kind="parentheses" separator="undefined">
        <identifier name="g" />
        <arguments>
          <identifier name="b" />
        </arguments>
      </apply>
    </operator>
  </part>
  <part keyword="in">
    <delimited separator="comma" kind="parentheses">
      <identifier name="x" />
      <identifier name="y" />
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
  "role": "unit",
  "src": "let.mg",
  "children": [
    {
      "role": "form",
      "syntax": "surround",
      "children": [
        {
          "role": "part",
          "keyword": "let",
          "children": [
            {
              "role": "operator",
              "syntax": "infix",
              "name": "=",
              "children": [
                {
                  "role": "identifier",
                  "name": "x"
                },
                {
                  "role": "apply",
                  "separator": "undefined",
                  "kind": "parentheses",
                  "children": [
                    {
                      "role": "identifier",
                      "name": "f"
                    },
                    {
                      "role": "arguments",
                      "children": [
                        {
                          "role": "identifier",
                          "name": "a"
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
              "name": "=",
              "children": [
                {
                  "role": "identifier",
                  "name": "y"
                },
                {
                  "role": "apply",
                  "kind": "parentheses",
                  "separator": "undefined",
                  "children": [
                    {
                      "role": "identifier",
                      "name": "g"
                    },
                    {
                      "role": "arguments",
                      "children": [
                        {
                          "role": "identifier",
                          "name": "b"
                        }
                      ]
                    }
                  ]
                }
              ]
            }
          ]
        },
        {
          "role": "part",
          "keyword": "in",
          "children": [
            {
              "role": "delimited",
              "kind": "parentheses",
              "separator": "comma",
              "children": [
                {
                  "role": "identifier",
                  "name": "x"
                },
                {
                  "role": "identifier",
                  "name": "y"
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
role: form
syntax: surround
children:
- role: part
  keyword: let
  children:
  - role: operator
    syntax: infix
    name: '='
    children:
    - role: identifier
      name: x
    - role: apply
      kind: parentheses
      separator: undefined
      children:
      - role: identifier
        name: f
      - role: arguments
        children:
        - role: identifier
          name: a
  - role: operator
    name: '='
    syntax: infix
    children:
    - role: identifier
      name: y
    - role: apply
      kind: parentheses
      separator: undefined
      children:
      - role: identifier
        name: g
      - role: arguments
        children:
        - role: identifier
          name: b
- role: part
  keyword: in
  children:
  - role: delimited
    separator: comma
    kind: parentheses
    children:
    - role: identifier
      name: x
    - role: identifier
      name: y

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc0000a8870" [label="unit: let.mg", shape="box", fillcolor="lightgray"];
  "node_0xc0000a87e0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc0000a8870" -> "node_0xc0000a87e0";
  "node_0xc0000a8630" [label="part: let", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a87e0" -> "node_0xc0000a8630";
  "node_0xc0000a83c0" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8630" -> "node_0xc0000a83c0";
  "node_0xc0000a8210" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a83c0" -> "node_0xc0000a8210";
  "node_0xc0000a8360" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a83c0" -> "node_0xc0000a8360";
  "node_0xc0000a8270" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8360" -> "node_0xc0000a8270";
  "node_0xc0000a8300" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8360" -> "node_0xc0000a8300";
  "node_0xc0000a82d0" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8300" -> "node_0xc0000a82d0";
  "node_0xc0000a85d0" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a8630" -> "node_0xc0000a85d0";
  "node_0xc0000a8420" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a85d0" -> "node_0xc0000a8420";
  "node_0xc0000a8570" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a85d0" -> "node_0xc0000a8570";
  "node_0xc0000a8480" [label="identifier: g", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8570" -> "node_0xc0000a8480";
  "node_0xc0000a8510" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8570" -> "node_0xc0000a8510";
  "node_0xc0000a84e0" [label="identifier: b", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8510" -> "node_0xc0000a84e0";
  "node_0xc0000a87b0" [label="part: in", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a87e0" -> "node_0xc0000a87b0";
  "node_0xc0000a8750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0000a87b0" -> "node_0xc0000a8750";
  "node_0xc0000a8690" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8750" -> "node_0xc0000a8690";
  "node_0xc0000a86f0" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8750" -> "node_0xc0000a86f0";
}
```

![Image generated for example](images/let.png)
