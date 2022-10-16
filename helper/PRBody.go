package helper

import "strings"

type Variables struct {
	Variable string
	Value    string
}

var _variables []Variables = []Variables{
	{
		Variable: "${#develop}",
		Value:    "`develop`",
	},
	{
		Variable: "${#rem}",
		Value:    "`rem`",
	},
	{
		Variable: "${#px}",
		Value:    "`px`",
	},
	{
		Variable: "${#lang}",
		Value:    "`lang`",
	},
	{
		Variable: "${#clg}",
		Value:    "`console.logs`",
	},
	{
		Variable: "${#feature-ticket}",
		Value:    "Feature/Ticket-ID",
	},
	{
		Variable: "${#developer}",
		Value:    "`developer`",
	},
}

var _PRBody string = `
# Ticket information: 

[![Jira](https://img.shields.io/badge/Jira-0052CC?style=for-the-badge&logo=Jira&logoColor=white)](https://minnek.atlassian.net/browse/${ticket})

- ${ticket}

### Code Review Checklist

- [ ] Update your branch with latest changes from ${#develop}
- [ ] Verify all the checks are passed (if applicable)
- [ ] Use clean code practices
- [ ] Everything sorted alphabetically
- [ ] Use ${#rem} not ${#px}
- [ ] Instructions for how reviewers can test the code locally
- [ ] Screenshot of the feature/bug fix (if applicable)
- [ ] If any new text is added, it's internationalized, use ${#lang} function
- [ ] Any new elements have aria labels (accessibility)
- [ ] No unintentional ${#clg} left behind after debugging
- [ ] Did I use the clear and concise names for variables and functions?
- [ ] Did I explain all possible solutions and why I chose the one I did?
- [ ] Added any comments to make new functions clearer
- [ ] Comment the code if it's really needed (the code show be self-explained)
- [ ] Added PR labels
- [ ] Update any history/changelog file

>**In case something is not applied, justify the reason why you skip one of the points above**

#### Notes before Merge

- Change pull request title to ${#feature-ticket} to match the branch name. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/git-version-control.md#pull-request).
- Write Short and Detailed Commit Messages. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/git-version-control.md#writing-commits).
- Set a ${#developer}, to Code Review this pull request before merge to the ${#develop} branch. See the docs [here](https://github.com/Minnek-Digital-Studio/minnek-developer-handbook/blob/master/development/code-review.md).
- Follow these Guides for Coding Standard: [SASS](https://github.com/bigcommerce/sass-style-guide), [JavaScript](https://developer.mozilla.org/en-US/docs/MDN/Guidelines/Code_guidelines/JavaScript) and [HTML5](https://developer.mozilla.org/en-US/docs/MDN/Guidelines/Code_guidelines/HTML#class_and_id_names).
`

func ReplaceValues(base string, origin []Variables) string {
	for _, v := range origin {
		if strings.Contains(base, v.Variable) {
			base = strings.ReplaceAll(base, v.Variable, v.Value)
		}
	}

	return base
}

var PRBody = ReplaceValues(_PRBody, _variables)
