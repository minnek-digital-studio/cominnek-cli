import subprocess
from .config import os as _config

def run_command(command, output=False):
    if(_config.is_windows):
        return _run_command_windows(command, output)
    else:
        return _run_command_unix(command, output)

def _run_command_windows(command, output=False):
    res = _run(['powershell.exe', command], output)
    
    if(output):
        return res

def _run_command_unix(command, output=False):
    res = _run(command, output)
    
    if(output):
        return res

def _run(cmd, output=False):
    try:
        if(output):
            msg = str(subprocess.check_output(cmd, shell=True))
        else:
            subprocess.call(cmd, shell=True)
    
    except subprocess.CalledProcessError as e:
        msg = str(e.stdout.decode('utf-8'))
        if(not output):
            print(msg)
            print("Exiting...")
            exit(1)
    
    finally:
        if(output):
            return msg
