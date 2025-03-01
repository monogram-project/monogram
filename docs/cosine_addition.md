# Cosine addition formula

## Monogram

```txt
cos(A) * cos(B) - sin(A) * sin(B)

```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  130824905262640["operator: -"]:::custom_operator;
  130824905262720["operator: *"]:::custom_operator;
  130824905262640 --> 130824905262720;
  130824905262800["apply"]:::custom_apply;
  130824905262720 --> 130824905262800;
  130824905262880["identifier: cos"]:::custom_identifier;
  130824905262800 --> 130824905262880;
  130824905262960["arguments"]:::custom_arguments;
  130824905262800 --> 130824905262960;
  130824905263040["identifier: A"]:::custom_identifier;
  130824905262960 --> 130824905263040;
  130824905263120["apply"]:::custom_apply;
  130824905262720 --> 130824905263120;
  130824905263200["identifier: cos"]:::custom_identifier;
  130824905263120 --> 130824905263200;
  130824905263280["arguments"]:::custom_arguments;
  130824905263120 --> 130824905263280;
  130824905263360["identifier: B"]:::custom_identifier;
  130824905263280 --> 130824905263360;
  130824905263440["operator: *"]:::custom_operator;
  130824905262640 --> 130824905263440;
  130824905263520["apply"]:::custom_apply;
  130824905263440 --> 130824905263520;
  130824905263600["identifier: sin"]:::custom_identifier;
  130824905263520 --> 130824905263600;
  130824905263680["arguments"]:::custom_arguments;
  130824905263520 --> 130824905263680;
  130824905263760["identifier: A"]:::custom_identifier;
  130824905263680 --> 130824905263760;
  130824905263840["apply"]:::custom_apply;
  130824905263440 --> 130824905263840;
  130824905263920["identifier: sin"]:::custom_identifier;
  130824905263840 --> 130824905263920;
  130824905264000["arguments"]:::custom_arguments;
  130824905263840 --> 130824905264000;
  130824905264080["identifier: B"]:::custom_identifier;
  130824905264000 --> 130824905264080;

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
  "131276364416560" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "131276364416640" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131276364416560" -> "131276364416640";
  "131276364416720" [label="apply", shape="box", fillcolor="lightgreen"];
  "131276364416640" -> "131276364416720";
  "131276364416800" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "131276364416720" -> "131276364416800";
  "131276364416880" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131276364416720" -> "131276364416880";
  "131276364416960" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "131276364416880" -> "131276364416960";
  "131276364417040" [label="apply", shape="box", fillcolor="lightgreen"];
  "131276364416640" -> "131276364417040";
  "131276364417120" [label="identifier: cos", shape="box", fillcolor="Honeydew"];
  "131276364417040" -> "131276364417120";
  "131276364417200" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131276364417040" -> "131276364417200";
  "131276364417280" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "131276364417200" -> "131276364417280";
  "131276364417360" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131276364416560" -> "131276364417360";
  "131276364417440" [label="apply", shape="box", fillcolor="lightgreen"];
  "131276364417360" -> "131276364417440";
  "131276364417520" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "131276364417440" -> "131276364417520";
  "131276364417600" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131276364417440" -> "131276364417600";
  "131276364417680" [label="identifier: A", shape="box", fillcolor="Honeydew"];
  "131276364417600" -> "131276364417680";
  "131276364417760" [label="apply", shape="box", fillcolor="lightgreen"];
  "131276364417360" -> "131276364417760";
  "131276364417840" [label="identifier: sin", shape="box", fillcolor="Honeydew"];
  "131276364417760" -> "131276364417840";
  "131276364417920" [label="arguments", shape="box", fillcolor="PaleTurquoise"];
  "131276364417760" -> "131276364417920";
  "131276364418000" [label="identifier: B", shape="box", fillcolor="Honeydew"];
  "131276364417920" -> "131276364418000";
}
```

![Image generated for example](images/cosine_addition.png)
