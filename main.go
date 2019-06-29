package main

import (
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cmd"
	"github.com/urfave/cli"
)

// Version that will be generated during the build as a constant
var Version string

// GitRevString that will be generated during the build as a constant - represents git revision value
var GitRevString string

func main() {
	app := cli.NewApp()
	app.Name = "cmctl"
	app.Usage = "CLI tool for handle CM clusters"
	app.EnableBashCompletion = true
	app.UsageText = "cmctl command [command options] [arguments...]"
	if len(Version) > 0 {
		app.Version = Version
	} else {
		app.Version = "0.1.0"
	}
	if len(GitRevString) > 0 {
		app.Version = app.Version + fmt.Sprintf(" (git short hash: %v)", GitRevString)
	}

	app.Commands = []cli.Command{}

	app.Commands = append(app.Commands, cmd.ServersCommand())
	app.Commands = append(app.Commands, cmd.ProfilesCommand())
	app.Commands = append(app.Commands, cmd.ClustersCommand())
	app.Commands = append(app.Commands, cmd.HostsCommand())
	app.Commands = append(app.Commands, cmd.ServicesCommand())
	app.Commands = append(app.Commands, cmd.RolesCommand())
	app.Commands = append(app.Commands, cmd.UsersCommand())
	app.Commands = append(app.Commands, cmd.AccountsCommand())
	app.Commands = append(app.Commands, cmd.ConfigsCommand())
	app.Commands = append(app.Commands, cmd.InventoryCommand())
	app.Commands = append(app.Commands, cmd.SaltCommand())
	app.Commands = append(app.Commands, cmd.ExecCommand())
	app.Commands = append(app.Commands, cmd.PlaybookCommands())

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
