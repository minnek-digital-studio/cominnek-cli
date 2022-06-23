import platform

current_os = platform.system()

is_windows = current_os == "Windows"
is_linux = current_os == "Linux"
is_mac = current_os == "Darwin"