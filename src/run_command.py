import subprocess
from . import config

def run_command(command, output=False):
    if(config.is_windows):
        return _run_command_windows(command, output)
    else:
        return _run_command_unix(command, output)


def _run_command_windows(command, output=False):
    if(output):
        return str(subprocess.check_output(['powershell.exe', command]))
    else:
        subprocess.call(['powershell.exe', command])

def _run_command_unix(command, output=False):
    if(output):
        return str(subprocess.check_output(command, shell=True))
    else:
        subprocess.call(command, shell=True)