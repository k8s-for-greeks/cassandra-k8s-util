//Copyright © 2017 NAME HERE K8S For Greeks / Vorstella
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

package cass

import (
	"io/ioutil"
	"reflect"
	"strconv"

	"fmt"

	"strings"

	"github.com/golang/glog"
	yaml "gopkg.in/yaml.v2"
	"github.com/serenize/snaker"
)

type CassandraConf struct {
	conf           *CassandraYaml
	inputFilename  string
	outputFilename string
}

func CreateCassandraConf(inputFilename string) *CassandraConf {

	return &CassandraConf{
		conf:           &CassandraYaml{},
		inputFilename:  inputFilename,
		outputFilename: "/etc/cassandra/conf/",
	}
}

func (cass CassandraConf) ReadCreateCassandraConf() (*CassandraConf, error) {

	if cass.inputFilename == "" {
		return nil, fmt.Errorf("inputFilename cannot be empty")
	}

	b, err := ioutil.ReadFile(cass.inputFilename)
	if err != nil {
		return nil, fmt.Errorf("reading input file: %q - %s", cass.inputFilename, err)
	}

	yaml.Unmarshal(b, &cass.conf)

	if err != nil {
		return nil, fmt.Errorf("reading input file: %q - %s", cass.inputFilename, err)
	}

	return &cass, nil
}

func (cass CassandraConf) CreateCassandraConf(env []string) (*CassandraConf, error) {
	for _, e := range env {

		if !strings.HasPrefix(e, "CASSANDRA_") || e == "CASSANDRA_DC" || e == "CASSANDRA_RACK" {
			continue
		}

		a := strings.Split(e, "=")
		key := strings.ToLower(strings.Replace(a[0], "CASSANDRA_", "", -1))

		value := a[1]

		if key == "seeds" {
			cass.conf.SeedProvider[0].Parameters[0].Seeds = value
			continue
		}


		key = snaker.SnakeToCamel(key)

		v := reflect.ValueOf(cass.conf).Elem().FieldByName(key)

		if v.IsValid() {
			if v.Kind() == reflect.Int {

				i, err := strconv.Atoi(value)
				if err != nil {
					glog.Errorf("unable to set %q as int", value)
				}
				v.SetInt(int64(i))

			} else if v.Kind() == reflect.Bool {
				b, err := strconv.ParseBool(value)
				if err != nil {
					glog.Errorf("unable to set %q as bool", value)
				}
				v.SetBool(b)
			} else if v.Kind() == reflect.String {
				v.SetString(value)
			} else if v.Kind() == reflect.Struct {
				return nil, fmt.Errorf("struct is not supported yet")
			}
		}

	}

	dir := "/var/lib/cassandra/"

	if len(cass.conf.DataFileDirectories) == 0 {
		cass.conf.DataFileDirectories = []string { dir + "data" }
	}
	if cass.conf.CommitlogDirectory == "" {
		cass.conf.DataFileDirectories = []string { dir + "commitlog" }
	}

	if cass.conf.CdcRawDirectory == "" {
		cass.conf.DataFileDirectories = []string { dir + "cdc_raw" }
	}
	if cass.conf.SavedCachesDirectory == "" {
		cass.conf.DataFileDirectories = []string { dir + "saved_caches" }
	}
	if cass.conf.HintsDirectory == "" {
		cass.conf.DataFileDirectories = []string { dir + "hints" }
	}


	return &cass, nil
}

func (cass CassandraConf) WriteCassandraConf() error {
	b, err := yaml.Marshal(cass.conf)

	if err != nil {
		return fmt.Errorf("unable to marshal cassandra yaml: %v", err)
	}

	err = ioutil.WriteFile(cass.outputFilename, b, 0644)
	if err != nil {
		glog.Fatalf("writing output: %s", err)
		return err
	}
	return nil
}
