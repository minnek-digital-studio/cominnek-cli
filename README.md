# Cominnek  [!["Pyhton"](https://img.shields.io/badge/python-3.9.1%20-gray.svg?longCache=true&logo=python&colorB=yellow)](https://www.python.org/downloads/release/python-391/)

Create commits & pull requests in an easy way

This is based in the [Git Version Control](https://docs.minnekdigital.com/development/git-version-control.html)

# Requirements 📃

- **[Git](https://git-scm.com/)**
- **[Python 3.9.1](https://www.python.org/downloads/release/python-391/) (or higher)**
- **[Github CLI](#install-github-cli)**
- **[Git-Flow (MacOS)](#git-flow-macos)**

# First steps 🦶:
## Install github-cli
`Comminek` use GitHub-CLI to interact with GitHub. For Example, create the `pull requests`. **[See more here](https://git-scm.com/)**

### macOS
`gh` is available via [Homebrew](https://brew.sh/), [MacPorts](https://www.macports.org/), [Conda](https://docs.conda.io/en/latest/), [Spack](https://spack.io/), and as a downloadable binary from the [releases page](https://github.com/cli/cli/releases/latest).

#### Homebrew

| Install:          | Upgrade:          |
| ----------------- | ----------------- |
| `brew install gh` | `brew upgrade gh` |
#### Conda

| Install:                                 | Upgrade:                                |
|------------------------------------------|-----------------------------------------|
| `conda install gh --channel conda-forge` | `conda update gh --channel conda-forge` |

Additional Conda installation options available on the [gh-feedstock page](https://github.com/conda-forge/gh-feedstock#installing-gh).

#### Spack

| Install:           | Upgrade:                                 |
| ------------------ | ---------------------------------------- |
| `spack install gh` | `spack uninstall gh && spack install gh` |
### Windows

`gh` is available via [WinGet](https://github.com/microsoft/winget-cli), [scoop](https://scoop.sh/), [Chocolatey](https://chocolatey.org/), and as downloadable MSI.

#### WinGet

| Install:            | Upgrade:            |
| ------------------- | --------------------|
| `winget install --id GitHub.cli` | `winget upgrade --id GitHub.cli` |

#### scoop

| Install:           | Upgrade:           |
| ------------------ | ------------------ |
| `scoop install gh` | `scoop update gh`  |

#### Chocolatey

| Install:           | Upgrade:           |
| ------------------ | ------------------ |
| `choco install gh` | `choco upgrade gh` |

#### Signed MSI

MSI installers are available for download on the [releases page](https://github.com/cli/cli/releases/latest).
## Install

1. Clone the repository
```bash
git clone git@github.com:Minnek-Digital-Studio/cominnek.git
```

```bash
cd ./cominnek
```

2. Install `cominnek` module

```bash
python setup.py install
```

Now you can run 
```bash
cominnek -v
```

# Git-Flow (MacOS)
To install git flow run: 

```bash
brew install git-flow
```

# Usage

- **[Update version](#update-version)**: Commit, push and publish the theme to BigCommerce
- **[Push](#push)**: Commit and push the branch to GitHub
- **[Publish](#publish)**: Commit, push and create the pull request to develop in GitHub.
- **[Feature](#feature)**: Create a new feature branch
- **[Stash](#stash)**: Stash changes in a branch to another branch 
  
## Update version
Commit, push and publish the theme to BigCommerce. The commit going to be "update version" by default. 

***Important:** Just use this in the test branch.*

```bash
cominnek update-version -a
```
the commit will be: `update version`


| flag               | type          | description                 |
| ------------------ | ------------- | --------------------------- |
| `-a --apply`       |Boolean        | Apply the theme automaticly |
## Push
Commit and push the branch to GitHub
```bash
cominnek push --feat "home" --message "do some modifications"
```
the commit will be: `feat(home):{Ticket} do some modifications`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-F --feat`        |string         | make the commit with the prefix feat() |
| `-f --fix`         |string         | make the commit with the prefix fix()  |
| `-m --message`     |string*        | Receives the commit message            |
| `-y --yes`         |Boolean (false)| Skip the confirmation question         |

*\* required*

## Publish
Commit, push and create the pull request to develop in GitHub.
The usage is the same as [push](#push) just with the difference that this creates a pull request.
```bash
cominnek publish --feat "home" --message "do some modifications"
```
## Feature
Create a new feature branch using git flow. Also, this going to checkout `develop` branch if it isn't checked out.

```bash
cominnek feature --ticket "{Ticket}"
```
This the equivalent of: `git flow feature start {Ticket}` or `git-flow feature start {Ticket}` on MacOS
| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-t --ticket`      |string*        | name to assign to the feature          |

*\* required*

## Stash
Stash all the changes in the current branch and apply the changes to another branch.

```bash
cominnek stash --ticket "{Ticket}"
```
*For features* Ex: `feature/{Ticket}`

```bash
cominnek stash --branch "{branch}"
```
*For branch that is not a feature* Ex: `develop`

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-t --ticket`      |string         | name of the feature that's will be applied the change     |
| `-b --branch`      |string         | name of the branch that's will be applied the changes       |

## Examples and more

You can add a body to commit using the `--message` flag twice.

```bash
cominnek push -f "home" -m "Changes in home page" -m "the title was aligned to right"
```
This is the same as `git commit` -m "fix(home):{Ticket} Changes in home page" -m "the title was aligned to right"`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

---
To do a commit without a scope, use the flag `-f, --fix` or `-F, --feat` with a space. 
```bash
cominnek push -F " " -m "theme setup"
```
the commit will be: `feat(): theme setup`

# Ticketing system
This takes the ticket number from the branch name. This is important to create the pull request with the correct ticket number.

Example: 

You are in the branch `feature/MJ-11` 

The ticker number will be `MJ-11`

If you execute the command:

```bash
cominnek push -F "home" -m "Changes in home page"
``` 

the commit will be: `feat(home):MJ-11 Changes in home page`

### What about if I am not in a feature branch?
This going to let you know that you are not in a feature branch and you will accept the commit without a ticket number.

Example: You're in the branch `develop`

```bash
$ cominnek push -F "home" -m "Changes in home page"
  This is not a feature. Do you want to continue? (yes or no)
        Commit message: "feat(home): Changes in home page"
```

> With ❤ by [isaacismaelx14](https://github.com/isaacismaelx14)