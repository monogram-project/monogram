# Simple arithmetic expression

## Monogram

```txt
2 * 100 * 100 + 100 - 1
```

## Mermaid diagram

We can target Mermaid's flowchart as an output format. 
And this is what it looks like:

```mermaid
graph LR
  133569368393024["operator: -"]:::custom_operator;
  133569368393104["operator: +"]:::custom_operator;
  133569368393024 --> 133569368393104;
  133569368393184["operator: *"]:::custom_operator;
  133569368393104 --> 133569368393184;
  133569368393264["number: 2"]:::custom_number;
  133569368393184 --> 133569368393264;
  133569368393344["operator: *"]:::custom_operator;
  133569368393184 --> 133569368393344;
  133569368393424["number: 100"]:::custom_number;
  133569368393344 --> 133569368393424;
  133569368393504["number: 100"]:::custom_number;
  133569368393344 --> 133569368393504;
  133569368393584["number: 100"]:::custom_number;
  133569368393104 --> 133569368393584;
  133569368393664["number: 1"]:::custom_number;
  133569368393024 --> 133569368393664;

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
  "131376597796160" [label="operator: -", shape="box", fillcolor="#C0FFC0"];
  "131376597796240" [label="operator: +", shape="box", fillcolor="#C0FFC0"];
  "131376597796160" -> "131376597796240";
  "131376597796320" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131376597796240" -> "131376597796320";
  "131376597796400" [label="number: 2", shape="box", fillcolor="lightgoldenrodyellow"];
  "131376597796320" -> "131376597796400";
  "131376597796480" [label="operator: *", shape="box", fillcolor="#C0FFC0"];
  "131376597796320" -> "131376597796480";
  "131376597796560" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131376597796480" -> "131376597796560";
  "131376597796640" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131376597796480" -> "131376597796640";
  "131376597796720" [label="number: 100", shape="box", fillcolor="lightgoldenrodyellow"];
  "131376597796240" -> "131376597796720";
  "131376597796800" [label="number: 1", shape="box", fillcolor="lightgoldenrodyellow"];
  "131376597796160" -> "131376597796800";
}
```

![Image generated for example](images/arithmetic.png)
