package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"github.com/FINTprosjektet/fint-model/github"
)

func CmdListBranches(c *cli.Context) {
	for _, b := range github.GetBranchList() {
		fmt.Println(b)
	}
}