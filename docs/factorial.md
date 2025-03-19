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
  130643014042944["form: surround"]:::custom_form;
  130643014043024["part: def"]:::custom_part;
  130643014042944 --> 130643014043024;
  130643014043104["apply"]:::custom_apply;
  130643014043024 --> 130643014043104;
  130643014043184["identifier: f"]:::custom_identifier;
  130643014043104 --> 130643014043184;
  130643014043264["arguments"]:::custom_arguments;
  130643014043104 --> 130643014043264;
  130643014043344["identifier: n"]:::custom_identifier;
  130643014043264 --> 130643014043344;
  130643014043424["part: _"]:::custom_part;
  130643014042944 --> 130643014043424;
  130643014043504["form: surround"]:::custom_form;
  130643014043424 --> 130643014043504;
  130643014043584["part: if"]:::custom_part;
  130643014043504 --> 130643014043584;
  130643014043664["operator: <="]:::custom_operator;
  130643014043584 --> 130643014043664;
  130643014043744["identifier: n"]:::custom_identifier;
  130643014043664 --> 130643014043744;
  130643014043824["number: 1"]:::custom_number;
  130643014043664 --> 130643014043824;
  130643014043904["part: _"]:::custom_part;
  130643014043504 --> 130643014043904;
  130643014043984["number: 1"]:::custom_number;
  130643014043904 --> 130643014043984;
  130643014044064["part: else"]:::custom_part;
  130643014043504 --> 130643014044064;
  130643014044144["operator: *"]:::custom_operator;
  130643014044064 --> 130643014044144;
  130643014044224["identifier: n"]:::custom_identifier;
  130643014044144 --> 130643014044224;
  130643014044304["apply"]:::custom_apply;
  130643014044144 --> 130643014044304;
  130643014044384["identifier: f"]:::custom_identifier;
  130643014044304 --> 130643014044384;
  130643014044464["arguments"]:::custom_arguments;
  130643014044304 --> 130643014044464;
  130643014044544["operator: -"]:::custom_operator;
  130643014044464 --> 130643014044544;
  130643014044624["identifier: n"]:::custom_identifier;
  130643014044544 --> 130643014044624;
  130643014044704["number: 1"]:::custom_number;
  130643014044544 --> 130643014044704;

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
        name: '*'
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
              name: '-'
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
  "node_0xc00007c9c0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc00007c360" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007c9c0" -> "node_0xc00007c360";
  "node_0xc00007c300" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007c360" -> "node_0xc00007c300";
  "node_0xc00007c210" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c300" -> "node_0xc00007c210";
  "node_0xc00007c2a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007c300" -> "node_0xc00007c2a0";
  "node_0xc00007c270" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c2a0" -> "node_0xc00007c270";
  "node_0xc00007c990" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007c9c0" -> "node_0xc00007c990";
  "node_0xc00007c900" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc00007c990" -> "node_0xc00007c900";
  "node_0xc00007c4e0" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007c900" -> "node_0xc00007c4e0";
  "node_0xc00007c480" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c4e0" -> "node_0xc00007c480";
  "node_0xc00007c3c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c480" -> "node_0xc00007c3c0";
  "node_0xc00007c420" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c480" -> "node_0xc00007c420";
  "node_0xc00007c5a0" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007c900" -> "node_0xc00007c5a0";
  "node_0xc00007c540" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c5a0" -> "node_0xc00007c540";
  "node_0xc00007c8d0" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007c900" -> "node_0xc00007c8d0";
  "node_0xc00007c870" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c8d0" -> "node_0xc00007c870";
  "node_0xc00007c600" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c870" -> "node_0xc00007c600";
  "node_0xc00007c810" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007c870" -> "node_0xc00007c810";
  "node_0xc00007c660" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c810" -> "node_0xc00007c660";
  "node_0xc00007c7b0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007c810" -> "node_0xc00007c7b0";
  "node_0xc00007c780" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007c7b0" -> "node_0xc00007c780";
  "node_0xc00007c6c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007c780" -> "node_0xc00007c6c0";
  "node_0xc00007c720" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007c780" -> "node_0xc00007c720";
}
```

![Image generated for example](images/factorial.png)
