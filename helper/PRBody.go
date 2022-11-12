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
	{
		Variable: "${#checks}",
		Value:    "`checks`",
	},
}

var _PRBody string = `
## Issue Info

<a href="https://minnek.atlassian.net/browse/${ticket}" target="_blank">
    <img src="https://img.shields.io/badge/Jira-0052CC?style=for-the-badge&logo=Jira&logoColor=white" alt="JIRA" title="${ticket}"/>
</a>

### Code Review Checklist for Authors

- [ ] Update your branch with the latest changes from ${#develop}.
- [ ] Verify all the ${#checks} are passed (if applicable).
- [ ] Use [code conventions](https://en.wikipedia.org/wiki/Coding_conventions) and [best practices](https://en.wikipedia.org/wiki/Best_practice).
- [ ] New features must have compatibility for a11y (Accessibility).
- [ ] No unintentional ${#clg} left behind after debugging.
- [ ] Notify the code reviewers on time to have an efficient review time.

### Code Review Checklist for Reviewers

- [ ] Take some time to understand the code you are reading.
- [ ] Use an inquisitive tone, do not make an order.
- [ ] Accept that many programming decisions are opinions. Engage a discussion and reach a resolution quickly.
- [ ] Seek to understand the author’s perspective.

**In case something is not applied, justify the reason why you skip one of the points above**
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
