import os
import platform

current_os = platform.system()

is_windows = current_os == "Windows"
is_linux = current_os == "Linux"
is_mac = current_os == "Darwin"

def get_path():
    path = ''
    if(is_windows):
        path = f'{os.getenv("APPDATA")}\\cominnek'
    elif(is_linux):
        path = f'{os.getenv("HOME")}/cominnek'
    elif(is_mac):
        path = f'{os.getenv("HOME")}/Library/Application Support/cominnek'
    else:
        path = f'{os.getenv("HOME")}/cominnek'

    if(not os.path.exists(path)):
        os.mkdir(path)
    
    return path

def path(_path, _file):
    if(is_windows):
        return f"{_path}\\{_file}"
    elif(is_linux):
        return f"{_path}/{_file}"
    elif(is_mac):
        return f"{_path}/{_file}"
    else:
        return f"{_path}/{_file}"

