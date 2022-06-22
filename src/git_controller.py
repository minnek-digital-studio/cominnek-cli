from cProfile import run
from .run_command import run_command
from .question import question

def _get_commit_cmd(message, desc = None) -> str:
    commit_command = f'git commit'
    msg = message
    commit_command = f'{commit_command} -m "{message}"'

    print(f'\n\tCommit message: "{message}"')

    if(desc):
        print(f"\tBody: {desc}")
        commit_command = f'{commit_command} -m "{desc}"'
    print("") # new line
    return commit_command

def add(show_status=False):
    run_command('git add .')
    if(show_status):
        run_command('git status')

def commit(message, desc=None, skip_question=False):
    commit_command = _get_commit_cmd(message, desc)
    add(True)

    if(skip_question == False):
        if(not question("Do you want to continue?")):
            run_command('git reset .')
            print("exiting...")
            return False

    run_command(commit_command)
    return True

def push():
    run_command('git push')

def pull_request(ticket):
    title = f'Feature/{ticket}'
    message = f'### Ticket info \n- {ticket}\n\n- https://minnek.atlassian.net/browse/{ticket}'
    run_command(
        f'gh pr create -t "{title}" -b "{message}" -B develop -a "@me"')

def get_current_branch():
    outCmd = str(run_command("git rev-parse --abbrev-ref HEAD", True))
    outCmd = outCmd.split("'")[1]
    outCmd = outCmd.split("\\")[0]
    return outCmd

def is_feature():
    outCmd = get_current_branch()
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

def pull():
    run_command('git pull')

def brach_switch(branch, make_pull_request=False):
    outCmd = run_command(f'git checkout {branch}', True)

    if(outCmd.find(f"Your branch is up to date with 'origin/{branch}'") == -1):
        print("Branch has changes. Pulling latest changes...")
        if(make_pull_request):
           pull()
    else:
        print("Branch is up to date.\n")

def feature_create(ticket):
    outCmd = get_current_branch()

    if("develop" not in outCmd):
        print("This is not the develop branch. Switching to develop.\n")
        brach_switch("develop", True)
        print("Ready to create feature branch.\n")
    
    print(f"Creating feature branch for {ticket}")
    run_command(f'git flow feature start {ticket}')

def stash(branch):
    run_command(f'git stash')
    run_command(f'git checkout {branch}')
    run_command(f'git stash apply')
