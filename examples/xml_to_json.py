import xml.etree.ElementTree as ET
import json
import argparse
from pathlib import Path

def xml_to_json(element):
    obj = {"role": element.tag}
    if element.tag == "number":
        for attr_name, attr_value in element.attrib.items():
            if attr_name == "value":
                s = attr_value
                try:
                    obj[attr_name] = int(s)
                except ValueError:
                    obj[attr_name] = float(s)                
            else:
                obj[attr_name] = attr_value
    else:
        for attr_name, attr_value in element.attrib.items():
            obj[attr_name] = attr_value
    
    children = [xml_to_json(child) for child in element]
    if children:
        obj["children"] = children
    
    return obj

class Main:

    def __init__(self):
        parser = argparse.ArgumentParser(description="Convert XML to Mermaid flowchart format.")
        parser.add_argument("--input", type=Path, required=True, help="Path to the input XML file.")
        args = parser.parse_args()
        self._args = args

    def main(self):
        xml_str = self._args.input.read_text()
        root = ET.fromstring(xml_str)
        json_obj = xml_to_json(root)
        json_str = json.dumps(json_obj, indent=4)
        print(json_str)

if __name__ == "__main__":
    Main().main()
