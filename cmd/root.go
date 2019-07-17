// Copyright © 2019 VMware, INC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"

	"github.com/edgexfoundry/edgex-cli/cmd/db"
	"github.com/edgexfoundry/edgex-cli/cmd/device"
	"github.com/edgexfoundry/edgex-cli/cmd/deviceservice"
	"github.com/edgexfoundry/edgex-cli/cmd/event"
	"github.com/edgexfoundry/edgex-cli/cmd/notification"
	"github.com/edgexfoundry/edgex-cli/cmd/profile"
	"github.com/edgexfoundry/edgex-cli/cmd/reading"
	"github.com/edgexfoundry/edgex-cli/cmd/status"
	"github.com/edgexfoundry/edgex-cli/cmd/subscription"
)

// NewCommand returns rootCmd which represents the base command when called without any subcommands
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edgex",
		Short: "EdgeX command line interface",
		Long: " _____    _           __  __  _____                     _            \n" +
			"| ____|__| | __ _  ___\\ \\/ / |  ___|__  _   _ _ __   __| |_ __ _   _ \n" +
			"|  _| / _` |/ _` |/ _ \\\\  /  | |_ / _ \\| | | | '_ \\ / _` | '__| | | |\n" +
			"| |__| (_| | (_| |  __//  \\  |  _| (_) | |_| | | | | (_| | |  | |_| |\n" +
			"|_____\\__,_|\\__, |\\___/_/\\_\\ |_|  \\___/ \\__,_|_| |_|\\__,_|_|   \\__, |\n" +
			"            |___/                                              |___/ \n" +
			`

https://www.edgexfoundry.org/
	`,
	}

	// Add all subcommands below:
	cmd.AddCommand(device.NewCommand())
	cmd.AddCommand(deviceservice.NewCommand())
	cmd.AddCommand(profile.NewCommand())
	cmd.AddCommand(event.NewCommand())
	cmd.AddCommand(reading.NewCommand())
	cmd.AddCommand(status.NewCommand())
	cmd.AddCommand(db.NewCommand())
	cmd.AddCommand(notification.NewCommand())
	cmd.AddCommand(subscription.NewCommand())

	// global flags
	Verbose := false
	cmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Print URL(s) used by the entered command.")

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	// set default config
	setConfig()
	if err := NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func setConfig() {
	viper.SetDefault("Host", "localhost")
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/.edgex-cli")
	viper.WriteConfig()
	viper.SafeWriteConfig()
}