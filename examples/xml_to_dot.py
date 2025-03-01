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

def xml_to_dot(xml_string):
    # Parse the XML string
    root = ET.fromstring(xml_string)
    
    # Initialize the DOT notation string
    # dot = "digraph G {\n"
    dot = 'digraph G {\n  bgcolor="transparent";\n  node [shape="box", style="filled", fontname="Ubuntu Mono"];\n'
    
    def add_edges(node, dot, parent_id=None):
        # Create a unique identifier for the current node
        node_id = id(node)
        # Get the node's tag and attributes
        label = f'{node.tag}'
        if len(node.attrib) == 1:
            key, value = next(iter(node.attrib.items()))
            label = f'{node.tag}: {value}'
        # Determine the fill color based on the tag name
        fillcolor = tag_colors.get(node.tag, "lightgray")
        # Add the current node to the DOT notation
        dot += f'  "{node_id}" [label="{label}", shape="box", fillcolor="{fillcolor}"];\n'
        if parent_id:
            # Add edge from parent to the current node
            dot += f'  "{parent_id}" -> "{node_id}";\n'
        for child in node:
            # Recursively add edges for child nodes
            dot = add_edges(child, dot, node_id)
        return dot
    
    # Start adding edges from the root
    dot = add_edges(root, dot)
    
    # Close the DOT notation string
    dot += "}"
    return dot


class Main:

    def __init__(self):
        parser = argparse.ArgumentParser(description="Convert XML to DOT format.")
        parser.add_argument("--input", type=Path, required=True, help="Path to the input XML file.")
        args = parser.parse_args()
        self._args = args

    def main(self):
        dot_representation = xml_to_dot(self._args.input.read_text())
        print(dot_representation)
        # with open("example.dot", "w") as f:
        #     f.write(dot_representation)



if __name__ == "__main__":
    Main().main()
