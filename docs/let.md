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
  node_0xc0000aec90["form: surround"]:::custom_form;
  node_0xc0000aeae0["part: let"]:::custom_part;
  node_0xc0000aec90 --> node_0xc0000aeae0;
  node_0xc0000ae870["operator: ="]:::custom_operator;
  node_0xc0000aeae0 --> node_0xc0000ae870;
  node_0xc0000ae6c0["identifier: x"]:::custom_identifier;
  node_0xc0000ae870 --> node_0xc0000ae6c0;
  node_0xc0000ae810["apply"]:::custom_apply;
  node_0xc0000ae870 --> node_0xc0000ae810;
  node_0xc0000ae720["identifier: f"]:::custom_identifier;
  node_0xc0000ae810 --> node_0xc0000ae720;
  node_0xc0000ae7b0["arguments"]:::custom_arguments;
  node_0xc0000ae810 --> node_0xc0000ae7b0;
  node_0xc0000ae780["identifier: a"]:::custom_identifier;
  node_0xc0000ae7b0 --> node_0xc0000ae780;
  node_0xc0000aea80["operator: ="]:::custom_operator;
  node_0xc0000aeae0 --> node_0xc0000aea80;
  node_0xc0000ae8d0["identifier: y"]:::custom_identifier;
  node_0xc0000aea80 --> node_0xc0000ae8d0;
  node_0xc0000aea20["apply"]:::custom_apply;
  node_0xc0000aea80 --> node_0xc0000aea20;
  node_0xc0000ae930["identifier: g"]:::custom_identifier;
  node_0xc0000aea20 --> node_0xc0000ae930;
  node_0xc0000ae9c0["arguments"]:::custom_arguments;
  node_0xc0000aea20 --> node_0xc0000ae9c0;
  node_0xc0000ae990["identifier: b"]:::custom_identifier;
  node_0xc0000ae9c0 --> node_0xc0000ae990;
  node_0xc0000aec60["part: in"]:::custom_part;
  node_0xc0000aec90 --> node_0xc0000aec60;
  node_0xc0000aec00["delimited"]:::custom_delimited;
  node_0xc0000aec60 --> node_0xc0000aec00;
  node_0xc0000aeb40["identifier: x"]:::custom_identifier;
  node_0xc0000aec00 --> node_0xc0000aeb40;
  node_0xc0000aeba0["identifier: y"]:::custom_identifier;
  node_0xc0000aec00 --> node_0xc0000aeba0;
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
    <operator name="=" syntax="infix">
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
    <delimited kind="parentheses" separator="comma">
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
  "role": "form",
  "syntax": "surround",
  "children": [
    {
      "role": "part",
      "keyword": "let",
      "children": [
        {
          "role": "operator",
          "name": "=",
          "syntax": "infix",
          "children": [
            {
              "role": "identifier",
              "name": "x"
            },
            {
              "role": "apply",
              "kind": "parentheses",
              "separator": "undefined",
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
              "separator": "undefined",
              "kind": "parentheses",
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
syntax: surround
children:
- role: part
  keyword: let
  children:
  - role: operator
    name: =
    syntax: infix
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
    name: =
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
    kind: parentheses
    separator: comma
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
  "node_0xc00019cc90" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc00019cae0" [label="part: let", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00019cc90" -> "node_0xc00019cae0";
  "node_0xc00019c870" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00019cae0" -> "node_0xc00019c870";
  "node_0xc00019c6c0" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc00019c870" -> "node_0xc00019c6c0";
  "node_0xc00019c810" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00019c870" -> "node_0xc00019c810";
  "node_0xc00019c720" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc00019c810" -> "node_0xc00019c720";
  "node_0xc00019c7b0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00019c810" -> "node_0xc00019c7b0";
  "node_0xc00019c780" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc00019c7b0" -> "node_0xc00019c780";
  "node_0xc00019ca80" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00019cae0" -> "node_0xc00019ca80";
  "node_0xc00019c8d0" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc00019ca80" -> "node_0xc00019c8d0";
  "node_0xc00019ca20" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00019ca80" -> "node_0xc00019ca20";
  "node_0xc00019c930" [label="identifier: g", shape="box", fillcolor="Honeydew"];
  "node_0xc00019ca20" -> "node_0xc00019c930";
  "node_0xc00019c9c0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00019ca20" -> "node_0xc00019c9c0";
  "node_0xc00019c990" [label="identifier: b", shape="box", fillcolor="Honeydew"];
  "node_0xc00019c9c0" -> "node_0xc00019c990";
  "node_0xc00019cc60" [label="part: in", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00019cc90" -> "node_0xc00019cc60";
  "node_0xc00019cc00" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc00019cc60" -> "node_0xc00019cc00";
  "node_0xc00019cb40" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc00019cc00" -> "node_0xc00019cb40";
  "node_0xc00019cba0" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc00019cc00" -> "node_0xc00019cba0";
}
```

![Image generated for example](images/let.png)
