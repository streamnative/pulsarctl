// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package token

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/algorithm"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func createSecretKey(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a secret key."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	outputToTerminal := cmdutils.Example{
		Desc:    "Create a secret key.",
		Command: "pulsarctl token create-secret-key",
	}

	outputWithBase64 := cmdutils.Example{
		Desc:    "Create a base64 encoded secret key.",
		Command: "pulsarctl token create-secret-key --base64",
	}

	outputToFile := cmdutils.Example{
		Desc:    "Create a secret key and save it to a file.",
		Command: "pulsarctl token create-secret-key --output (file path)",
	}
	examples = append(examples, outputToTerminal, outputWithBase64, outputToFile)
	desc.CommandExamples = examples

	o := make([]byte, 32)

	var out []cmdutils.Output
	toTerminal := cmdutils.Output{
		Desc: "Create a secret key successfully.",
		Out:  fmt.Sprintf("%+v", o),
	}

	withBase64 := cmdutils.Output{
		Desc: "Write a base64 encoded secret key to the terminal.",
		Out:  base64.StdEncoding.EncodeToString(o),
	}

	toFile := cmdutils.Output{
		Desc: "Write the secret key to a file successfully.",
		Out:  "Write secret to the file (filename) successfully.",
	}

	toFileError := cmdutils.Output{
		Desc: "Writing the secret key to a file was failed.",
		Out:  "[✖]  writing the secret key to the file (filename) was failed",
	}

	invalidSignatureAlgorithmError := cmdutils.Output{
		Desc: "Using invalid signature algorithm to generate secret key.",
		Out: "[✖]  the signature algorithm '(signature algorithm)' is invalid. Valid options are: " +
			"'HS256', 'HS384', 'HS512'",
	}

	out = append(out, toTerminal, withBase64, toFile, toFileError, invalidSignatureAlgorithmError)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-secret-key",
		"Create a secret key",
		desc.ToString(),
		desc.ExampleToString())

	var signatureAlgorithm string
	var base64Encoded bool
	var output string

	vc.SetRunFunc(func() error {
		return doCreateSecretKey(vc, signatureAlgorithm, output, base64Encoded)
	})

	vc.FlagSetGroup.InFlagSet("Create secret key", func(set *pflag.FlagSet) {
		set.StringVarP(&signatureAlgorithm, "signature-algorithm", "a", "HS256",
			"The signature algorithm used for generating the secret key. Valid options are:"+
				"'HS256', 'HS384', 'HS512'.")
		set.StringVarP(&output, "output-file", "o", "",
			"The file that the secret key is written to.")
		set.BoolVarP(&base64Encoded, "base64", "b", false,
			"Generate a base64 encoded secret key.")
	})

}

func doCreateSecretKey(vc *cmdutils.VerbCmd, signatureAlgorithm, outputFile string, base64Encoded bool) error {
	admin := cmdutils.NewPulsarClient()
	secret, err := admin.Token().CreateSecretKey(algorithm.Algorithm(signatureAlgorithm))
	if err != nil {
		return err
	}

	var output []byte
	if base64Encoded {
		output = []byte(base64.StdEncoding.EncodeToString(secret))
	} else {
		output = secret
	}

	if outputFile != "" {
		err := ioutil.WriteFile(outputFile, output, 0644)
		if err != nil {
			return errors.Errorf("writing the secret key to the file %s was failed\n", outputFile)
		}
		vc.Command.Printf("Write the secret key to the file %s successfully\n", outputFile)
		return nil
	}

	vc.Command.Println(string(output))
	return nil
}
