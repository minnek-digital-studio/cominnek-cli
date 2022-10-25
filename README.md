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

Create commits & pull requests easily. `Cominnek` is based on the [Semantic Commit Messages](https://www.conventionalcommits.org/en/v1.0.0/) specification.
## Index
 - **[Requirements](#requirements)**
     - **[Git-Flow (MacOS)](#git-flow-macos)**
 - **[Install](#installation)**
 - **[First Step](#first-steps)**
 - **[Usage](#usage)**
 - **[Ticketing System](#ticketing-system)**

# Requirements

- **[Git](https://git-scm.com/)**
- **[Git-Flow (MacOS)](#git-flow-macos)**

### Git-Flow (MacOS)
To install git flow run: 

```bash
brew install git-flow
```

# Installation

1. Download installer

[![Macos](https://img.shields.io/badge/mac%20os-0078D6?style=for-the-badge&logo=apple&logoColor=white)](https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-2.2.0.dmg) [![Windows](https://img.shields.io/badge/Windows-0078D6?style=for-the-badge&logo=windows&logoColor=white
)](https://github.com/Minnek-Digital-Studio/cominnek/releases/latest/download/cominnek-2.2.0.exe)

2. Run installer

3. Now you can run 
```bash
cominnek -v
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

- **[Update version](#update-version)**: Create and push a commit with the correct template for conventional commits.
- **[Push](#push)**: Commit and push the branch to GitHub.
- **[Publish](#publish)**: Commit, push and create the pull request as a draft to develop in GitHub.
- **[Commit](#commit)**: Commit the changes to the branch.
- **[Flow](#flow)**: Create a new branch and start the flow.
- **[Stash](#stash)**: Stash changes from one branch to another one.
- **[PR](#pr)**: Create a Pull Request as a draft to develop in GitHub.
- **[Merge](#merge)**: Merge the branch into the received branch.
- **[Update](#update)**: Update the cominnek version.
- **[Config](#config)**: Configure the cominnek.
  - **[PR](#pr-1)**: Configure the pull request template.
- **[Auth](#auth)**: Configure the GitHub authentication.
  - **[Login](#login)**: Login into your GitHub account.
  - **[Test](#test)**: Test the connection with GitHub.
  - **[Logout](#logout)**: Logout from your GitHub account.
## Update version
Create and push a commit with the correct template for conventional commits.

```bash
cominnek update-version <version>
```
the commit will be: `build(version): update version`

| flag               | type          | description                 |
| ------------------ | ------------- | --------------------------- |
| `<version>`         |String         | The version to update       |
## Push
Commit and push the branch to GitHub
```bash
cominnek push "do some modifications" --fix "home"
```
the commit will be: `fix(home):{Ticket} do some modifications`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<message>`        |string *         | Commit message                         |
| `-m --body`     |string        | Receives the commit body message            |
| `-M --merge`       |string         | Receives a brach to merge the current branch in the received one  |
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

## Publish
Commit, push and create the pull request as a draft to develop in GitHub.
the commit will be: `feat(home):{Ticket} do some modifications`

The usage is the same as [push](#push) just with the difference that this creates a pull request.
```bash
cominnek publish "do some modifications" --fix "home"
```
## Commit
Will commit the changes to the branch.
```bash
cominnek commit "do some modifications" --fix "home"
```
the commit will be: `feat(home):{Ticket} do some modifications`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<message>`        |string*        | Commit message                         |
| `-m --body`        |string         | Receives the commit body message       |
| `-a --all`         |bolean         | Add all files to the commit            |
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
## Flow
Create a new branch with the prefix `feature/`, `bugfix/`, `hotfix/` or `release/` and the name of the branch will be the ticket number.

```bash
cominnek flow feature "<Ticket>"
```
This the equivalent of: `git flow feature start {Ticket}` or `git-flow feature start {Ticket}` on MacOS

| Command               | description                            |
| ------------------ | ---------------------------------------|
| `feature`       |  create a new feature branch |
| `bugfix`        |  create a new bugfix branch |
| `hotfix`        |  create a new hotfix branch |
| `release`       |  create a new release branch |

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

| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `-t --ticket`      |string         | name of the feature that's will be applied the change     |

## Merge
Merge the current branch into the received one. This will help you save time when you are working on a feature branch and you want to merge, for example, the feature branch into the `test` branch.


```bash
cominnek merge "<branch>"
```


| flag               | type          | description                            |
| ------------------ | ------------- | ---------------------------------------|
| `<branch>`      |string*         | name of the branch that's will be applied the changes       |

*\* required*

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
cominnek push "Changes in home page" -f "home" -m "the title was aligned to right"
```
This is the same as
`git commit -m "fix(home):{Ticket} Changes in home page" -m "the title was aligned to right"`

*{Ticket} is the ticket number* See more information in the [Ticketing system](#ticketing-system)

---
To do a commit without a scope, use the flag and on windows, you should use it with a space. 

Mac Os and Linux:
```bash
cominnek push "theme setup" -b
```
Windows:

```powershell
cominnek push "theme setup" -b " "
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
cominnek push "Changes on homepage" -F "home"
``` 

the commit will be: `feat(home): MJ-11 Changes on the homepage`

### What about if I am not in a feature branch?
This going to let you know that you are not in a feature branch and you will accept the commit without a ticket number.

Example: You're in the branch `develop`

```bash
$ cominnek push "Changes in home page" -F "home"
  This is not a feature. Do you want to continue? (yes or no)
        Commit message: "feat(home): Changes in home page"
```


Cominnek `V2.2.0`
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
