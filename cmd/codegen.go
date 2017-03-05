// Copyright © 2017 NAME HERE K8S For Greeks / Vorstella
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
	"io/ioutil"
	"os"

	"fmt"

	"github.com/ChimeraCoder/gojson"
	"github.com/golang/glog"
	helper "github.com/k8s-for-greeks/cassandra-k8s-util/pkg/util/cmd"
	"github.com/spf13/cobra"
)

type CodeGenOptions struct {
	yamlFile   string
	codePath   string
	structName string
	fileName   string
	pkgName    string
}

func NewCodeGenCmd(out io.Writer) *cobra.Command {

	options := &CodeGenOptions{}
	// codegenCmd represents the codegen command
	command := &cobra.Command{
		Use:   "codegen",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			err := RunCodeGen(cmd, args, out, options)
			if err != nil {
				helper.ExitWithError(err)
			}
		},
	}

	command.Flags().StringVarP(&options.codePath, "code-path", "c", "", "code file")
	command.Flags().StringVarP(&options.yamlFile, "yaml-file", "y", "", "yaml file")
	command.Flags().StringVarP(&options.structName, "struct-name", "s", "", "struct name")
	command.Flags().StringVarP(&options.pkgName, "pkg-name", "p", "", "package name")
	command.Flags().StringVarP(&options.fileName, "file-name", "f", "", "go file name")
	return command
}

func RunCodeGen(cmd *cobra.Command, args []string, out io.Writer, options *CodeGenOptions) error {

	parser := gojson.ParseYaml

	tagList := make([]string, 0)
	tagList = append(tagList, "yaml")
	tagList = append(tagList, "json")

	var input io.Reader
	input = os.Stdin

	if options.yamlFile != "" {
		f, err := os.Open(options.yamlFile)
		if err != nil {
			glog.Fatalf("reading input file: %s", err)
		}
		defer f.Close()
		input = f
	}

	if output, err := gojson.Generate(input, parser, options.structName, options.pkgName, tagList, true); err != nil {
		return fmt.Errorf("error parsing: %v", err)
	} else {
		if options.codePath != "" && options.fileName != "" {
			f := fmt.Sprintf("%s/%s", options.codePath, options.fileName)
			s := fmt.Sprintf("%s\n\n%s", license, string(output[:]))

			err := ioutil.WriteFile(f, []byte(s), 0644)
			if err != nil {
				glog.Fatalf("writing output: %s", err)
				return err
			}
			glog.Infof("Created go file: %q", f)
		} else {
			fmt.Printf("%s", output)
		}
	}

	return nil
}

const license = `//Copyright © 2017 NAME HERE K8S For Greeks / Vorstella
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

////
// Generated code, more details cassandra-k8s-util codegen -h
////`
