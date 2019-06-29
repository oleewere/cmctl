package cmd

import (
	"fmt"

	"github.com/oleewere/cmctl/cm"
	"github.com/urfave/cli"
)

// AccountsCommand External account related operations
func AccountsCommand() cli.Command {
	return cli.Command{
		Name:  "accounts",
		Usage: "External account related operations",
		Subcommands: []cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "Print all external accounts",
				Action: func(c *cli.Context) error {
					cmServer := cm.GetActiveCM()
					validateActiveCM(cmServer)
					categories := cmServer.GetExternalAccountCategories()
					var counter int
					for _, category := range categories {
						accountTypes := cmServer.GetExternalAccountTypesByCategory(category)
						for _, accountType := range accountTypes {
							externalAccountData := cmServer.GetExternalAccountsByType(accountType)
							for externalAccount, accountConfigs := range externalAccountData {
								var tableData [][]string
								for configKey, configValue := range accountConfigs {
									tableData = append(tableData, []string{configKey, configValue})
								}
								prefixData := ""
								if counter > 0 {
									prefixData = "\n"
								}
								counter++
								header := fmt.Sprintf("%vExternal Account - %v (type: %v, category: %v):", prefixData, externalAccount, accountType, category)
								printTable(header, []string{"NAME", "VALUE"}, tableData, c)
							}
						}
					}
					return nil
				},
			},
		},
	}
}
