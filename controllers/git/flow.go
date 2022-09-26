package git_controller

func Feature(ticket string) string {
	return "git flow feature start " + ticket
}

func Hotfix(ticket string) string {
	return "git flow hotfix start " + ticket
}

func Release(ticket string) string {
	return "git flow release start " + ticket
}

func Support(ticket string) string {
	return "git flow support start " + ticket
}

func Init() string {
	return "git flow init"
}
