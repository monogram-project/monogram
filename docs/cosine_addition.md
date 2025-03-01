# Simple arithmetic expression

## Monogram

```txt
cos(A) * cos(B) - sin(A) * sin(B)```

## Mermaid diagram

```mermaid
graph TD
  128805228348976["operator: -"]:::custom_operator;
  128805228349056["operator: *"]:::custom_operator;
  128805228348976 --> 128805228349056;
  128805228349136["apply"]:::custom_apply;
  128805228349056 --> 128805228349136;
  128805228349216["identifier: cos"]:::custom_identifier;
  128805228349136 --> 128805228349216;
  128805228349296["arguments"]:::custom_arguments;
  128805228349136 --> 128805228349296;
  128805228349376["identifier: A"]:::custom_identifier;
  128805228349296 --> 128805228349376;
  128805228349456["apply"]:::custom_apply;
  128805228349056 --> 128805228349456;
  128805228349536["identifier: cos"]:::custom_identifier;
  128805228349456 --> 128805228349536;
  128805228349616["arguments"]:::custom_arguments;
  128805228349456 --> 128805228349616;
  128805228349696["identifier: B"]:::custom_identifier;
  128805228349616 --> 128805228349696;
  128805228349776["operator: *"]:::custom_operator;
  128805228348976 --> 128805228349776;
  128805228349856["apply"]:::custom_apply;
  128805228349776 --> 128805228349856;
  128805228349936["identifier: sin"]:::custom_identifier;
  128805228349856 --> 128805228349936;
  128805228350016["arguments"]:::custom_arguments;
  128805228349856 --> 128805228350016;
  128805228350096["identifier: A"]:::custom_identifier;
  128805228350016 --> 128805228350096;
  128805228350176["apply"]:::custom_apply;
  128805228349776 --> 128805228350176;
  128805228350256["identifier: sin"]:::custom_identifier;
  128805228350176 --> 128805228350256;
  128805228350336["arguments"]:::custom_arguments;
  128805228350176 --> 128805228350336;
  128805228350416["identifier: B"]:::custom_identifier;
  128805228350336 --> 128805228350416;

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
<operator name="-">
    <operator name="*">
        <apply kind="parentheses" separator="undefined">
            <identifier name="cos"/>
            <arguments>
                <identifier name="A"/>
            </arguments>
        </apply>
        <apply kind="parentheses" separator="undefined">
            <identifier name="cos"/>
            <arguments>
                <identifier name="B"/>
            </arguments>
        </apply>
    </operator>
    <operator name="*">
        <apply kind="parentheses" separator="undefined">
            <identifier name="sin"/>
            <arguments>
                <identifier name="A"/>
            </arguments>
        </apply>
        <apply kind="parentheses" separator="undefined">
            <identifier name="sin"/>
            <arguments>
                <identifier name="B"/>
            </arguments>
        </apply>
    </operator>
</operator>
```

## JSON

```json
{
    "role": "operator",
    "name": "-",
    "children": [
        {
            "role": "operator",
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
```

## YAML

```yaml
role: operator
name: '-'
children:
- role: operator
  name: '*'
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
  "140463112832560" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "140463112832640" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "140463112832560" -> "140463112832640";
  "140463112832720" [label="apply", shape="box", fillcolor="lightgreen"];
  "140463112832640" -> "140463112832720";
  "140463112832800" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "140463112832720" -> "140463112832800";
  "140463112832880" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "140463112832720" -> "140463112832880";
  "140463112832960" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "140463112832880" -> "140463112832960";
  "140463112833040" [label="apply", shape="box", fillcolor="lightgreen"];
  "140463112832640" -> "140463112833040";
  "140463112833120" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "140463112833040" -> "140463112833120";
  "140463112833200" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "140463112833040" -> "140463112833200";
  "140463112833280" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "140463112833200" -> "140463112833280";
  "140463112833360" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "140463112832560" -> "140463112833360";
  "140463112833440" [label="apply", shape="box", fillcolor="lightgreen"];
  "140463112833360" -> "140463112833440";
  "140463112833520" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "140463112833440" -> "140463112833520";
  "140463112833600" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "140463112833440" -> "140463112833600";
  "140463112833680" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "140463112833600" -> "140463112833680";
  "140463112833760" [label="apply", shape="box", fillcolor="lightgreen"];
  "140463112833360" -> "140463112833760";
  "140463112833840" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "140463112833760" -> "140463112833840";
  "140463112833920" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "140463112833760" -> "140463112833920";
  "140463112834000" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "140463112833920" -> "140463112834000";
}
```

![Image generated for example](images/arithmetic.png)
