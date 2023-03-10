/*
Copyright © 2022 comwork.io contact.comwork.io
*/
package cmd

import (
	"cwc/cmd/admin"
	"cwc/cmd/bucket"
	"cwc/cmd/configure"
	"cwc/cmd/dnszones"
	"cwc/cmd/environment"
	"cwc/cmd/instance"
	"cwc/cmd/login"
	"cwc/cmd/project"
	"cwc/cmd/provider"
	"cwc/cmd/region"
	"cwc/handlers/user"
	"fmt"

	"cwc/cmd/registry"
	"os"

	"github.com/spf13/cobra"
	"moul.io/banner"
)

var (
	fversion    bool
	cli_version string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cwc",
	Short: "\nA Command Line interface to manage your cloud resources in comwork cloud",
	Long:  "\nA Command Line interface to manage your cloud resources in comwork cloud.\nComplete documentation is available here: https://doc.cloud.comwork.io/docs/tutorials/api/cli/",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(banner.Inline("cwcloud cli"))
		if fversion {
			user.HandleVersion(cli_version)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	cli_version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cwc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	rootCmd.Flags().BoolVarP(&fversion, "version", "v", false, "The cli version")
	rootCmd.AddCommand(admin.AdminCmd)

	rootCmd.AddCommand(project.ProjectCmd)
	rootCmd.AddCommand(bucket.BucketCmd)
	rootCmd.AddCommand(instance.InstanceCmd)
	rootCmd.AddCommand(registry.RegistryCmd)
	rootCmd.AddCommand(login.LoginCmd)
	rootCmd.AddCommand(provider.ProviderCmd)
	rootCmd.AddCommand(environment.EnvironmentCmd)
	rootCmd.AddCommand(region.RegionCmd)
	rootCmd.AddCommand(dnszones.DnsZonesCmd)
	rootCmd.AddCommand(configure.ConfigureCmd)

}
