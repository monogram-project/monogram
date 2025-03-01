import xml.etree.ElementTree as ET
from pathlib import Path
import argparse

# Define colors for specific tag names
tag_colors = {
    "form": "lightpink",
    "part": "#FFD8E1",
    "apply": "lightgreen",
    "identifier": "Honeydew",
    "arguments": "PaleTurquoise",
    "operator": "#C0FFC0",
    "number": "lightgoldenrodyellow"
}

def xml_to_mermaid(xml_string):
    # Parse the XML string
    root = ET.fromstring(xml_string)
    
    # Initialize the Mermaid flowchart notation string
    mermaid = 'graph TD\n'
    
    def add_edges(node, mermaid, parent_id=None):
        # Create a unique identifier for the current node
        node_id = id(node)
        # Get the node's tag and attributes
        label = f'{node.tag}'
        if len(node.attrib) == 1:
            key, value = next(iter(node.attrib.items()))
            label = f'{node.tag}: {value}'
        # Determine the fill color based on the tag name
        fillcolor = tag_colors.get(node.tag, "lightgray")
        # Add the current node to the Mermaid notation with custom class
        node_class = node.tag
        mermaid += f'  {node_id}["{label}"]:::custom_{node_class};\n'
        if parent_id:
            # Add edge from parent to the current node
            mermaid += f'  {parent_id} --> {node_id};\n'
        for child in node:
            # Recursively add edges for child nodes
            mermaid = add_edges(child, mermaid, node_id)
        return mermaid
    
    # Start adding edges from the root
    mermaid = add_edges(root, mermaid)
    
    # Add custom styles for nodes
    mermaid += "\n"
    for tag, color in tag_colors.items():
        mermaid += f'classDef custom_{tag} fill:{color},stroke:#333,stroke-width:2px;\n'
    
    return mermaid


class Main:

    def __init__(self):
        parser = argparse.ArgumentParser(description="Convert XML to Mermaid flowchart format.")
        parser.add_argument("--input", type=Path, required=True, help="Path to the input XML file.")
        args = parser.parse_args()
        self._args = args

    def main(self):
        mermaid_representation = xml_to_mermaid(self._args.input.read_text())
        print(mermaid_representation)
        # with open("example.mmd", "w") as f:
        #     f.write(mermaid_representation)

if __name__ == "__main__":
    Main().main()
