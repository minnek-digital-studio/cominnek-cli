from .run_command import run_command
from .question import question
from .config import os as _config
from . import file
from colorama import Fore
from colorama import Style

class GitController:
    def __init__(self) -> None:
        pass

    def __get_commit_cmd(self, message, desc=None) -> str:
        commit_command = f'git commit'
        commit_command = f'{commit_command} -m "{message}"'

        print(f'\n\tCommit message: "{message}"')

        if(desc):
            print(f"\tBody: {desc}")
            commit_command = f'{commit_command} -m "{desc}"'
        print("")  # new line
        return commit_command

    def add(self, show_status=False):
        run_command('git add .')
        if(show_status):
            run_command('git status')

    def commit(self, message, desc=None, skip_question=False):
        commit_command = self.__get_commit_cmd(message, desc)

        if(skip_question == False):
            res = question(f"{Fore.LIGHTWHITE_EX}Do you want to continue?{Style.RESET_ALL}")

            if(not res):
                run_command('git reset .')
                print("exiting...")
                return False

        run_command(commit_command)
        return True

    def push(self, publish=False):
        if(publish):
            print("Pushing to remote...")
            run_command(
                f'git push --set-upstream origin {self.get_current_branch()}')
        else:
            run_command('git push')

    def pull_request(self, ticket):
        title = f'Feature/{ticket}'
        message = f'### Ticket info\n- {ticket}\n- https://minnek.atlassian.net/browse/{ticket}'
        route = _config.path(_config.get_path(), 'pull_request.md')
        file.create(route, message)

        run_command(f'gh pr create -t "{title}" -F "{route}" -B develop -a "@me" -d')

    def get_current_branch(self):
        outCmd = str(run_command("git rev-parse --abbrev-ref HEAD", True))
        outCmd = outCmd.split("'")[1]
        outCmd = outCmd.split("\\")[0]
        return outCmd

    def is_feature(self):
        outCmd = self.get_current_branch()
        res = None

        if("feature" not in outCmd):
            msg = "This is not a feature. Do you want to continue?"
            res = question(f"{Fore.LIGHTWHITE_EX}{msg}{Style.RESET_ALL} ")

        if(res == True):
            return ""
        elif(res == False):
            exit()

        if("/" in outCmd):
            outCmd = outCmd.split('/')[1]

        return outCmd

    def pull(self) -> str:
        return run_command('git pull', True)

    def __check_conflicts(self, outcmd, callback=None):
        if(outcmd.find("CONFLICT") != -1 or outcmd.find("conflict") != -1):
            msg = "Conflicts found. Please checkout the branch, resolve conflicts and try again.\n"
            print(f"\n{Fore.LIGHTRED_EX}{msg}{Style.RESET_ALL}")
            print(outcmd)
            print(f"{Fore.LIGHTYELLOW_EX}Cleanin...\n{Style.RESET_ALL}")
            run_command('git reset --merge ORIG_HEAD')

            if(callback != None):
                callback()

            print("Exiting...")
            exit(1)

    def brach_switch(self, branch, branchCome=None, _stash=False, make_pull_request=False):
        outCmd = run_command(f'git checkout {branch}', True)

        if(outCmd.find(f"Your branch is up to date with 'origin/{branch}'") == -1):
            print("Branch has changes. Pulling latest changes...")

            if(make_pull_request):
                p_out = self.pull()

            def fallBack():
                if(branchCome):
                    self.brach_switch(branchCome)

                    if(_stash):
                        run_command(f'git stash apply')

            self.__check_conflicts(p_out, fallBack)
        else:
            print("Branch is up to date.\n")

    def feature_create(self, ticket, stash=False):
        if(stash):
            run_command(f'git stash')

        outCmd = self.get_current_branch()

        if("develop" not in outCmd):
            print("This is not the develop branch. Switching to develop.\n")

        self.brach_switch("develop", outCmd, stash, True)
        print("Ready to create feature branch.\n")

        print(f"Creating feature branch for {ticket}")

        if(_config.is_windows):
            run_command(f'git flow feature start {ticket}')
        else:
            run_command(f'git-flow feature start {ticket}')

        if(stash):
            print("Stashing changes... \n")
            run_command(f'git stash apply')

    def merge(self, swich_branch, mergeable_branch):
        self.brach_switch(swich_branch, mergeable_branch, False, True)
        self.__merge(mergeable_branch)

    def __merge(self, branch):
        run_command(f'git merge {branch}')

    def stash(self, branch):
        run_command(f'git stash')
        run_command(f'git checkout {branch}')
        run_command(f'git stash apply')
