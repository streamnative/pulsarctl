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
	"io/ioutil"
	"os"
	"strings"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"

	"github.com/spf13/pflag"
)

func show(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "This command is used for showing the content of a token."
	desc.CommandPermission = "This command does not need any permission."

	var examples []cmdutils.Example
	readTokenFromEnv := cmdutils.Example{
		Desc:    "Read a token from the env TOKEN.",
		Command: "pulsarctl token show",
	}

	readTokenFromString := cmdutils.Example{
		Desc:    "Read a token from a given string.",
		Command: "pulsarctl token show --token-string (token)",
	}

	readTokenFromFile := cmdutils.Example{
		Desc:    "Read a token from a given file.",
		Command: "pulsarctl token show --token-file (token)",
	}
	examples = append(examples, readTokenFromEnv, readTokenFromString, readTokenFromFile)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	defaultOutput := cmdutils.Output{
		Desc: "Show the content of the given token.",
		Out:  "The algorithm and subject of the token are (signature algorithm), (subject).",
	}

	tokenNotSpecifiedErr := cmdutils.Output{
		Desc: "There is no token to show.",
		Out:  "[✖]  " + errNoTokenStringOrFile.Error(),
	}

	tokenSpecifiedMoreThanOneErr := cmdutils.Output{
		Desc: "Too many tokens to show.",
		Out:  "[✖]  " + errTokenSpecifiedMoreThanOne.Error(),
	}

	out = append(out, defaultOutput, tokenNotSpecifiedErr, tokenSpecifiedMoreThanOneErr)
	desc.CommandOutput = out

	vc.SetDescription(
		"show",
		"Show the algorithm and subject of a token.",
		desc.ToString(),
		desc.ExampleToString())

	var tokenString, tokenFile string

	vc.SetRunFunc(func() error {
		return doShow(vc, tokenString, tokenFile)
	})

	vc.FlagSetGroup.InFlagSet("Show token", func(set *pflag.FlagSet) {
		set.StringVar(&tokenString, "token-string", "",
			"The token string you would like to show the content.")
		set.StringVar(&tokenFile, "token-file", "",
			"The token file you would like to show the content.")
	})
	vc.EnableOutputFlagSet()
}

func doShow(vc *cmdutils.VerbCmd, tokenString, tokenFile string) error {
	token, err := readToken(tokenString, tokenFile)
	if err != nil {
		return err
	}

	tokenUtil := cmdutils.NewPulsarClient().Token()
	algorithm, err := tokenUtil.GetAlgorithm(token)
	if err != nil {
		return err
	}
	subject, err := tokenUtil.GetSubject(token)
	if err != nil {
		return err
	}

	vc.Command.Printf("The algorithm and subject of the token are %s, %s.\n", algorithm, subject)

	return nil
}

func readToken(tokenString, tokenFile string) (string, error) {
	tokenString = strings.TrimSpace(tokenString)
	tokenFile = strings.TrimSpace(tokenFile)
	switch {
	case tokenString == "" && tokenFile == "":
		token := os.Getenv("TOKEN")
		if strings.TrimSpace(token) != "" {
			return token, nil
		}
		return "", errNoTokenStringOrFile
	case tokenString != "" && tokenFile != "":
		return "", errTokenSpecifiedMoreThanOne
	case tokenString != "":
		return tokenString, nil
	case tokenFile != "":
		data, err := ioutil.ReadFile(tokenFile)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return "", nil
}
