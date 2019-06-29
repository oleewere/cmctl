package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// UsersCommand user related operations
func UsersCommand() cli.Command {
	return cli.Command{
		Name:  "users",
		Usage: "User related opreations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all users per cluster",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					users := cmServer.GetUsers()
					if len(users) > 0 {
						var tableData [][]string
						for _, user := range users {
							authRoles := user.AuthRoles
							authRolesStrList := make([]string, 0)
							for _, authRole := range authRoles {
								authRolesStrList = append(authRolesStrList, authRole.Name)
							}
							authRolesStr := strings.Join(authRolesStrList[:], ",")
							tableData = append(tableData, []string{user.Name, authRolesStr})
						}
						header := "USERS: "
						printTable(header, []string{"NAME", "AUTH ROLES"}, tableData, c)
					} else {
						fmt.Println("Not found users!")
						os.Exit(1)
					}

					return nil
				},
			},
		},
	}
}
