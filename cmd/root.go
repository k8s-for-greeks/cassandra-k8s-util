// Copyright © 2017 K8s For Greeks / Vorstella
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
	goflag "flag"
	helper "github.com/k8s-for-greeks/cassandra-k8s-util/pkg/util/cmd"
	"github.com/spf13/cobra"
	"io"
	"os"
)

// FIXME put this into a struct
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cassandra-k8s-util",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	goflag.Set("logtostderr", "true")
	goflag.CommandLine.Parse([]string{})
	if err := RootCmd.Execute(); err != nil {
		helper.ExitWithError(err)
	}
}

func init() {

	NewCmdRoot(os.Stdout)
}

func NewCmdRoot(out io.Writer) *cobra.Command {

	RootCmd.AddCommand(NewCodeGenCmd(out))
	RootCmd.AddCommand(NewCreateConfigCmd(out))

	return RootCmd
}
