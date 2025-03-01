# Simple arithmetic expression

## Monogram

```txt
2 * 100 * 100 + 100 - 1
```

## Mermaid diagram

```mermaid
graph TD
  128564441744704["operator: -"]:::custom_operator;
  128564441744784["operator: +"]:::custom_operator;
  128564441744704 --> 128564441744784;
  128564441744864["operator: *"]:::custom_operator;
  128564441744784 --> 128564441744864;
  128564441744944["number: 2"]:::custom_number;
  128564441744864 --> 128564441744944;
  128564441745024["operator: *"]:::custom_operator;
  128564441744864 --> 128564441745024;
  128564441745104["number: 100"]:::custom_number;
  128564441745024 --> 128564441745104;
  128564441745184["number: 100"]:::custom_number;
  128564441745024 --> 128564441745184;
  128564441745264["number: 100"]:::custom_number;
  128564441744784 --> 128564441745264;
  128564441745344["number: 1"]:::custom_number;
  128564441744704 --> 128564441745344;

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
    <operator name="+">
        <operator name="*">
            <number value="2"/>
            <operator name="*">
                <number value="100"/>
                <number value="100"/>
            </operator>
        </operator>
        <number value="100"/>
    </operator>
    <number value="1"/>
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
            "name": "+",
            "children": [
                {
                    "role": "operator",
                    "name": "*",
                    "children": [
                        {
                            "role": "number",
                            "value": 2
                        },
                        {
                            "role": "operator",
                            "name": "*",
                            "children": [
                                {
                                    "role": "number",
                                    "value": 100
                                },
                                {
                                    "role": "number",
                                    "value": 100
                                }
                            ]
                        }
                    ]
                },
                {
                    "role": "number",
                    "value": 100
                }
            ]
        },
        {
            "role": "number",
            "value": 1
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
  name: +
  children:
  - role: operator
    name: '*'
    children:
    - role: number
      value: 2
    - role: operator
      name: '*'
      children:
      - role: number
        value: 100
      - role: number
        value: 100
  - role: number
    value: 100
- role: number
  value: 1

```

## Graphviz Dot format

```dot
digraph G {
  bgcolor="transparent";
  node [shape="box", style="filled", fontname="Ubuntu Mono"];
  "139254878537024" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "139254878537104" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "139254878537024" -> "139254878537104";
  "139254878537184" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "139254878537104" -> "139254878537184";
  "139254878537264" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "139254878537184" -> "139254878537264";
  "139254878537344" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "139254878537184" -> "139254878537344";
  "139254878537424" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139254878537344" -> "139254878537424";
  "139254878537504" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139254878537344" -> "139254878537504";
  "139254878537584" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "139254878537104" -> "139254878537584";
  "139254878537664" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "139254878537024" -> "139254878537664";
}
```

![Image generated for example](images/arithmetic.png)
