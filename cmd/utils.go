package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/oleewere/cmctl/cm"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

func printTable(title string, headers []string, data [][]string, c *cli.Context) {
	fmt.Println(title)
	if len(data) > 0 {
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader(headers)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	} else {
		for i := 1; i <= len(title); i++ {
			fmt.Print("-")
		}
		fmt.Println()
		fmt.Println("NO ENTRIES FOUND!")
	}
}

func printJSON(b []byte) {
	fmt.Println(formatJSON(b).String())
}

func formatJSON(b []byte) *bytes.Buffer {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &out
}

func validateActiveCM(cmServer cm.CMServer) {
	if len(cmServer.Name) == 0 {
		fmt.Println("No active CM server selected. (see 'servers use' command)")
		os.Exit(1)
	} else {
		if cmServer.UseGateway && len(cmServer.ConnectionProfile) == 0 {
			fmt.Println("No connection profile attached to CM server/gateway. (see 'profiles' command)")
			os.Exit(1)
		}
	}
}
