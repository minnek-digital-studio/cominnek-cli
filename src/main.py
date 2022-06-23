import argparse
from . import git_controller as git
from .run_command import run_command

_VERSION = "1.1.1"

def addArguments():
    parser = argparse.ArgumentParser('cominnek')
    parser.add_argument('-v', '--version', action='version', version='%(prog)s ' + _VERSION)
    subparser = parser.add_subparsers(dest='command')

    update_version = subparser.add_parser('update-version', help='Create a commit with the message "update version" before pushing to github and BigCommerce')
    publish = subparser.add_parser('publish', help='create a pull request after commit')
    push = subparser.add_parser('push', help="commit and push the branch")
    feature = subparser.add_parser('feature', help="Create a feature branch")
    stash = subparser.add_parser('stash', help="Take all changes in current branch and stash them to another branch")

    feature.add_argument('-t', '--ticket', required=True, help="The feature name")

    def addArguments(toAdd):
        toAdd.add_argument("-f", "--fix", help="make the commit with the prefix fix()" )
        toAdd.add_argument("-F", "--feat", type=str, help="make the commit with the prefix feat()" )
        toAdd.add_argument("-m", "--message", action='append', type=str, required=True, help="the commit message")
        toAdd.add_argument("-y", "--yes", type=bool, default=False, action=argparse.BooleanOptionalAction, help="skip the question")

    addArguments(push)
    addArguments(publish)

    update_version.add_argument("-a", "--apply", type=bool, default=False, action=argparse.BooleanOptionalAction, help="Upload the theme to BigCommerce and Apply it" )
    
    stash.add_argument("-t", "--ticket", type=str, help="The ticket name")
    stash.add_argument("-b", "--branch", type=str, help="The branch name")
    return parser.parse_args()

def textValidate(text):
    if(text == "" or not isinstance(text, str)):
        return ""
    return text.strip()

def getState(args):
    state = None
    if(args.fix):
        state = f"fix({textValidate(args.fix)}):"
    if(args.feat):
        state = f"feat({textValidate(args.feat)}):"

    if(not state):
        raise Exception("Sorry, a state is required. use --feat or --fix")

    return state

def push(pr, args):
    state = getState(args)
    desc = None
    msg = args.message[0]
    ticket = git.is_feature()
    message = f"{state}{ticket} {msg}"

    if(len(args.message) > 1):
        desc = args.message[1]

    commit_exec = git.commit(message, desc, skip_question = args.yes)

    if(commit_exec == False):
        return
    
    git.push(pr)

    if(pr):
        git.pull_request(ticket)

def updateVersion(args):
    stencil = "stencil push"
    git.commit("update version", skip_question = True)
    git.push()

    if(args.apply):
        stencil = f"{stencil} -a"
    
    run_command(stencil)

def feature(args):
    git.feature_create(args.ticket)

def stash(args):
    branch = ""
    if(not args.ticket and not args.branch):
        raise Exception("Sorry, a ticket or branch is required. use --ticket or --branch")

    if(args.ticket): branch = f"feature/{args.ticket}"
    if(args.branch): branch = args.branch

    git.stash(branch)

def main():
    args = addArguments()

    if(args.command == "update-version"):
        updateVersion(args)

    if(args.command == "publish"):
        push(True, args)

    if(args.command == "push"):
        push(False, args)
    
    if(args.command == "feature"):
        feature(args)
    
    if(args.command == "stash"):
        stash(args)


if __name__ == "__main__":
    main()