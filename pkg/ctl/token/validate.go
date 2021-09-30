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
	"io/ioutil"
	"strings"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/algorithm"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/keypair"

	"github.com/pkg/errors"
	"github.com/spf13/pflag"
)

type validateCmdArgs struct {
	signatureAlgorithm string
	tokenString        string
	tokenFile          string
	secretKeyString    string
	secretKeyFile      string
	publicKeyFile      string
	base64Encoded      bool
}

var errNoTokenStringOrFile = errors.New("both the token string and the token file are not specified")
var errTokenSpecifiedMoreThanOne = errors.New("both the token string and token file are specified")
var errNoValidateKeySpecified = errors.New("none of the validate keys is specified")
var errValidateKeySpecifiedMoreThanOne = errors.New("the validate key is specified more than one")

func validate(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for validating a token."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	validateTokenString := cmdutils.Example{
		Desc:    "Validate a token string using the specified secret key string.",
		Command: "pulsarctl token validate --token-string (token) --secret-key-string (secret-key-string)",
	}

	validateTokenFile := cmdutils.Example{
		Desc:    "Validate a token file using the specified secret key file.",
		Command: "pulsarctl token validate --token-string (token) --secret-key-file (secret-key-file-path)",
	}

	validateTokenStringWithPublicKey := cmdutils.Example{
		Desc:    "Validate a token string using the specified public key file.",
		Command: "pulsarctl token validate --token-string (token) --public-key-file (public-key-file-path)",
	}

	validateTokenFileWithBase64EncodedSecretKeyString := cmdutils.Example{
		Desc:    "Validate a token string using the specified base64 encoded secret key string.",
		Command: "pulsarctl token validate --token-string (token) --secret-key-string (secret-key-string) --base64",
	}

	validateTokenFileWithSignatureAlgorithm := cmdutils.Example{
		Desc: "Validate a token file that signed with the specified secret key string and the specified " +
			"signature algorithm.",
		Command: "pulsarctl toke validate --token-string (token) --secret-key-file (secret-key-file-path) " +
			"--signature-algorithm (algorithm)",
	}
	examples = append(examples, validateTokenString, validateTokenFile, validateTokenStringWithPublicKey,
		validateTokenFileWithBase64EncodedSecretKeyString, validateTokenFileWithSignatureAlgorithm)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	defaultOutput := cmdutils.Output{
		Desc: "The token is valid.",
		Out:  "The subject is (subject), and the expire time is (time).",
	}

	noTokenStringOrFileErr := cmdutils.Output{
		Desc: "Both the token string and the token file are not specified.",
		Out:  "[✖]  " + errNoTokenStringOrFile.Error(),
	}

	tokenSpecifiedMoreThanOneErr := cmdutils.Output{
		Desc: "Both the token string and the token file are specified.",
		Out:  "[✖]  " + errTokenSpecifiedMoreThanOne.Error(),
	}

	noValidateKeySpecifiedErr := cmdutils.Output{
		Desc: "There is no key to validate the token.",
		Out:  "[✖]  " + errNoValidateKeySpecified.Error(),
	}

	validateKeySpecifiedMoreThanOneErr := cmdutils.Output{
		Desc: "The key used to validate token is specified more than one.",
		Out:  "[✖]  " + errValidateKeySpecifiedMoreThanOne.Error(),
	}
	out = append(out, defaultOutput, noTokenStringOrFileErr, tokenSpecifiedMoreThanOneErr,
		noValidateKeySpecifiedErr, validateKeySpecifiedMoreThanOneErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"validate",
		"Validate a token.",
		desc.ToString(),
		desc.ExampleToString())

	args := new(validateCmdArgs)

	vc.SetRunFunc(func() error {
		return doValidate(vc, args)
	})

	vc.FlagSetGroup.InFlagSet("Validate a token", func(set *pflag.FlagSet) {
		set.StringVarP(&args.signatureAlgorithm, "signature-algorithm", "a", "RS256",
			"The signature algorithm is used for generating the token. Valid options are: "+
				"'HS256', 'HS384', 'HS512', 'RS256', 'RS384', 'RS512', 'PS256', 'PS384', 'PS512', 'ES256', "+
				"'ES384', 'ES512'.")
		set.StringVar(&args.tokenString, "token-string", "",
			"The token string that will be validated.")
		set.StringVar(&args.tokenFile, "token-file", "",
			"The token file that will be validated.")
		set.StringVar(&args.secretKeyString, "secret-key-string", "",
			"The secret key string that used to validate a token.")
		set.StringVar(&args.secretKeyFile, "secret-key-file", "",
			"The secret key file that used to validate a token.")
		set.StringVar(&args.publicKeyFile, "public-key-file", "",
			"The public key file that used to validate a token.")
		set.BoolVar(&args.base64Encoded, "base64", false,
			"The secret key is base64 encoded or not.")
	})
	vc.EnableOutputFlagSet()
}

func doValidate(vc *cmdutils.VerbCmd, args *validateCmdArgs) error {
	args = trimSpaceForValidateArgs(args)
	err := validateCmdCheckArgs(args)
	if err != nil {
		return err
	}

	var tokenString string
	switch {
	case args.tokenString != "":
		tokenString = args.tokenString
	case args.tokenFile != "":
		data, err := ioutil.ReadFile(args.tokenFile)
		if err != nil {
			return err
		}
		tokenString = string(data)
	}

	keyData, err := readValidateKeyData(args)
	if err != nil {
		return err
	}

	token := cmdutils.NewPulsarClient().Token()
	subject, expireTime, err := token.Validate(algorithm.Algorithm(args.signatureAlgorithm), tokenString, keyData)
	if err != nil {
		return err
	}

	if expireTime == 0 {
		vc.Command.Printf("The subject is %s and it will never expire.\n", subject)
	} else {
		vc.Command.Printf("The subject is %s, and will expire at %s.\n", subject, time.Unix(expireTime, 0))
	}

	return nil
}

func trimSpaceForValidateArgs(args *validateCmdArgs) *validateCmdArgs {
	return &validateCmdArgs{
		tokenString:        strings.TrimSpace(args.tokenString),
		tokenFile:          strings.TrimSpace(args.tokenFile),
		secretKeyString:    strings.TrimSpace(args.secretKeyString),
		secretKeyFile:      strings.TrimSpace(args.secretKeyFile),
		signatureAlgorithm: strings.TrimSpace(args.signatureAlgorithm),
		publicKeyFile:      strings.TrimSpace(args.publicKeyFile),
	}
}

func validateCmdCheckArgs(args *validateCmdArgs) error {
	if args.tokenString == "" && args.tokenFile == "" {
		return errNoTokenStringOrFile
	}

	if args.tokenString != "" && args.tokenFile != "" {
		return errTokenSpecifiedMoreThanOne
	}

	if args.secretKeyString == "" && args.secretKeyFile == "" && args.publicKeyFile == "" {
		return errNoValidateKeySpecified
	}

	switch {
	case args.secretKeyString != "" && args.secretKeyFile != "":
		fallthrough
	case args.secretKeyString != "" && args.publicKeyFile != "":
		fallthrough
	case args.secretKeyFile != "" && args.publicKeyFile != "":
		return errValidateKeySpecifiedMoreThanOne
	}
	return nil
}

func readValidateKeyData(args *validateCmdArgs) (interface{}, error) {
	switch {
	case args.secretKeyString != "":
		return []byte(args.secretKeyString), nil
	case args.secretKeyString != "" && args.base64Encoded:
		return base64.StdEncoding.DecodeString(args.secretKeyString)
	case args.secretKeyFile != "":
		return ioutil.ReadFile(args.secretKeyFile)
	case args.secretKeyFile != "" && args.base64Encoded:
		data, err := ioutil.ReadFile(args.secretKeyFile)
		if err != nil {
			return nil, err
		}
		return base64.StdEncoding.DecodeString(string(data))
	case args.publicKeyFile != "":
		data, err := ioutil.ReadFile(args.publicKeyFile)
		if err != nil {
			return nil, err
		}
		switch algorithm.Algorithm(args.signatureAlgorithm) {
		case algorithm.RS256:
			fallthrough
		case algorithm.RS384:
			fallthrough
		case algorithm.RS512:
			fallthrough
		case algorithm.PS256:
			fallthrough
		case algorithm.PS384:
			fallthrough
		case algorithm.PS512:
			return keypair.DecodeRSAPublicKey(data)
		case algorithm.ES256:
			fallthrough
		case algorithm.ES384:
			fallthrough
		case algorithm.ES512:
			return keypair.DecodeECDSAPublicKey(data)
		}
	}

	return nil, errors.New("no matching decoder found")
}
