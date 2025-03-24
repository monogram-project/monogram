# Recursive factorial function

## Monogram

```txt
# The factorial function in monogram.
def f(n):
    if n <= 1:
        1
    else:
        n * f(n - 1)
    endif
enddef
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  node_0xc0000aee70["form: surround"]:::custom_form;
  node_0xc0000ae810["part: def"]:::custom_part;
  node_0xc0000aee70 --> node_0xc0000ae810;
  node_0xc0000ae7b0["apply"]:::custom_apply;
  node_0xc0000ae810 --> node_0xc0000ae7b0;
  node_0xc0000ae6c0["identifier: f"]:::custom_identifier;
  node_0xc0000ae7b0 --> node_0xc0000ae6c0;
  node_0xc0000ae750["arguments"]:::custom_arguments;
  node_0xc0000ae7b0 --> node_0xc0000ae750;
  node_0xc0000ae720["identifier: n"]:::custom_identifier;
  node_0xc0000ae750 --> node_0xc0000ae720;
  node_0xc0000aee40["part: _"]:::custom_part;
  node_0xc0000aee70 --> node_0xc0000aee40;
  node_0xc0000aedb0["form: surround"]:::custom_form;
  node_0xc0000aee40 --> node_0xc0000aedb0;
  node_0xc0000ae990["part: if"]:::custom_part;
  node_0xc0000aedb0 --> node_0xc0000ae990;
  node_0xc0000ae930["operator: <="]:::custom_operator;
  node_0xc0000ae990 --> node_0xc0000ae930;
  node_0xc0000ae870["identifier: n"]:::custom_identifier;
  node_0xc0000ae930 --> node_0xc0000ae870;
  node_0xc0000ae8d0["number: 1"]:::custom_number;
  node_0xc0000ae930 --> node_0xc0000ae8d0;
  node_0xc0000aea50["part: _"]:::custom_part;
  node_0xc0000aedb0 --> node_0xc0000aea50;
  node_0xc0000ae9f0["number: 1"]:::custom_number;
  node_0xc0000aea50 --> node_0xc0000ae9f0;
  node_0xc0000aed80["part: else"]:::custom_part;
  node_0xc0000aedb0 --> node_0xc0000aed80;
  node_0xc0000aed20["operator: *"]:::custom_operator;
  node_0xc0000aed80 --> node_0xc0000aed20;
  node_0xc0000aeab0["identifier: n"]:::custom_identifier;
  node_0xc0000aed20 --> node_0xc0000aeab0;
  node_0xc0000aecc0["apply"]:::custom_apply;
  node_0xc0000aed20 --> node_0xc0000aecc0;
  node_0xc0000aeb10["identifier: f"]:::custom_identifier;
  node_0xc0000aecc0 --> node_0xc0000aeb10;
  node_0xc0000aec60["arguments"]:::custom_arguments;
  node_0xc0000aecc0 --> node_0xc0000aec60;
  node_0xc0000aec30["operator: -"]:::custom_operator;
  node_0xc0000aec60 --> node_0xc0000aec30;
  node_0xc0000aeb70["identifier: n"]:::custom_identifier;
  node_0xc0000aec30 --> node_0xc0000aeb70;
  node_0xc0000aebd0["number: 1"]:::custom_number;
  node_0xc0000aec30 --> node_0xc0000aebd0;
classDef custom_apply fill:lightgreen,stroke:#333,stroke-width:2px;
classDef custom_identifier fill:Honeydew,stroke:#333,stroke-width:2px;
classDef custom_arguments fill:PaleTurquoise,stroke:#333,stroke-width:2px;
classDef custom_operator fill:#C0FFC0,stroke:#333,stroke-width:2px;
classDef custom_number fill:lightgoldenrodyellow,stroke:#333,stroke-width:2px;
classDef custom_form fill:lightpink,stroke:#333,stroke-width:2px;
classDef custom_part fill:#FFD8E1,stroke:#333,stroke-width:2px;
```

## XML

```xml
<form syntax="surround">
  <part keyword="def">
    <apply kind="parentheses" separator="undefined">
      <identifier name="f" />
      <arguments>
        <identifier name="n" />
      </arguments>
    </apply>
  </part>
  <part keyword="_">
    <form syntax="surround">
      <part keyword="if">
        <operator name="&lt;=" syntax="infix">
          <identifier name="n" />
          <number value="1" />
        </operator>
      </part>
      <part keyword="_">
        <number value="1" />
      </part>
      <part keyword="else">
        <operator name="*" syntax="infix">
          <identifier name="n" />
          <apply kind="parentheses" separator="undefined">
            <identifier name="f" />
            <arguments>
              <operator name="-" syntax="infix">
                <identifier name="n" />
                <number value="1" />
              </operator>
            </arguments>
          </apply>
        </operator>
      </part>
    </form>
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
      "keyword": "def",
      "children": [
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
                  "name": "n"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "role": "part",
      "keyword": "_",
      "children": [
        {
          "role": "form",
          "syntax": "surround",
          "children": [
            {
              "role": "part",
              "keyword": "if",
              "children": [
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": "<=",
                  "children": [
                    {
                      "role": "identifier",
                      "name": "n"
                    },
                    {
                      "role": "number",
                      "value": "1"
                    }
                  ]
                }
              ]
            },
            {
              "role": "part",
              "keyword": "_",
              "children": [
                {
                  "role": "number",
                  "value": "1"
                }
              ]
            },
            {
              "role": "part",
              "keyword": "else",
              "children": [
                {
                  "role": "operator",
                  "syntax": "infix",
                  "name": "*",
                  "children": [
                    {
                      "role": "identifier",
                      "name": "n"
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
                              "role": "operator",
                              "syntax": "infix",
                              "name": "-",
                              "children": [
                                {
                                  "role": "identifier",
                                  "name": "n"
                                },
                                {
                                  "role": "number",
                                  "value": "1"
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
  keyword: def
  children:
  - role: apply
    kind: parentheses
    separator: undefined
    children:
    - role: identifier
      name: f
    - role: arguments
      children:
      - role: identifier
        name: n
- role: part
  keyword: _
  children:
  - role: form
    syntax: surround
    children:
    - role: part
      keyword: if
      children:
      - role: operator
        name: <=
        syntax: infix
        children:
        - role: identifier
          name: n
        - role: number
          value: 1
    - role: part
      keyword: _
      children:
      - role: number
        value: 1
    - role: part
      keyword: else
      children:
      - role: operator
        name: "*"
        syntax: infix
        children:
        - role: identifier
          name: n
        - role: apply
          kind: parentheses
          separator: undefined
          children:
          - role: identifier
            name: f
          - role: arguments
            children:
            - role: operator
              name: "-"
              syntax: infix
              children:
              - role: identifier
                name: n
              - role: number
                value: 1
```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "node_0xc0000aee70" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc0000ae810" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000aee70" -> "node_0xc0000ae810";
  "node_0xc0000ae7b0" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000ae810" -> "node_0xc0000ae7b0";
  "node_0xc0000ae6c0" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae7b0" -> "node_0xc0000ae6c0";
  "node_0xc0000ae750" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000ae7b0" -> "node_0xc0000ae750";
  "node_0xc0000ae720" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae750" -> "node_0xc0000ae720";
  "node_0xc0000aee40" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000aee70" -> "node_0xc0000aee40";
  "node_0xc0000aedb0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc0000aee40" -> "node_0xc0000aedb0";
  "node_0xc0000ae990" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000aedb0" -> "node_0xc0000ae990";
  "node_0xc0000ae930" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000ae990" -> "node_0xc0000ae930";
  "node_0xc0000ae870" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000ae930" -> "node_0xc0000ae870";
  "node_0xc0000ae8d0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000ae930" -> "node_0xc0000ae8d0";
  "node_0xc0000aea50" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000aedb0" -> "node_0xc0000aea50";
  "node_0xc0000ae9f0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000aea50" -> "node_0xc0000ae9f0";
  "node_0xc0000aed80" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000aedb0" -> "node_0xc0000aed80";
  "node_0xc0000aed20" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000aed80" -> "node_0xc0000aed20";
  "node_0xc0000aeab0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aed20" -> "node_0xc0000aeab0";
  "node_0xc0000aecc0" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000aed20" -> "node_0xc0000aecc0";
  "node_0xc0000aeb10" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aecc0" -> "node_0xc0000aeb10";
  "node_0xc0000aec60" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000aecc0" -> "node_0xc0000aec60";
  "node_0xc0000aec30" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000aec60" -> "node_0xc0000aec30";
  "node_0xc0000aeb70" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000aec30" -> "node_0xc0000aeb70";
  "node_0xc0000aebd0" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000aec30" -> "node_0xc0000aebd0";
}
```

![Image generated for example](images/factorial.png)
