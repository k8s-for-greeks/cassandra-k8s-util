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
	"syscall"
	"testing"
)

func TestReadCassandraConf(t *testing.T) {

	c := CreateCassandraConf("testdata/cassandra.yaml", "")
	c, err := c.ReadCreateCassandraConf()
	if err != nil {
		t.Fatal(err)
	}

	if c.conf == nil {
		t.Fatal("Conf struct is nil")
	}

}

func TestCreateCassandraConf(t *testing.T) {
	c := CreateCassandraConf("testdata/cassandra.yaml", "")
	c, err := c.ReadCreateCassandraConf()
	if err != nil {
		t.Fatal(err)
	}
	env := []string{
		"FOO",
		"CASSANDRA_SEEDS=\"127.0.0.1,10.0.0.1\"",
		"CASSANDRA_CLUSTER_NAME=foo",
	}
	c, err = c.CreateCassandraConf(env)

	if err != nil {
		t.Fatal(err)
	}

	if c.conf.ClusterName != "foo" {
		t.Fatalf("cluster name does not equal the correct value")
	}
}

func TestWriteCreateCassandraConf(t *testing.T) {

	f, err := ioutil.TempFile("", "testconf")
	if err != nil {
		t.Fatal(err)
	}
	defer syscall.Unlink(f.Name())

	c := CreateCassandraConf("testdata/cassandra.yaml", f.Name())

	c, err = c.ReadCreateCassandraConf()

	if err != nil {
		t.Fatal(err)
	}

	err = c.WriteCassandraConf()

	if err != nil {
		t.Fatal(err)
	}

	c = CreateCassandraConf(f.Name(), "")
	c, err = c.ReadCreateCassandraConf()

	if err != nil {
		t.Fatal(err)
	}

}
