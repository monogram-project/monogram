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
  133213536224576["form: surround"]:::custom_form;
  133213536224656["part: def"]:::custom_part;
  133213536224576 --> 133213536224656;
  133213536224736["apply"]:::custom_apply;
  133213536224656 --> 133213536224736;
  133213536224816["identifier: f"]:::custom_identifier;
  133213536224736 --> 133213536224816;
  133213536224896["arguments"]:::custom_arguments;
  133213536224736 --> 133213536224896;
  133213536224976["identifier: n"]:::custom_identifier;
  133213536224896 --> 133213536224976;
  133213536225056["part: _"]:::custom_part;
  133213536224576 --> 133213536225056;
  133213536225136["form: surround"]:::custom_form;
  133213536225056 --> 133213536225136;
  133213536225216["part: if"]:::custom_part;
  133213536225136 --> 133213536225216;
  133213536225296["operator: <="]:::custom_operator;
  133213536225216 --> 133213536225296;
  133213536225376["identifier: n"]:::custom_identifier;
  133213536225296 --> 133213536225376;
  133213536225456["number: 1"]:::custom_number;
  133213536225296 --> 133213536225456;
  133213536225536["part: _"]:::custom_part;
  133213536225136 --> 133213536225536;
  133213536225616["number: 1"]:::custom_number;
  133213536225536 --> 133213536225616;
  133213536225696["part: else"]:::custom_part;
  133213536225136 --> 133213536225696;
  133213536225776["operator: *"]:::custom_operator;
  133213536225696 --> 133213536225776;
  133213536225856["identifier: n"]:::custom_identifier;
  133213536225776 --> 133213536225856;
  133213536225936["apply"]:::custom_apply;
  133213536225776 --> 133213536225936;
  133213536226016["identifier: f"]:::custom_identifier;
  133213536225936 --> 133213536226016;
  133213536226096["arguments"]:::custom_arguments;
  133213536225936 --> 133213536226096;
  133213536226176["operator: -"]:::custom_operator;
  133213536226096 --> 133213536226176;
  133213536226256["identifier: n"]:::custom_identifier;
  133213536226176 --> 133213536226256;
  133213536226336["number: 1"]:::custom_number;
  133213536226176 --> 133213536226336;

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
  "node_0xc00007ea50" [label="unit: factorial.mg", shape="box", fillcolor="lightgray"];
  "node_0xc00007e9c0" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc00007ea50" -> "node_0xc00007e9c0";
  "node_0xc00007e360" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007e9c0" -> "node_0xc00007e360";
  "node_0xc00007e300" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e360" -> "node_0xc00007e300";
  "node_0xc00007e210" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e300" -> "node_0xc00007e210";
  "node_0xc00007e2a0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e300" -> "node_0xc00007e2a0";
  "node_0xc00007e270" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e2a0" -> "node_0xc00007e270";
  "node_0xc00007e990" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007e9c0" -> "node_0xc00007e990";
  "node_0xc00007e900" [label="form: surround", shape="box", fillcolor="lightpink"];
  "node_0xc00007e990" -> "node_0xc00007e900";
  "node_0xc00007e4e0" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007e900" -> "node_0xc00007e4e0";
  "node_0xc00007e480" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e4e0" -> "node_0xc00007e480";
  "node_0xc00007e3c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e480" -> "node_0xc00007e3c0";
  "node_0xc00007e420" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e480" -> "node_0xc00007e420";
  "node_0xc00007e5a0" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007e900" -> "node_0xc00007e5a0";
  "node_0xc00007e540" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e5a0" -> "node_0xc00007e540";
  "node_0xc00007e8d0" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "node_0xc00007e900" -> "node_0xc00007e8d0";
  "node_0xc00007e870" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e8d0" -> "node_0xc00007e870";
  "node_0xc00007e600" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e870" -> "node_0xc00007e600";
  "node_0xc00007e810" [label="apply", shape="box", fillcolor="lightgreen"];
  "node_0xc00007e870" -> "node_0xc00007e810";
  "node_0xc00007e660" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e810" -> "node_0xc00007e660";
  "node_0xc00007e7b0" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "node_0xc00007e810" -> "node_0xc00007e7b0";
  "node_0xc00007e780" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "node_0xc00007e7b0" -> "node_0xc00007e780";
  "node_0xc00007e6c0" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "node_0xc00007e780" -> "node_0xc00007e6c0";
  "node_0xc00007e720" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "node_0xc00007e780" -> "node_0xc00007e720";
}
```

![Image generated for example](images/factorial.png)
