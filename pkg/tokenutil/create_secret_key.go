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

package tokenutil

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/tokenutil/internal/algorithm"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

func createSecretKey(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a secret key."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	outputToTerminal := cmdutils.Example{
		Desc:    "Create a secret key and print it to the terminal",
		Command: "pulsarctl token create-secret-key",
	}

	outputWithBase64 := cmdutils.Example{
		Desc:    "Create a secret key and print it with base64 encode to the terminal",
		Command: "pulsarctl token create-secret-key --base64",
	}

	outputToFile := cmdutils.Example{
		Desc:    "Create a secret key and output to a file",
		Command: "pulsarctl token create-secret-key --output (file path)",
	}
	examples = append(examples, outputToTerminal, outputWithBase64, outputToFile)
	desc.CommandExamples = examples

	o := make([]byte, 32)

	var out []cmdutils.Output
	toTerminal := cmdutils.Output{
		Desc: "normal output",
		Out:  fmt.Sprintf("%+v", o),
	}

	withBase64 := cmdutils.Output{
		Desc: "Write the secret key to the terminal and encode with base64",
		Out:  base64.StdEncoding.EncodeToString(o),
	}

	toFile := cmdutils.Output{
		Desc: "Write the secret key to a file",
		Out:  "Write secret to file (filename) successfully.",
	}

	invalidSignatureAlgorithmError := cmdutils.Output{
		Desc: "Using invalid signature algorithm to generate secret key",
		Out: "[✖]  the signature algorithm '(signature algorithm)' is invalid. Valid options include: " +
			"'HS256', 'HS384', 'HS512'",
	}

	signatureAlgorithmIsNotHMACError := cmdutils.Output{
		Desc: "Not using HMAC signature algorithm to generate secret key",
		Out: "[✖]  the signature algorithm '(signature algorithm)' is invalid. Valid options include: " +
			"'HS256', 'HS384', 'HS512'",
	}

	out = append(out, toTerminal, withBase64, toFile, invalidSignatureAlgorithmError, signatureAlgorithmIsNotHMACError)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-secret-key",
		"Create a secret key",
		desc.ToString(),
		desc.ExampleToString())

	var signatureAlgorithm string
	var encode bool
	var output string

	vc.SetRunFunc(func() error {
		return doCreateSecretKey(vc, signatureAlgorithm, output, encode)
	})

	vc.FlagSetGroup.InFlagSet("Create secret key", func(set *pflag.FlagSet) {
		set.StringVarP(&signatureAlgorithm, "signature-algorithm", "a", "HS256",
			"The signature algorithm of generate secret key, valid options include 'HS256', 'HS384', 'HS512'")
		set.StringVarP(&output, "output-file", "o", "",
			"The file that the secret key write to")
		set.BoolVarP(&encode, "base64", "b", false,
			"Using base64 to encode the secret key")
	})

}

func doCreateSecretKey(vc *cmdutils.VerbCmd, signatureAlgorithm, outputFile string, encode bool) error {
	sa, err := algorithm.GetSignatureAlgorithm(signatureAlgorithm)
	if err != nil {
		return err
	}
	if !sa.IsHMAC() {
		return errors.Errorf("the signature algorithm '%s' is invalid. Valid options include "+
			"'HS256', 'HS384', 'HS512'\n", signatureAlgorithm)
	}

	secret := sa.GenerateSecret()

	if encode {
		r := base64.StdEncoding.EncodeToString(secret)
		vc.Command.Println(r)
		return nil
	}

	if outputFile != "" {
		err := ioutil.WriteFile(outputFile, secret, 0644)
		if err != nil {
			return errors.Errorf("write the secret key to the file %s failed\n", outputFile)
		}
		vc.Command.Printf("Write the secret key to the file %s successfully.\n", outputFile)
		return nil
	}

	vc.Command.Println(secret)
	return nil
}
