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
  140419325330752["form: surround"]:::custom_form;
  140419325330832["part: def"]:::custom_part;
  140419325330752 --> 140419325330832;
  140419325330912["apply"]:::custom_apply;
  140419325330832 --> 140419325330912;
  140419325330992["identifier: f"]:::custom_identifier;
  140419325330912 --> 140419325330992;
  140419325331072["arguments"]:::custom_arguments;
  140419325330912 --> 140419325331072;
  140419325331152["identifier: n"]:::custom_identifier;
  140419325331072 --> 140419325331152;
  140419325331232["part: _"]:::custom_part;
  140419325330752 --> 140419325331232;
  140419325331312["form: surround"]:::custom_form;
  140419325331232 --> 140419325331312;
  140419325331392["part: if"]:::custom_part;
  140419325331312 --> 140419325331392;
  140419325331472["operator: <="]:::custom_operator;
  140419325331392 --> 140419325331472;
  140419325331552["identifier: n"]:::custom_identifier;
  140419325331472 --> 140419325331552;
  140419325331632["number: 1"]:::custom_number;
  140419325331472 --> 140419325331632;
  140419325331712["part: _"]:::custom_part;
  140419325331312 --> 140419325331712;
  140419325331792["number: 1"]:::custom_number;
  140419325331712 --> 140419325331792;
  140419325331872["part: else"]:::custom_part;
  140419325331312 --> 140419325331872;
  140419325331952["operator: *"]:::custom_operator;
  140419325331872 --> 140419325331952;
  140419325332032["identifier: n"]:::custom_identifier;
  140419325331952 --> 140419325332032;
  140419325332112["apply"]:::custom_apply;
  140419325331952 --> 140419325332112;
  140419325332192["identifier: f"]:::custom_identifier;
  140419325332112 --> 140419325332192;
  140419325332272["arguments"]:::custom_arguments;
  140419325332112 --> 140419325332272;
  140419325332352["operator: -"]:::custom_operator;
  140419325332272 --> 140419325332352;
  140419325332432["identifier: n"]:::custom_identifier;
  140419325332352 --> 140419325332432;
  140419325332512["number: 1"]:::custom_number;
  140419325332352 --> 140419325332512;

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
        <operator syntax="infix" name="&lt;=">
          <identifier name="n" />
          <number value="1" />
        </operator>
      </part>
      <part keyword="_">
        <number value="1" />
      </part>
      <part keyword="else">
        <operator syntax="infix" name="*">
          <identifier name="n" />
          <apply kind="parentheses" separator="undefined">
            <identifier name="f" />
            <arguments>
              <operator syntax="infix" name="-">
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
  "role": "unit",
  "src": "factorial.mg",
  "children": [
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
        syntax: infix
        name: <=
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
        syntax: infix
        name: '*'
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
              syntax: infix
              name: '-'
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
  "node_0xc0000a8a50" [label="unit: factorial.mg", shape="box", fillcolor="lightgray"];
  "node_0xc0000a89c0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc0000a8a50" -> "node_0xc0000a89c0";
  "node_0xc0000a8360" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a89c0" -> "node_0xc0000a8360";
  "node_0xc0000a8300" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a8360" -> "node_0xc0000a8300";
  "node_0xc0000a8210" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8300" -> "node_0xc0000a8210";
  "node_0xc0000a82a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8300" -> "node_0xc0000a82a0";
  "node_0xc0000a8270" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a82a0" -> "node_0xc0000a8270";
  "node_0xc0000a8990" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a89c0" -> "node_0xc0000a8990";
  "node_0xc0000a8900" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc0000a8990" -> "node_0xc0000a8900";
  "node_0xc0000a84e0" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a8900" -> "node_0xc0000a84e0";
  "node_0xc0000a8480" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a84e0" -> "node_0xc0000a8480";
  "node_0xc0000a83c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8480" -> "node_0xc0000a83c0";
  "node_0xc0000a8420" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8480" -> "node_0xc0000a8420";
  "node_0xc0000a85a0" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a8900" -> "node_0xc0000a85a0";
  "node_0xc0000a8540" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a85a0" -> "node_0xc0000a8540";
  "node_0xc0000a88d0" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "node_0xc0000a8900" -> "node_0xc0000a88d0";
  "node_0xc0000a8870" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a88d0" -> "node_0xc0000a8870";
  "node_0xc0000a8600" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8870" -> "node_0xc0000a8600";
  "node_0xc0000a8810" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc0000a8870" -> "node_0xc0000a8810";
  "node_0xc0000a8660" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8810" -> "node_0xc0000a8660";
  "node_0xc0000a87b0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc0000a8810" -> "node_0xc0000a87b0";
  "node_0xc0000a8780" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc0000a87b0" -> "node_0xc0000a8780";
  "node_0xc0000a86c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc0000a8780" -> "node_0xc0000a86c0";
  "node_0xc0000a8720" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc0000a8780" -> "node_0xc0000a8720";
}
```

![Image generated for example](images/factorial.png)
