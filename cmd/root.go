/*
Copyright © 2021 Shruti Patel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/go-cmd/cmd"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/viper"
)

var cfgFile string
var is_verbose bool = false
var cfg config

const OPEN_URL_CONFIG_KEY = "open_urls"
const URL_KEY = "url"
const BROWSER_KEY = "browser"
const OPEN_APPS_CONFIG_KEY = "open_apps"
const APP_KEY = "app"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "catapault",
	Short: "Automates setting up different workspaces for different projects.",
	Long: `Automates setting up different workspaces for different projects. 
	Setup inludes opening applications/files/urls/etc. and running shell commands.`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("%+v\n", cfg)
		if is_verbose {
			fmt.Printf("%+v\n", cfg)
		}

		fmt.Printf("Opening urls.\n")
		openUrls(cfg, is_verbose)

		fmt.Printf("Opening apps.\n")
		openApps(cfg, is_verbose)

		fmt.Printf("Running cmds.\n")
		runShellCmds(cfg, is_verbose)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.workspace-setup.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&is_verbose, "verbose", "v", false, "Run in verbose mode (default false)")

	// Config file found and successfully parsed
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		fmt.Printf("Config file passed.\n")
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {

		fmt.Printf("Config file not passed.\n")
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		const Appname = "workspace-setup"
		viper.SetConfigName("config")                             // name of config file (without extension)
		viper.SetConfigType("yaml")                               // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(fmt.Sprintf("/etc/%s", Appname))      // path to look for the config file in
		viper.AddConfigPath(fmt.Sprintf("%s/.%s", home, Appname)) // call multiple times to add many search paths
		viper.AddConfigPath(".")                                  // optionally look for config in the working directory
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(fmt.Errorf("Config file was not found: %w \n", err))
		} else {
			// Config file was found but another error was produced
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
	} else {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	if err := viper.UnmarshalExact(&cfg); err != nil {
		panic(fmt.Errorf("Fatal error unmarshaling config file: %w \n", err))
	}
}

func openUrls(cfg config, is_verbose bool) {
	for _, url_meta := range cfg.Urls {
		fmt.Printf("Opening url: %s\n", url_meta.URL)
		if url_meta.Browser != "" {
			open.StartWith(url_meta.URL, url_meta.Browser)
		} else {
			open.Start(url_meta.URL)
		}
	}
}

func openApps(cfg config, is_verbose bool) {
	for _, app_meta := range cfg.Apps {
		fmt.Printf("Opening App: %s\n", app_meta.App)
		if app_meta.App != "" {
			open.StartWith(app_meta.Args, app_meta.App)
		} else {
			open.Start(app_meta.Args)
		}
	}
}

func runShellCmds(cfg config, is_verbose bool) {
	for _, shell_cmd_meta := range cfg.ShellCmds {
		fmt.Printf("Running cmd: %s\n", shell_cmd_meta.Cmd)
		if len(shell_cmd_meta.Cmd) < 2 {
			panic(fmt.Errorf(""))
		}
		c := cmd.NewCmd(shell_cmd_meta.Cmd[0], shell_cmd_meta.Cmd[1:]...)
		<-c.Start() // Blocks until call is completed
	}
}
