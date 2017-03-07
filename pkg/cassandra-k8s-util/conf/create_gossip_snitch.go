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

type GossipSnitchOptions struct {
	outputFilename string
}

func CreateGossipSnitchOptions(outputFilename string) *GossipSnitchOptions {
	return &GossipSnitchOptions{
		outputFilename: outputFilename,
	}
}

func (snitch GossipSnitchOptions) WriteGossipSnitchOptions(dc string, rack string) error {

	if dc != "" && rack != "" {

		f, err := os.OpenFile(snitch.outputFilename, os.O_RDWR, 0644)
		if err != nil {
			return fmt.Errorf("Unable to open file: %v", err)
		}

		defer f.Close()

		rack = fmt.Sprintf("rack=%s\n", rack)
		if _, err = f.WriteString(rack); err != nil {
			return fmt.Errorf("Unable to write to file: %v", err)
		}

		dc = fmt.Sprintf("dc=%s\n", rack)
		if _, err = f.WriteString(dc); err != nil {
			return fmt.Errorf("Unable to write to file: %v", err)
		}

	}

	return nil
}
