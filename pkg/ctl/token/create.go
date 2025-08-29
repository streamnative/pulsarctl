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
	"os"
	"strings"
	"time"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/algorithm"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/keypair"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type createCmdArgs struct {
	base64Encoded      bool
	signatureAlgorithm string
	subject            string
	expireTime         string
	headers            map[string]string
	secretKeyString    string
	secretKeyFile      string
	privateKeyFile     string
}

var errNoKeySpecified = errors.New("none of the signing keys is specified")
var errKeySpecifiedMoreThanOne = errors.New("the signing key is specified more than one")

func create(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for create a token string."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	createTokenWithSecretKeyString := cmdutils.Example{
		Desc:    "Create a token using a secret key string.",
		Command: "pulsarctl token create --secret-key-string (secret-key-string) --subject (subject)",
	}

	createTokenWithSecretKeyFile := cmdutils.Example{
		Desc:    "Create a token using a secret key file.",
		Command: "pulsarctl token create --secret-key-file (secret-key-file-path) --subject (subject)",
	}

	createTokenWithPrivateKeyFile := cmdutils.Example{
		Desc:    "Create a token using a private key file.",
		Command: "pulsarctl token create --private-key-file (private-key-file-path) --subject (subject)",
	}

	createTokenWithExpireTime := cmdutils.Example{
		Desc:    "Create a token with expire time.",
		Command: "pulsarctl token create --secret-key-string (secret-key-string) --subject (subject) --expire 1m",
	}

	createTokenWithHeaders := cmdutils.Example{
		Desc: "Create a token with headers.",
		Command: "pulsarctl token create --secret-key-string (secret-key-string) --subject (subject)" +
			" -headers kid=kid1,key2=value2",
	}

	createTokenWithBase64EncodedSecretKeyString := cmdutils.Example{
		Desc:    "Create a token using a base64 encoded secret key.",
		Command: "pulsarctl token create --secret-key-string (secret-key-string) --base64 --subject (subject)",
	}
	examples = append(examples, createTokenWithSecretKeyString, createTokenWithSecretKeyFile,
		createTokenWithPrivateKeyFile, createTokenWithExpireTime, createTokenWithBase64EncodedSecretKeyString,
		createTokenWithHeaders)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	defaultOutput := cmdutils.Output{
		Desc: "Create a token successfully.",
		Out:  "eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJoZWxsby10ZXN0In0.qxaczygeZaZDlK7jQHHXCaQRbwd2wxIHjCH3y_Lo2Q4",
	}

	keysNotSpecifiedErr := cmdutils.Output{
		Desc: "None of the signing keys is specified.",
		Out:  "[✖]  " + errNoKeySpecified.Error(),
	}

	KeySpecifiedMoreThanOneErr := cmdutils.Output{
		Desc: "Signing key is specified more than one.",
		Out:  "[✖]  " + errKeySpecifiedMoreThanOne.Error(),
	}
	out = append(out, defaultOutput, keysNotSpecifiedErr, KeySpecifiedMoreThanOneErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"create",
		"Create a token string.",
		desc.ToString(),
		desc.ExampleToString())

	args := new(createCmdArgs)

	vc.SetRunFunc(func() error {
		return doCreate(vc, args)
	})

	vc.FlagSetGroup.InFlagSet("Create a token", func(set *pflag.FlagSet) {
		set.StringVarP(&args.signatureAlgorithm, "signature-algorithm", "a", "RS256",
			"The signature algorithm used to generate the secret key or the private key "+
				"Valid options are: 'HS256', 'HS384', 'HS512', 'RS256', 'RS384', 'RS512', 'PS256', "+
				"'PS384', 'PS512', 'ES256', 'ES384', 'ES512'.")
		set.StringVar(&args.secretKeyString, "secret-key-string", "",
			"The secret key string that used to sign a token.")
		set.StringVar(&args.secretKeyFile, "secret-key-file", "",
			"The secret key file that used to sign a token.")
		set.StringVar(&args.privateKeyFile, "private-key-file", "",
			"The private key file that used to sign a toke.")
		set.StringVar(&args.subject, "subject", "",
			"The 'subject' or 'principal' associate with this token.")
		set.StringVar(&args.expireTime, "expire", "",
			"The expire time for a token. e.g. 1s, 1m, 1h")
		set.BoolVar(&args.base64Encoded, "base64", false,
			"The secret key is base64 encoded or not.")
		set.StringToStringVar(&args.headers, "headers", nil,
			"The headers for a token. e.g. key1=value1,key2=value2")
		_ = cobra.MarkFlagRequired(set, "subject")
	})
	vc.EnableOutputFlagSet()
}

func doCreate(vc *cmdutils.VerbCmd, args *createCmdArgs) error {
	args = trimSpaceArgs(args)
	err := createCmdCheckArgs(args)
	if err != nil {
		return err
	}
	signKey, err := parseSigningKeyData(args)
	if err != nil {
		return err
	}

	token := cmdutils.NewPulsarClient().Token()

	var expireTime int64
	if args.expireTime != "" {
		d, err := time.ParseDuration(args.expireTime)
		if err != nil {
			return err
		}
		expireTime = time.Now().Add(d).Unix()
	}

	var claims *jwt.MapClaims
	if expireTime <= 0 {
		claims = &jwt.MapClaims{
			"sub": args.subject,
		}
	} else {
		claims = &jwt.MapClaims{
			"sub": args.subject,
			"exp": jwt.NewNumericDate(time.Unix(expireTime, 0)),
		}
	}

	// Covert headers to map[string]interface{}
	headers := make(map[string]interface{})
	if len(args.headers) > 0 {
		for key, value := range args.headers {
			headers[key] = value
		}
	}

	tokenString, err := token.CreateToken(algorithm.Algorithm(args.signatureAlgorithm), signKey, claims, headers)
	if err != nil {
		return err
	}
	vc.Command.Println(tokenString)

	return nil
}

func createCmdCheckArgs(args *createCmdArgs) error {
	if args.secretKeyString == "" && args.secretKeyFile == "" && args.privateKeyFile == "" {
		return errNoKeySpecified
	}

	switch {
	case args.secretKeyFile != "" && args.secretKeyString != "":
		fallthrough
	case args.secretKeyFile != "" && args.privateKeyFile != "":
		fallthrough
	case args.secretKeyString != "" && args.privateKeyFile != "":
		return errKeySpecifiedMoreThanOne
	default:
		return nil
	}
}

func trimSpaceArgs(args *createCmdArgs) *createCmdArgs {
	return &createCmdArgs{
		signatureAlgorithm: strings.TrimSpace(args.signatureAlgorithm),
		subject:            strings.TrimSpace(args.subject),
		expireTime:         strings.TrimSpace(args.expireTime),
		secretKeyString:    strings.TrimSpace(args.secretKeyString),
		secretKeyFile:      strings.TrimSpace(args.secretKeyFile),
		privateKeyFile:     strings.TrimSpace(args.privateKeyFile),
		headers:            args.headers,
	}
}

func parseSigningKeyData(args *createCmdArgs) (interface{}, error) {
	switch {
	case args.secretKeyString != "":
		return []byte(args.secretKeyString), nil
	case args.secretKeyString != "" && args.base64Encoded:
		return base64.StdEncoding.DecodeString(args.secretKeyString)
	case args.secretKeyFile != "":
		return os.ReadFile(args.secretKeyFile)
	case args.secretKeyFile != "" && args.base64Encoded:
		data, err := os.ReadFile(args.secretKeyFile)
		if err != nil {
			return nil, err
		}
		return base64.StdEncoding.DecodeString(string(data))
	case args.privateKeyFile != "":
		data, err := os.ReadFile(args.privateKeyFile)
		if err != nil {
			return nil, err
		}
		switch {
		case strings.HasPrefix(args.signatureAlgorithm, "RS") || strings.HasPrefix(args.signatureAlgorithm, "PS"):
			kp, err := keypair.DecodePrivateKey(keypair.RSA, data)
			if err != nil {
				return nil, err
			}
			return kp.GetRsaPrivateKey()
		case strings.HasPrefix(args.signatureAlgorithm, "ES"):
			kp, err := keypair.DecodePrivateKey(keypair.ECDSA, data)
			if err != nil {
				return nil, err
			}
			return kp.GetEcdsaPrivateKey()
		case strings.HasPrefix(args.signatureAlgorithm, "HS"):
			return nil, errors.New("invalid type of the signature algorithm")
		}
	}

	return nil, errors.New("no way to decode the signature key was found")
}
