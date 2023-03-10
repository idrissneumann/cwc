/*
Copyright © 2022 comwork.io contact.comwork.io

*/
package ls

import (
	"cwc/handlers/admin"

	"github.com/spf13/cobra"
)

var (
	instanceId string
)

// lsCmd represents the ls command
var LsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List availble instances in the cloud",
	Long: `This command lets you list your available instances in the cloud
This command takes no arguments`,
	Run: func(cmd *cobra.Command, args []string) {
		if *&instanceId == "" {

			admin.HandleGetInstances()
		} else {
			admin.HandleGetInstance(&instanceId)
		}

	},
}

func init() {

	LsCmd.Flags().StringVarP(&instanceId, "instance", "i", "", "The instance id")

}
