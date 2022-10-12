package git

import (
	"fmt"
	"log"

	git_controller "github.com/Minnek-Digital-Studio/cominnek/controllers/git"
	"github.com/Minnek-Digital-Studio/cominnek/pkg/shell"
)

func Config(name string, email string) {
	cmd := git_controller.Config(name, email);

	err, out, errout := shell.Out(cmd)
	
	if err != nil {
		fmt.Println(out);
		fmt.Println(errout);
		log.Fatal(errout);
	}

	fmt.Println(out);
	fmt.Println(errout);
}
