import subprocess
import argparse

parser = argparse.ArgumentParser('something')
subparser = parser.add_subparsers(dest='command')

publish = subparser.add_parser(
    'publish', help='create a pull request after commit')
push = subparser.add_parser('push', help="commit and push the branch")
update_version = subparser.add_parser('update-version')

update_version.add_argument("-a", "--apply", type=bool, default=False, action=argparse.BooleanOptionalAction)


def addArguments(toAdd):
    toAdd.add_argument("-f", "--fix", metavar="fix",
                       help="make the commit with the prefix fix()")
    toAdd.add_argument(
        "-F", "--feat",  help="make the commit with the prefix feat()")
    toAdd.add_argument("-m", "--message", type=str,
                       required=True, help="the commit message")
    toAdd.add_argument("-y", "--yes", type=bool, default=False, action=argparse.BooleanOptionalAction)



addArguments(push)
addArguments(publish)

args = parser.parse_args()


def question(question):
    i = 0
    while i < 2:
        answer = input(f"{question} (yes or no)")
        if any(answer.lower() == f for f in ["yes", 'y', '1', 'ye']):
            return True
        elif any(answer.lower() == f for f in ['no', 'n', '0']):
            return False
        else:
            i += 1
            if i < 4:
                print('Please enter yes or no')
            else:
                print("Nothing done")
                return False


def run_command(command, output=False):
    if(output):
        return subprocess.check_output(['powershell.exe', command])
    else:
        subprocess.call(['powershell.exe', command])


def pullRequet(ticket):
    title = f'Feature/{ticket}'
    message = f'### Ticket info \n- {ticket}\n\n- https://minnek.atlassian.net/browse/{ticket}'
    run_command(
        f'gh pr create -t "{title}" -b "{message}" -B develop -a "@me"')


def getBranch():
    outCmd = str(run_command("git rev-parse --abbrev-ref HEAD", True))
    outCmd = outCmd.split("'")[1]
    outCmd = outCmd.split("\\")[0]
    res = None

    if("feature" not in outCmd):
        res = question("This is not a feature. Do you want to continue?")

    if(res == True):
        return ""
    elif(res == False):
        exit()

    if("/" in outCmd):
        outCmd = outCmd.split('/')[1]

    return outCmd


def textValidate(text: str):
    if(text == "n"):
        return ""
    return text.strip()


def push(pr):
    state = None
    msg = args.message
    # ticket = getBranch()
    ticket = "getBranch()"

    if(args.fix):
        state = f"fix({textValidate(args.fix)}):"
    if(args.feat):
        state = f"feat({textValidate(args.feat)}):"

    if(not state):
        raise Exception("Sorry, a state is required. use --feat or --fix")

    message = f"{state}{ticket} {msg}"
    print(message)

    # run_command('git add .')
    # run_command('git status')

    if(not question("Do you want to continue?") or args.fix != False):
        print("no")
        # run_command('git reset .')
    else:
        print("yes")
        # run_command(f'git commit -m "{message}"')
        # run_command(f'git push')

        # if(pr):
        #     pullRequet(ticket)


def main():
    if(args.command == "update-version"):
        stencil = "stencil push"
        run_command('git add .')
        run_command('git commit -m "update version"; git push;')

        if(args.apply):
            stencil = f"{stencil} -a"
        
        run_command(stencil)

    if(args.command == "publish"):
        push(True)

    if(args.command == "push"):
        push(False)
