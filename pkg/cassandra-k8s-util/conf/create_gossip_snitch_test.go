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

func TestWriteGossipSnitch(t *testing.T) {

	f, err := ioutil.TempFile("", "testconf")

	if err != nil {
		t.Fatal(err)
	}

	defer syscall.Unlink(f.Name())
	c := CreateGossipSnitchOptions(f.Name())

	err = c.WriteGossipSnitchOptions("cass", "rack1")
	if err != nil {
		t.Fatal(err)
	}

}
