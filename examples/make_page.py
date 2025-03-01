import argparse
from pathlib import Path
import re

def lines( path: Path ):
    with path.open() as f:
        for line in f:
            yield line

class Main:

    def __init__(self):
        parser = argparse.ArgumentParser(description='Create a page')
        parser.add_argument('--stem', type=str, help='name of the page')
        self._args = parser.parse_args()

    def main(self):
        template = tuple( lines( Path( self._args.stem ).with_suffix('.md') ) )
        target = ( Path('_build') / self._args.stem ).with_suffix('.md')
        print(target)
        with target.open('w') as f:
            for line in template:
                if m := re.match( r'#include\s+([\w_]+)', line ):
                    extn = m.group(1)
                    include_file = ( Path( '_build') / self._args.stem ).with_suffix( '.' + extn )
                    for line in lines( include_file ):
                        f.write( line )
                else:
                    f.write( line )
        
        
if __name__ == "__main__":
    Main().main()