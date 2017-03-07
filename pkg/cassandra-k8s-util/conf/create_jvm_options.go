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
	"fmt"
	"os"
)

type JVMOptions struct {
	inputFilename  string
	outputFilename string
}

func CreateJVMOptions(inputFilename string, outputFilename string) *JVMOptions {
	return &JVMOptions{
		inputFilename:  inputFilename,
		outputFilename: outputFilename,
	}
}

func (jvm JVMOptions) WriteJVMOptions(maxHeap string, ringDelay string, migrationWait string) error {

	if jvm.inputFilename == "" {
		return fmt.Errorf("inputFilename cannot be empty")
	}

	f, err := os.OpenFile(jvm.inputFilename, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("Unable to open file: %v", err)
	}

	defer f.Close()

	if migrationWait == "" {
		migrationWait = "-1"
	}

	s := fmt.Sprintf("-Dcassandra.migration_task_wait_in_seconds=%s\n", migrationWait)
	if _, err = f.WriteString(s); err != nil {
		return fmt.Errorf("Unable to write to file: %v", err)
	}

	if ringDelay == "" {
		ringDelay = "30000"
	}

	s = fmt.Sprintf("-D-Dcassandra.ring_delay_ms=%s\n", ringDelay)
	if _, err = f.WriteString(s); err != nil {
		return fmt.Errorf("Unable to write to file: %v", err)
	}

	if maxHeap != "" {

		s := fmt.Sprintf("-Xms%s\n", maxHeap)
		if _, err = f.WriteString(s); err != nil {
			return fmt.Errorf("Unable to write to file: %v", err)
		}

		s = fmt.Sprintf("-Xmx%s\n", maxHeap)
		if _, err = f.WriteString(s); err != nil {
			return fmt.Errorf("Unable to write to file: %v", err)
		}
	}

	return nil
}
