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
enddef```

## Mermaid diagram

```mermaid
graph TD
  123776115246400["form"]:::custom_form;
  123776115246480["part: def"]:::custom_part;
  123776115246400 --> 123776115246480;
  123776115246560["apply"]:::custom_apply;
  123776115246480 --> 123776115246560;
  123776115246640["identifier: f"]:::custom_identifier;
  123776115246560 --> 123776115246640;
  123776115246720["arguments"]:::custom_arguments;
  123776115246560 --> 123776115246720;
  123776115246800["identifier: n"]:::custom_identifier;
  123776115246720 --> 123776115246800;
  123776115246880["part: _"]:::custom_part;
  123776115246400 --> 123776115246880;
  123776115246960["form"]:::custom_form;
  123776115246880 --> 123776115246960;
  123776115247040["part: if"]:::custom_part;
  123776115246960 --> 123776115247040;
  123776115247120["operator: <="]:::custom_operator;
  123776115247040 --> 123776115247120;
  123776115247200["identifier: n"]:::custom_identifier;
  123776115247120 --> 123776115247200;
  123776115247280["number: 1"]:::custom_number;
  123776115247120 --> 123776115247280;
  123776115247360["part: _"]:::custom_part;
  123776115246960 --> 123776115247360;
  123776115247440["number: 1"]:::custom_number;
  123776115247360 --> 123776115247440;
  123776115247520["part: else"]:::custom_part;
  123776115246960 --> 123776115247520;
  123776115247600["operator: *"]:::custom_operator;
  123776115247520 --> 123776115247600;
  123776115247680["identifier: n"]:::custom_identifier;
  123776115247600 --> 123776115247680;
  123776115247760["apply"]:::custom_apply;
  123776115247600 --> 123776115247760;
  123776115247920["identifier: f"]:::custom_identifier;
  123776115247760 --> 123776115247920;
  123776115248080["arguments"]:::custom_arguments;
  123776115247760 --> 123776115248080;
  123776115248240["operator: -"]:::custom_operator;
  123776115248080 --> 123776115248240;
  123776115248400["identifier: n"]:::custom_identifier;
  123776115248240 --> 123776115248400;
  123776115248560["number: 1"]:::custom_number;
  123776115248240 --> 123776115248560;

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
  "132914985665856" [label="form", shape="box", fillcolor="lightpink"];
  "132914985665936" [label="part: def", shape="box", fillcolor="#FFD8E1"];
  "132914985665856" -> "132914985665936";
  "132914985666016" [label="apply", shape="box", fillcolor="lightgreen"];
  "132914985665936" -> "132914985666016";
  "132914985666096" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "132914985666016" -> "132914985666096";
  "132914985666176" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "132914985666016" -> "132914985666176";
  "132914985666256" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "132914985666176" -> "132914985666256";
  "132914985666336" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "132914985665856" -> "132914985666336";
  "132914985666416" [label="form", shape="box", fillcolor="lightpink"];
  "132914985666336" -> "132914985666416";
  "132914985666496" [label="part: if", shape="box", fillcolor="#FFD8E1"];
  "132914985666416" -> "132914985666496";
  "132914985666576" [label="operator: <=", shape="box", fillcolor="#C0FFC0"];
  "132914985666496" -> "132914985666576";
  "132914985666656" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "132914985666576" -> "132914985666656";
  "132914985666736" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "132914985666576" -> "132914985666736";
  "132914985666816" [label="part: _", shape="box", fillcolor="#FFD8E1"];
  "132914985666416" -> "132914985666816";
  "132914985666896" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "132914985666816" -> "132914985666896";
  "132914985666976" [label="part: else", shape="box", fillcolor="#FFD8E1"];
  "132914985666416" -> "132914985666976";
  "132914985667056" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "132914985666976" -> "132914985667056";
  "132914985667136" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "132914985667056" -> "132914985667136";
  "132914985667216" [label="apply", shape="box", fillcolor="lightgreen"];
  "132914985667056" -> "132914985667216";
  "132914985667376" [label="identifier: f", shape="box", fillcolor="Honeydew"];
  "132914985667216" -> "132914985667376";
  "132914985667536" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "132914985667216" -> "132914985667536";
  "132914985667696" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "132914985667536" -> "132914985667696";
  "132914985667856" [label="identifier: n", shape="box", fillcolor="Honeydew"];
  "132914985667696" -> "132914985667856";
  "132914985668016" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "132914985667696" -> "132914985668016";
}
```

![Image generated for example](images/factorial.png)
