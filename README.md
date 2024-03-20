<p align="center">
  <a href="https://minnekdigital.com/">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://assets.minnekdigital.com/logo-md.jpg">
      <img alt="Minnek Logo" src="https://assets.minnekdigital.com/logo-md.jpg">
    </picture>    
  </a>
</p>

---

# Cominnek  [!["Go"](https://img.shields.io/badge/go-1.18.3%20-gray.svg?longCache=true&logo=go&colorB=blue)](https://go.dev/doc/go1.18)

<img src="./assets/banner.png" />

Create commits & pull requests easily.  `Cominnek` is based on [TAYO]([https://mnk-docs.ngrok.io/guide/development/version-control/branch-management.html](https://docs.minnekdigital.com/engineering/development/version-control/branch/tayo-flow.html)) by Minnek.
## Index
 - **[Requirements](#requirements)**
 - **[Install](#installation)**
 - **[Update](#update)**
 - **[First Step](#first-steps)**
 - **[Usage](#usage)**
 - **[Ticketing System](#ticketing-system)**

# Requirements

- **[Git](https://git-scm.com/)**

# Installation

1. Download installer

[![Macos](https://img.shields.io/badge/mac%20os-0078D6?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.dmg) [![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white
)](https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.exe) [![Linux](https://img.shields.io/badge/Linux-0078D6?style=for-the-badge&logo=linux&logoColor=white)](https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-4.0.0.deb)

2. Run installer ([See MAC Os steps](#how-to-install-on-mac))

3. Now you can run 
```bash
cominnek -v
```

## How to install on Mac

After mounting the installer maybe you would have some issues at the moment to run the installer on Mac since OS blocks the installer app. So you will need to run the `installer.sh` manually. 

### Install manually

1. Open the terminal.

2. Move into `Volumes` directory.
```bash
$ cd  /Volumes/cominnek-$version/
```

3. Run the bash installer
```bash
$ sudo bash installer.sh
```
# Update

You can get the latest version of `cominnek` by running the following command:

```bash
cominnek update
```



# First steps:

Login into your GitHub account

```bash
cominnek auth login
```

Then you can test the connection with the:

```bash
cominnek auth test
```


# Usage

- **[Init](#init)**: Initialize cominnek.
- **[Update version](#update-version)**: Create and push a commit with the correct template for conventional commits.
- **[Push](#push)**: Commit and push the branch to GitHub.
- **[Publish](#publish)**: Commit, push and create the pull request as a draft to develop in GitHub.
- **[Commit](#commit)**: Commit the changes to the branch.
- **[Branch](#branch)**: Create a new branch.
- **[Stash](#stash)**: Stash changes from one branch to another one.
- **[PR](#pr)**: Create a Pull Request as a draft to develop in GitHub.
- **[Merge](#merge)**: Merge the branch into the received branch.
- **[Reset](#reset)**: Reset the branch to the selected commit.
- **[Release](#release-experimental)**: Release a new version. (⚠️ Experimental)
- **[Update](#update)**: Update the cominnek version.
- **[Config](#config)**: Configure the cominnek.
  - **[PR](#pr-1)**: Configure the pull request template.
- **[Auth](#auth)**: Configure the GitHub authentication.
  - **[Login](#login)**: Login into your GitHub account.
  - **[Test](#test)**: Test the connection with GitHub.
  - **[Logout](#logout)**: Logout from your GitHub account.

## Init
Initialize cominnek.

```bash
cominnek init
```
This will create a `.minnekrc.json` file in the root of your project.
## Update version
Create and push a commit with the correct template for conventional commits.

```bash
cominnek update-version <version>
```
the commit will be: `build: update version to <version>`

| flag               | type          | description                 |
| ------------------ | ------------- | --------------------------- |
| `<version>`         |String         | The version to update       |
## Push
Commit and push the branch to GitHub
```bash
cominnek push -m "do some modifications" --fix "home"
```
the commit will be: `fix(home):{Ticket} do some modifications`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-m --message`     |string[]       | Receives the commit message & body message |
| `-M --merge`       |string         | Receives a Branch to merge after end the push|
| `-F --feat`        |string         | make the commit with the prefix feat() |
| `-f --fix`         |string         | make the commit with the prefix fix()  |
| `-d --docs`        |string         | make the commit with the prefix docs() |
| `   --style`       |string         | make the commit with the prefix style()|
| `-r --refactor`    |string         | make the commit with the prefix refactor()|
| `   --perf`        |string         | make the commit with the prefix perf()|
| `   --test`        |string         | make the commit with the prefix test()|
| `-b --build`       |string         | make the commit with the prefix build()|
| `   --ci`          |string         | make the commit with the prefix ci()   |
| `   --chore`       |string         | make the commit with the prefix chore()|
| `   --revert`      |string         | make the commit with the prefix revert()|
| `   --skip-commit` |string         | Skip the commit and only push the branch|


*\* required*

## Publish
Commit, push and create the pull request as a draft to develop in GitHub.
the commit will be: `feat(home):{Ticket} do some modifications`

The usage is the same as [push](#push) just with the difference that this creates a pull request.
```bash
cominnek publish -m "do some modifications" --fix "home"
```
## Commit
Will commit the changes to the branch.
```bash
cominnek commit -m "do some modifications" --fix "home"
```
the commit will be: `feat(home):{Ticket} do some modifications`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-m --message`     |string[]       | Receives the commit message & body message |
| `-a --all`         |boolean        | Add all files to the commit            |
| `-F --feat`        |string         | make the commit with the prefix feat() |
| `-f --fix`         |string         | make the commit with the prefix fix()  |
| `-d --docs`        |string         | make the commit with the prefix docs() |
| `   --style`       |string         | make the commit with the prefix style()|
| `-r --refactor`    |string         | make the commit with the prefix refactor()|
| `   --perf`        |string         | make the commit with the prefix perf()|
| `   --test`        |string         | make the commit with the prefix test()|
| `-b --build`       |string         | make the commit with the prefix build()|
| `   --ci`          |string         | make the commit with the prefix ci()   |
| `   --chore`       |string         | make the commit with the prefix chore()|
| `   --revert`      |string         | make the commit with the prefix revert()|

*\* required*
## Branch
Create a new branch with the prefix `feature/`, `bugfix/`, `hotfix/` or `release/` and the name of the branch will be the ticket number.

We are following a Git-Flow variant called [TAYO](https://mnk-docs.ngrok.io/guide/development/version-control/branch-management.html) by Minnek

```bash
cominnek branch feature "<Ticket>"
```
This the equivalent of: `git branch feature/{Ticket}`

| Command               | description                            |
| ------------------ | ---------------------------------------|
| `feature`       |  create a new feature branch from `develop` |
| `bugfix`        |  create a new bugfix branch from `develop` |
| `hotfix`        |  create a new hotfix branch from `master`|
| `release`       |  create a new release branch from `develop`|
| `support`       |  create a new support branch from `master`|
| `test`          |  create a new test branch from `develop`|
| `sync`          |  create a new sync branch from `develop`|

| Flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<Ticket>`         |string*         |  ticket number |
| `-s --stash`       |boolean        | take the changes in the current branch and apply it to the new feature's branch          |

*\* required*

## Stash
Stash all the changes in the current branch and apply the changes to another branch.

```bash
cominnek stash "<Branch>"
```

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<Branch>`         |string*         |  branch name |
## PR
Create a pull request as a draft directly to develop

```bash
cominnek pr
```
The flag `--ticket` is optional. If it's not provided this will take the ticket number from the current branch.

| flag               | type          | default | description                            |
| ------------------ | ------------- | ------ | ---------------------------------------|
| `-t --ticket`      |string         | take by branch| name of the feature that's will be applied the change     |
| `-b --base`      |string         | `develop` |   base branch of the pull request.  |

In a case of a release branch, it will create a pull request to `master` and `develop`.

## Merge
Merge the current branch into the received one. This will help you save time when you are working on a feature branch and you want to merge, for example, the feature branch into the `test` branch.


```bash
cominnek merge "<branch>"
```


| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<branch>`      |string*         | name of the branch that's will be applied the changes       |

*\* required*

## Reset 
Reset the current branch to the selected commit. By default, it will reset to the last commit.

```bash
cominnek reset
```
This the equivalent of: `git reset --soft HEAD~1`

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-c --commit`      |string         | commit hash       |
| `-y --confirm`     |boolean        | confirm the reset       |
| `   --hard`        |boolean        | Reset HEAD, index and working tree   |
| `   --soft`        |boolean        | Reset only HEAD      |
| `   --mixed`       |boolean        | Reset HEAD and index |
| `   --keep`        |boolean        | Reset HEAD, index and working tree      |
| `   --merge`       |boolean        | Reset HEAD, index and working tree      |
| `-n -number`       |number         | number of commits to reset       |

## Release (Experimental)

⚠️ This is an experimental feature. You can report any issue [here](https://github.com/minnek-digital-studio/cominnek/issues)

Release a new version using [semantic versioning](https://semver.org/)

```bash
cominnek release
```

## Update
Update the Cominnek version.

```bash
cominnek update
```

## Config

Configure the cominnek.
  
  ```bash
  cominnek config -h
  ```
### PR
This command will help you to modify the pull request template.

```bash
cominnek config pr -b path/to/template.md
```

| Flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-b --body`       |string*        | path to the template file |

*\* required*
#### Template file Example:
In some cases, you will want to set some dynamic values, for example, ticket number.

| Key                |    description     |
| ------------------ | ------------------ |
| `${ticket}`        | Ticket number      |
| `${branch}`        | Head branch name   |


```markdown
# Example

- Ticket: ${ticket}
- Branch: ${branch}
```
## Auth

This command will help you to set the GitHub token.

```bash
cominnek auth -h
```

### Login
This command will help you to set the GitHub token.

```bash
cominnek auth login
```

### Test
This command will help you to test the GitHub token.

```bash
cominnek auth test
```

### Logout

This command will help you to remove the GitHub token.

```bash
cominnek auth logout
```

----
## Examples and more

You can add a body to commit using the `--message` flag twice.

```bash
cominnek push -m "Changes in home page" -f "home" -m "the title was aligned to right"
```
This is the same as
`git commit -m "fix(home):{Ticket} Changes in home page" -m "the title was aligned to right"`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

---
To do a commit without a scope, use the flag and on windows, you should use it with a space. 

Mac Os and Linux:
```bash
cominnek push -m "theme setup" -b
```
Windows:

```powershell
cominnek push -m "theme setup" -b " "
```
the commit will be: `build: theme setup`

___ 
Move your changes to a new feature branch

```bash
cominnek feature "<Ticket>" -s
```

# Ticketing system
This takes the ticket number from the branch name. This is important to create the pull request with the correct ticket number.

Example: 

You are in the branch `feature/MJ-11` 

The ticker number will be `MJ-11`

If you execute the command:

```bash
cominnek push -m "Changes on homepage" -F "home"
``` 

the commit will be: `feat(home): MJ-11 Changes on the homepage`

### What about if I am not in a feature branch?
This going to let you know that you are not in a feature branch and you will accept the commit without a ticket number.

Example: You're in the branch `develop`

```bash
$ cominnek push -m "Changes in home page" -F "home"
  This is not a feature. Do you want to continue? (yes or no)
        Commit message: "feat(home): Changes in home page"
```

## Contributing

If you want to contribute to this project, please read the [contributing guide](/CONTRIBUTING.md)

Cominnek `V4.0.0`
> With ❤ by [isaacismaelx14](https://github.com/isaacismaelx14)

## About

<a href="https://minnekdigital.com/">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="https://assets.minnekdigital.com/logo-sm.jpg">
    <img alt="Minnek Logo" src="https://assets.minnekdigital.com/logo-sm.jpg">
  </picture>
</a>

This project is maintained and funded by Minnek.

We ❤️ open source and do our part in sharing our work with the community!
See [our other projects][community] or [hire our team][hire] to help build your product.

Want to join? [Check out our Jobs][jobs]!

[community]: https://github.com/Minnek-Digital-Studio
[hire]: https://minnekdigital.com/
[jobs]: https://minnekdigital.com/careers
