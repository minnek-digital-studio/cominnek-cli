import os

def create(fileName, content):
    f = open(fileName, "w")
    f.write(content)
    f.close()


def delete(fileName):
    os.remove(fileName)
    