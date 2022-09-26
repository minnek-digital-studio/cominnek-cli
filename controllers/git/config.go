package git_controller

import "fmt"

func Config(name string, email string) string {
	return fmt.Sprintf("git config --global user.name \"%v\" && git config --global user.email \"%v\"", name, email);
}