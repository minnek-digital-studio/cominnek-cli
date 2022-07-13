from .args_ctrl import addArguments
from .git_controller import GitController
from .run_command import run_command

gitCtrl = GitController()

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
    ticket = gitCtrl.is_feature()
    gitCtrl.add(True)

    commit_exec = commit(args, ticket)

    if(commit_exec == False):
        return
    
    gitCtrl.push(pr)

    if(pr):
        pull_request(None, ticket)

def commit(args, ticket = None):
    state = getState(args)
    desc = None
    msg = args.message[0]
    
    if(args.add_all):
        print("Adding all files...")
        gitCtrl.add(True)

    if(ticket == None):
        ticket = gitCtrl.is_feature()
    
    message = f"{state}{ticket} {msg}"

    if(len(args.message) > 1):
            desc = args.message[1]

    return gitCtrl.commit(message, desc, skip_question=args.yes)

def pull_request(args, ticket = None):
    if(ticket == None):
        ticket = gitCtrl.is_feature() or args.ticket
        
        if(not ticket):
            print("\nSorry, a ticket is required. use --ticket")
            exit(1)

    gitCtrl.pull_request(ticket)

def update_version(args):
    stencil = "stencil push"
    gitCtrl.commit("update version", skip_question = True)
    gitCtrl.push()

    if(args.apply):
        stencil = f"{stencil} -a"
    
    run_command(stencil)

def feature(args):
    gitCtrl.feature_create(args.ticket, args.stash)

def stash(args):
    branch = ""
    if(not args.ticket and not args.branch):
        raise Exception("Sorry, a ticket or branch is required. use --ticket or --branch")

    if(args.ticket): branch = f"feature/{args.ticket}"
    if(args.branch): branch = args.branch

    gitCtrl.stash(branch)

def main():
    args = addArguments()

    if(args.command == "update-version"):
        update_version(args)

    if(args.command == "publish"):
        push(True, args)

    if(args.command == "push"):
        push(False, args)
    
    if(args.command == "feature"):
        feature(args)
    
    if(args.command == "stash"):
        stash(args)

    if(args.command == "pr"):
        pull_request(args)

    if(args.command == "commit"):
        commit(args)

if __name__ == "__main__":
    main()