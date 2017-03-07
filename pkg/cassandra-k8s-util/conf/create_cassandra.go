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

package conf

import (
	"io/ioutil"
	"reflect"
	"strconv"

	"fmt"

	"strings"

	"github.com/golang/glog"
	"github.com/serenize/snaker"
	yaml "gopkg.in/yaml.v2"
)

type CassandraConf struct {
	conf           *CassandraYaml
	inputFilename  string
	outputFilename string
}

func CreateCassandraConf(inputFilename string, outputFilename string) *CassandraConf {
	return &CassandraConf{
		conf:           &CassandraYaml{},
		inputFilename:  inputFilename,
		outputFilename: outputFilename,
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

// TODO - this function is messy ... how do we make it better?

func (cass CassandraConf) CreateCassandraConf(env []string) (*CassandraConf, error) {

	// TODO migration wait and ring delay into jvm

	var setNumTokens, setAutoBootstrap = true, true

	for _, e := range env {
		if strings.HasPrefix(e, "POD_IP") {

			a := strings.Split(e, "=")

			v := reflect.ValueOf(cass.conf).Elem().FieldByName("ListenAddress")
			if v.IsValid() {
				v.SetString(a[1])
			}

		}

		if !strings.HasPrefix(e, "CASSANDRA_") || e == "CASSANDRA_DC" || e == "CASSANDRA_RACK" {
			continue
		}

		a := strings.Split(e, "=")
		key := strings.ToLower(strings.Replace(a[0], "CASSANDRA_", "", -1))

		if key == "NUM_TOKENS" {
			setNumTokens = false
		} else if key == "AUTO_BOOTSTRAP" {
			setAutoBootstrap = false
		}

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
			} else {
				return nil, fmt.Errorf("not sure how we got here, but we should have not, we did not find the type.")
			}
		} else {
			glog.Warningf("unable to find %q field in cassandraYaml struct", key)
		}

	}

	// We are hard coding this regardless
	v := reflect.ValueOf(cass.conf).Elem().FieldByName("RPCAddress")
	if v.IsValid() {
		v.SetString("0.0.0.0")
	} else {
		glog.Warningf("unable to find RPCAddress field in cassandraYaml struct")
	}

	if setNumTokens {
		v := reflect.ValueOf(cass.conf).Elem().FieldByName("NumTokens")
		if v.IsValid() {
			v.SetInt(int64(32))
		}
	} else {
		glog.Warningf("unable to find NumTokens field in cassandraYaml struct")
	}

	if setAutoBootstrap {
		v := reflect.ValueOf(cass.conf).Elem().FieldByName("AutoBootstrap")
		if v.IsValid() {
			v.SetBool(false)
		}

	} else {
		glog.Warningf("unable to find AutoBootstrap field in AutoBootstrap struct")
	}

	dir := "/var/lib/cassandra/"

	if len(cass.conf.DataFileDirectories) == 0 {
		cass.conf.DataFileDirectories = []string{dir + "data"}
	}
	if cass.conf.CommitlogDirectory == "" {
		cass.conf.DataFileDirectories = []string{dir + "commitlog"}
	}

	if cass.conf.CdcRawDirectory == "" {
		cass.conf.DataFileDirectories = []string{dir + "cdc_raw"}
	}
	if cass.conf.SavedCachesDirectory == "" {
		cass.conf.DataFileDirectories = []string{dir + "saved_caches"}
	}
	if cass.conf.HintsDirectory == "" {
		cass.conf.DataFileDirectories = []string{dir + "hints"}
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
