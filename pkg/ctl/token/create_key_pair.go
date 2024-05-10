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
	"os"
	"strings"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/algorithm"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func createKeyPair(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for creating a private and public key pair."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	defaultCreate := cmdutils.Example{
		Desc:    "Create a private and public key pair using RS256 signature algorithm.",
		Command: "pulsarctl token create-key-pair --output-private-key (filepath) --output-public-key (filepath)",
	}

	createWithSignatureAlgorithm := cmdutils.Example{
		Desc: "Create a private and public key pair using the specified signature algorithm.",
		Command: "pulsarctl toke create-key-pair --signature-algorithm (algorithm) --output-private-key (filepath) " +
			"--output-public-key (filepath)",
	}
	examples = append(examples, defaultCreate, createWithSignatureAlgorithm)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	defaultOutput := cmdutils.Output{
		Desc: "Create a key pair successfully.",
		Out: "The private key and public key are generated to (private-key-path) and " +
			"(public-key-path) successfully.",
	}

	writePrivateKeyFailed := cmdutils.Output{
		Desc: "Writing a private key to a file was failed.",
		Out:  "[✖]  failed to write private key to the file (private-key-path)",
	}

	writePublicKeyFailed := cmdutils.Output{
		Desc: "Writing a public key failed to a file was failed.",
		Out:  "[✖]  failed to write public key to the file (public-key-path)",
	}

	keyFilePathEmptyError := cmdutils.Output{
		Desc: "The specified output key file path is empty.",
		Out:  "[✖]  the private key file path and the public key file path can not be empty",
	}
	out = append(out, defaultOutput, writePrivateKeyFailed, writePublicKeyFailed, keyFilePathEmptyError)
	desc.CommandOutput = out

	vc.SetDescription(
		"create-key-pair",
		"Create a private and public key pair.",
		desc.ToString(),
		desc.ExampleToString())

	var signatureAlgorithm string
	var outputPrivateKeyPath string
	var outputPublicKeyPath string

	vc.SetRunFunc(func() error {
		return doCreateKeyPair(vc, signatureAlgorithm, outputPrivateKeyPath, outputPublicKeyPath)
	})

	vc.FlagSetGroup.InFlagSet("Create key pair", func(set *pflag.FlagSet) {
		set.StringVarP(&signatureAlgorithm, "signature-algorithm", "a", "RS256",
			"The signature algorithm is used for generating the key pair. Valid options are: "+
				"'RS256', 'RS384', 'RS512', 'ES256', 'ES384', 'ES512'.")
		set.StringVar(&outputPrivateKeyPath, "output-private-key", "private.key",
			"The file that the private key is written to.")
		set.StringVar(&outputPublicKeyPath, "output-public-key", "public.key",
			"The file that the public key is written to.")
		cobra.MarkFlagRequired(set, "output-private-key")
		cobra.MarkFlagRequired(set, "output-private-key")
	})
	vc.EnableOutputFlagSet()
}

func doCreateKeyPair(vc *cmdutils.VerbCmd, signatureAlgorithm, outputPrivateKeyFilePath,
	outputPublicKeyFilePath string) error {

	if strings.TrimSpace(outputPrivateKeyFilePath) == "" || strings.TrimSpace(outputPublicKeyFilePath) == "" {
		return errors.New("the private key file path and the public key file path can not be empty")
	}

	tokenUtil := cmdutils.NewPulsarClient().Token()
	keyPair, err := tokenUtil.CreateKeyPair(algorithm.Algorithm(signatureAlgorithm))
	if err != nil {
		return err
	}

	privateKey, err := keyPair.EncodedPrivateKey()
	if err != nil {
		return err
	}

	publicKey, err := keyPair.EncodedPublicKey()
	if err != nil {
		return err
	}

	err = os.WriteFile(outputPrivateKeyFilePath, privateKey, 0644)
	if err != nil {
		return errors.WithMessagef(err,
			"failed to write private key to the file %s\n", outputPrivateKeyFilePath)
	}

	err = os.WriteFile(outputPublicKeyFilePath, publicKey, 0644)
	if err != nil {
		return errors.WithMessagef(err,
			"failed to write public key to the file %s\n", outputPublicKeyFilePath)
	}

	vc.Command.Printf("The private key and public key are generated to %s and "+
		"%s successfully.\n", outputPrivateKeyFilePath, outputPublicKeyFilePath)

	return nil
}
