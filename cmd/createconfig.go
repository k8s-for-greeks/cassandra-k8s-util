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
	"io"

	"github.com/k8s-for-greeks/cassandra-k8s-util/pkg/cassandra-k8s-util/conf"
	helper "github.com/k8s-for-greeks/cassandra-k8s-util/pkg/util/cmd"
	"github.com/spf13/cobra"

	"fmt"
	"os"
)

type CreateConfigOptions struct {
	cassandraConfDir  string
	cassandraConfInput string
	jvmConfInput string
	jvmConfOutput string
}

func NewCreateConfigCmd(out io.Writer) *cobra.Command {
	options := &CreateConfigOptions{}
	// createconfigCmd represents the createconfig command
	command := &cobra.Command{
		Use:   "create-config",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCreateConfig(cmd, args, out, options)
			if err != nil {
				helper.ExitWithError(err)
			}
		},
	}

	// TODO add new option values
	command.Flags().StringVarP(&options.cassandraConfDir, "cassandra-conf-dir", "d", options.cassandraConfDir, "code file")
	command.Flags().StringVarP(&options.cassandraConfInput, "cassandra-conf-input", "in", options.cassandraConfInput, "code file")

	return command
}


// TODO trim last "/"

func RunCreateConfig(cmd *cobra.Command, args []string, output io.Writer, options *CreateConfigOptions) error {

	in := "/usr/local/apache-cassandra/conf/cassandra.yaml"
	out := "/etc/conf/cassandra.yaml"
	snitch := "/etc/cassandra/cassandra-rackdc.properties"
	jvmIn := "/usr/local/apache-cassandra/conf/jvm.options"
	jvmOut := "/etc/cassandra/jvm.options"

	if options.cassandraConfInput != "" {
		in = options.cassandraConfInput
	}

	if options.cassandraConfDir != "" {
		out = options.cassandraConfDir + "/cassandra.yaml"
		snitch = options.cassandraConfDir + "/cassandra-rackdc.options"
		jvmOut = options.cassandraConfDir + "/jvm.options"
	}

	c := conf.CreateCassandraConf(in, out)
	c, err := c.ReadCreateCassandraConf()

	if err != nil {
		return fmt.Errorf("unable to read cassandra config: %v", err)
	}

	c, err = c.CreateCassandraConf(os.Environ())

	if err != nil {
		return fmt.Errorf("error creating cassandra config: %v", err)
	}

	if err = c.WriteCassandraConf(); err != nil {
		return fmt.Errorf("error writing cassandra config: %v", err)
	}


	if options.jvmConfInput != "" {
		jvmIn = options.jvmConfInput
	}


	j := conf.CreateJVMOptions(jvmIn,jvmOut)

	heap := os.Getenv("CASSANDRA_MAX_HEAP")
	ringDelay := os.Getenv("CASSANDRA_RING_DELAY")
	migrationWait := os.Getenv("CASSANDRA_MIGRATION_WAIT")

	err = j.WriteJVMOptions(heap,ringDelay,migrationWait)

	if err != nil {
		return fmt.Errorf("error creating jvm options file: %v", err)
	}

	s := conf.CreateGossipSnitchOptions(snitch)

	dc := os.Getenv("CASSANDRA_DC")
	rack := os.Getenv("CASSANDRA_RACK")

	err = s.WriteGossipSnitchOptions(dc, rack)

	if err != nil {
		return fmt.Errorf("error creating snitch file: %v", err)
	}

	return nil

}
