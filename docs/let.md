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
  136060684305728["form: surround"]:::custom_form;
  136060684305808["part: let"]:::custom_part;
  136060684305728 --> 136060684305808;
  136060684305888["operator: ="]:::custom_operator;
  136060684305808 --> 136060684305888;
  136060684305968["identifier: x"]:::custom_identifier;
  136060684305888 --> 136060684305968;
  136060684306048["apply"]:::custom_apply;
  136060684305888 --> 136060684306048;
  136060684306128["identifier: f"]:::custom_identifier;
  136060684306048 --> 136060684306128;
  136060684306208["arguments"]:::custom_arguments;
  136060684306048 --> 136060684306208;
  136060684306288["identifier: a"]:::custom_identifier;
  136060684306208 --> 136060684306288;
  136060684306368["operator: ="]:::custom_operator;
  136060684305808 --> 136060684306368;
  136060684306448["identifier: y"]:::custom_identifier;
  136060684306368 --> 136060684306448;
  136060684306528["apply"]:::custom_apply;
  136060684306368 --> 136060684306528;
  136060684306608["identifier: g"]:::custom_identifier;
  136060684306528 --> 136060684306608;
  136060684306688["arguments"]:::custom_arguments;
  136060684306528 --> 136060684306688;
  136060684306768["identifier: b"]:::custom_identifier;
  136060684306688 --> 136060684306768;
  136060684306848["part: in"]:::custom_part;
  136060684305728 --> 136060684306848;
  136060684306928["delimited"]:::custom_delimited;
  136060684306848 --> 136060684306928;
  136060684307008["identifier: x"]:::custom_identifier;
  136060684306928 --> 136060684307008;
  136060684307088["identifier: y"]:::custom_identifier;
  136060684306928 --> 136060684307088;

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
          "syntax": "infix",
          "name": "=",
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
          "name": "=",
          "syntax": "infix",
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
    name: '='
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
  "node_0xc0001147e0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc000114630" [label="part: let", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0001147e0" -> "node_0xc000114630";
  "node_0xc0001143c0" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc000114630" -> "node_0xc0001143c0";
  "node_0xc000114210" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc0001143c0" -> "node_0xc000114210";
  "node_0xc000114360" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0001143c0" -> "node_0xc000114360";
  "node_0xc000114270" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc000114360" -> "node_0xc000114270";
  "node_0xc000114300" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc000114360" -> "node_0xc000114300";
  "node_0xc0001142d0" [label="identifier: a", shape="box", fillcolor="Honeydew"];
  "node_0xc000114300" -> "node_0xc0001142d0";
  "node_0xc0001145d0" [label="operator: =", shape="box", fillcolor="#C0FFC0"];
  "node_0xc000114630" -> "node_0xc0001145d0";
  "node_0xc000114420" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc0001145d0" -> "node_0xc000114420";
  "node_0xc000114570" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0001145d0" -> "node_0xc000114570";
  "node_0xc000114480" [label="identifier: g", shape="box", fillcolor="Honeydew"];
  "node_0xc000114570" -> "node_0xc000114480";
  "node_0xc000114510" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc000114570" -> "node_0xc000114510";
  "node_0xc0001144e0" [label="identifier: b", shape="box", fillcolor="Honeydew"];
  "node_0xc000114510" -> "node_0xc0001144e0";
  "node_0xc0001147b0" [label="part: in", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0001147e0" -> "node_0xc0001147b0";
  "node_0xc000114750" [label="delimited", shape="box", fillcolor="lightgray"];
  "node_0xc0001147b0" -> "node_0xc000114750";
  "node_0xc000114690" [label="identifier: x", shape="box", fillcolor="Honeydew"];
  "node_0xc000114750" -> "node_0xc000114690";
  "node_0xc0001146f0" [label="identifier: y", shape="box", fillcolor="Honeydew"];
  "node_0xc000114750" -> "node_0xc0001146f0";
}
```

![Image generated for example](images/let.png)
