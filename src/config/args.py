import argparse

def ticket_flag(not_required=False):
    return {
        "name": "ticket",
        "short": "t",
        "help": "The ticket name",
        "required": not not_required,
        "type": str
    }

yes_flag = {
    "name": "yes",
    "short": "y",
    "default": False,
    "action": argparse.BooleanOptionalAction,
    "help": "skip the question",
    "type": bool
}

def commits_flags(add_all=False):
    __add_all = {
        "name": "add-all",
        "short": "a",
        "help": "Add all files to the commit",
        "action": argparse.BooleanOptionalAction,
        "default": False,
        "type": bool
    }

    flags = [
        {
            "name": "fix",
            "short": "f",
            "help": "make the commit with the prefix fix()",
            "type": str
        },
        {
            "name": "feat",
            "short": "F",
            "help": "make the commit with the prefix feat()",
            "type": str
        },
        {
            "name": "message",
            "short": "m",
            "help": "the commit message",
            "action": "append",
            "required": True,
            "type": str
        },
        yes_flag
    ]

    if(add_all):
        flags.append(__add_all)

    return flags

def git_publish_flags():
    flags = commits_flags()
    flags.append(
        {
            "name": "merge",
            "short": "M",
            "help": "merge the commit with the received branch",
            "type": str
        }
    )
    return flags

items = [
    {
        "value": "feature",
        "help": "Create a feature branch",
        "flags": [
            {
                "name": "stash",
                "short": "s",
                "help": "Stash the current branch",
                "default": False,
                "action": argparse.BooleanOptionalAction,
                "type": bool
            },
            ticket_flag()
        ]
    },
    {
        "value": "update-version",
        "help": "Create a commit with the message \"update version\" before pushing to github and BigCommerce",
        "flags": [
            {
                "name": "apply",
                "short": "a",
                "help": "Upload the theme to BigCommerce and Apply it",
                "action": argparse.BooleanOptionalAction,
                "default": False,
                "type": bool
            }
        ]
    },
    {
        "value": "publish",
        "help": "create a pull request after commit",
        "flags": git_publish_flags()
    },
    {
        "value": "push",
        "help": "commit and push the branch",
        "flags": git_publish_flags()
    },
    {
        "value": "stash",
        "help": "Take all changes in current branch and stash them to another branch",
        "flags": [
            {
                "name": "branch",
                "short": "b",
                "help": "The branch name",
                "type": str
            },
            ticket_flag()
        ]
    },
    {
        "value": "pr",
        "help": "Create a pull request directly to develop's branch",
        "flags": [
            ticket_flag(not_required=True)
        ]
    },
    {
        "value": "commit",
        "help": "Create a commit",
        "flags": commits_flags(add_all=True)
    },
    {
        "value": "merge",
        "help": "Merge the current branch with the received branch",
        "flags": [
            {
                "name": "branch",
                "short": "b",
                "help": "The branch name",
                "type": str
            },
        ]
    }
]

newItem = []

for item in items:
    for flag in item["flags"]:
        if(not "required" in flag):
            flag["required"] = False
        if(not "action" in flag):
            flag["action"] = None
        if(not "default" in flag):
            flag["default"] = None
