import subprocess

def run_command(command, output=False):
    if(output):
        return str(subprocess.check_output(['powershell.exe', command]))
    else:
        subprocess.call(['powershell.exe', command])