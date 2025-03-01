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
  133465698845936["form"]:::custom_form;
  133465698846016["part: def"]:::custom_part;
  133465698845936 --> 133465698846016;
  133465698846096["apply"]:::custom_apply;
  133465698846016 --> 133465698846096;
  133465698846176["identifier: f"]:::custom_identifier;
  133465698846096 --> 133465698846176;
  133465698846256["arguments"]:::custom_arguments;
  133465698846096 --> 133465698846256;
  133465698846336["identifier: n"]:::custom_identifier;
  133465698846256 --> 133465698846336;
  133465698846416["part: _"]:::custom_part;
  133465698845936 --> 133465698846416;
  133465698846496["form"]:::custom_form;
  133465698846416 --> 133465698846496;
  133465698846576["part: if"]:::custom_part;
  133465698846496 --> 133465698846576;
  133465698846656["operator: <="]:::custom_operator;
  133465698846576 --> 133465698846656;
  133465698846736["identifier: n"]:::custom_identifier;
  133465698846656 --> 133465698846736;
  133465698846816["number: 1"]:::custom_number;
  133465698846656 --> 133465698846816;
  133465698846896["part: _"]:::custom_part;
  133465698846496 --> 133465698846896;
  133465698846976["number: 1"]:::custom_number;
  133465698846896 --> 133465698846976;
  133465698847056["part: else"]:::custom_part;
  133465698846496 --> 133465698847056;
  133465698847136["operator: *"]:::custom_operator;
  133465698847056 --> 133465698847136;
  133465698847216["identifier: n"]:::custom_identifier;
  133465698847136 --> 133465698847216;
  133465698847296["apply"]:::custom_apply;
  133465698847136 --> 133465698847296;
  133465698847456["identifier: f"]:::custom_identifier;
  133465698847296 --> 133465698847456;
  133465698847616["arguments"]:::custom_arguments;
  133465698847296 --> 133465698847616;
  133465698847776["operator: -"]:::custom_operator;
  133465698847616 --> 133465698847776;
  133465698847936["identifier: n"]:::custom_identifier;
  133465698847776 --> 133465698847936;
  133465698848096["number: 1"]:::custom_number;
  133465698847776 --> 133465698848096;

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
<form>
    <part keyword="def">
        <apply kind="parentheses" separator="undefined">
            <identifier name="f"/>
            <arguments>
                <identifier name="n"/>
            </arguments>
        </apply>
    </part>
    <part keyword="_">
        <form>
            <part keyword="if">
                <operator name="&lt;=">
                    <identifier name="n"/>
                    <number value="1"/>
                </operator>
            </part>
            <part keyword="_">
                <number value="1"/>
            </part>
            <part keyword="else">
                <operator name="*">
                    <identifier name="n"/>
                    <apply kind="parentheses" separator="undefined">
                        <identifier name="f"/>
                        <arguments>
                            <operator name="-">
                                <identifier name="n"/>
                                <number value="1"/>
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
                    "children": [
                        {
                            "role": "part",
                            "keyword": "if",
                            "children": [
                                {
                                    "role": "operator",
                                    "name": "<=",
                                    "children": [
                                        {
                                            "role": "identifier",
                                            "name": "n"
                                        },
                                        {
                                            "role": "number",
                                            "value": 1
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
                                    "value": 1
                                }
                            ]
                        },
                        {
                            "role": "part",
                            "keyword": "else",
                            "children": [
                                {
                                    "role": "operator",
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
                                                            "name": "-",
                                                            "children": [
                                                                {
                                                                    "role": "identifier",
                                                                    "name": "n"
                                                                },
                                                                {
                                                                    "role": "number",
                                                                    "value": 1
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
```

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
    children:
    - role: part
      keyword: if
      children:
      - role: operator
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
  "125100571592000" [label="form", shape="box", fillcolor="lightpink"];
  "125100571592080" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "125100571592000" -> "125100571592080";
  "125100571592160" [label="apply", shape="box", fillcolor="lightgreen"];
  "125100571592080" -> "125100571592160";
  "125100571592240" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "125100571592160" -> "125100571592240";
  "125100571592320" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "125100571592160" -> "125100571592320";
  "125100571592400" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "125100571592320" -> "125100571592400";
  "125100571592480" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "125100571592000" -> "125100571592480";
  "125100571592560" [label="form", shape="box", fillcolor="lightpink"];
  "125100571592480" -> "125100571592560";
  "125100571592640" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "125100571592560" -> "125100571592640";
  "125100571592720" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "125100571592640" -> "125100571592720";
  "125100571592800" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "125100571592720" -> "125100571592800";
  "125100571592880" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "125100571592720" -> "125100571592880";
  "125100571592960" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "125100571592560" -> "125100571592960";
  "125100571593040" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "125100571592960" -> "125100571593040";
  "125100571593120" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "125100571592560" -> "125100571593120";
  "125100571593200" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "125100571593120" -> "125100571593200";
  "125100571593280" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "125100571593200" -> "125100571593280";
  "125100571593360" [label="apply", shape="box", fillcolor="lightgreen"];
  "125100571593200" -> "125100571593360";
  "125100571593520" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "125100571593360" -> "125100571593520";
  "125100571593680" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "125100571593360" -> "125100571593680";
  "125100571593840" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "125100571593680" -> "125100571593840";
  "125100571594000" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "125100571593840" -> "125100571594000";
  "125100571594160" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "125100571593840" -> "125100571594160";
}
```

![Image generated for example](images/factorial.png)
