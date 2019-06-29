package cmd

import (
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// SaltCommand salt related commands against CB gateway
func SaltCommand() cli.Command {
	return cli.Command{
		Name:  "salt",
		Usage: "Execute salt commands on the CB gateway",
		Subcommands: []cli.Command{
			{
				Name:    "exec",
				Aliases: []string{"x"},
				Usage:   "Execute salt commands on the CB gateway",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					if !cmServer.UseGateway {
						fmt.Println("Cannot run salt commands as CM Server is not CB managed!")
						os.Exit(1)
					}
					command := c.String("command")
					if len(command) == 0 {
						fmt.Println("Command parameter is missing! (use 'command' or 'c')")
						os.Exit(1)
					}
					cmServer.RunSaltCommand(command, c.String("raw-command"), c.String("prefix"), c.String("binary"))
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "command, c", Usage: "Command to execute on the minions (run cmd)"},
					cli.StringFlag{Name: "raw-command, r", Usage: "Raw command to execute on the minions (e.g.: \"'*' cmd.run 'echo hello'\")"},
					cli.StringFlag{Name: "binary, b", Usage: "Salt binary to use (default: salt)"},
					cli.StringFlag{Name: "prefix, p", Usage: "Salt binary prefix path (default: /opt/salt_*/bin)"},
				},
			},
			{
				Name:  "sync",
				Usage: "Synchronize salt resources from CB source",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					if !cmServer.UseGateway {
						fmt.Println("Cannot run salt commands as CM Server is not CB managed!")
						os.Exit(1)
					}
					if len(c.String("source")) == 0 {
						fmt.Println("The 'source' parameter is required for sync salt scripts!")
						os.Exit(1)
					}
					cmServer.SyncSaltScripts(c.String("source"), c.String("target"), c.String("filter"), c.Bool("clear"))
					return nil
				},
				Flags: []cli.Flag{
					cli.StringFlag{Name: "source, s", Usage: "Source path for the local salt script folder to sync"},
					cli.StringFlag{Name: "target, t", Usage: "Target remote path for the salt scripts (default: /srv)"},
					cli.StringFlag{Name: "filter, f", Usage: "Filter to use only second level folders for uploading (e.g.: pillar)"},
					cli.BoolFlag{Name: "clear", Usage: "Delete salt backup(s) from the gateway"},
				},
			},
		},
	}
}
