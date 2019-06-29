package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// PlaybookCommands Execute a list of tasks defined in playbook file
func PlaybookCommands() cli.Command {
	return cli.Command{
		Name:  "playbook",
		Usage: "Execute a list of tasks defined in playbook file",
		Action: func(c *cli.Context) error {
			cmServer := cm.GetActiveCM()
			validateActiveCM(cmServer)
			playbookFileParam := c.String("file")
			inventoryFile := c.String("inventory")
			if len(playbookFileParam) == 0 {
				fmt.Println("Provide -f or --file parameter")
				os.Exit(1)
			}
			playbookFiles := strings.Split(playbookFileParam, ",")
			var inventory *cm.Inventory
			if len(inventoryFile) > 0 {
				inventoryVal := cm.ReadInventoryFromFile(inventoryFile)
				inventory = &inventoryVal
			}
			for _, playbookFile := range playbookFiles {
				playbook := cm.LoadPlaybookFile(playbookFile, c.String("vars"))
				cmServer.ExecutePlaybook(playbook, inventory)
			}
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{Name: "file, f", Usage: "Playbook file"},
			cli.StringFlag{Name: "inventory, i", Usage: "Hosts inventory file"},
			cli.StringFlag{Name: "vars, v", Usage: "Provided extra variables (e.g.: --vars='myvar1=myvalue1 myvar2=myvalue2')"},
		},
	}
}
