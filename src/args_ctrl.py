import argparse

from .config import args

_VERSION = "1.3.2"

def addArguments():
    parser = argparse.ArgumentParser('cominnek')
    parser.add_argument('-v', '--version', action='version', version='%(prog)s ' + _VERSION)
    subparser = parser.add_subparsers(dest='command')

    for(key) in args.items:
        sub = subparser.add_parser(key['value'] , help=key['help'])
        
        for(arg) in key['flags']:
            name = f"--{arg['name']}"
            short = f"-{arg['short']}"
            
            sub.add_argument(short, name, type=arg['type'], required=arg["required"], help=arg['help'], action=arg['action'] or 'store', default=arg['default'])
    return parser.parse_args()
