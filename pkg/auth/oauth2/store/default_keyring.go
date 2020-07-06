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

package store

import (
	"path/filepath"

	"github.com/99designs/keyring"
	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

const (
	serviceName  = "pulsar"
	keyChainName = "pulsarctl"
)

func MakeKeyringStore() (Store, error) {
	kr, err := makeKeyring()
	if err != nil {
		return nil, err
	}
	return NewKeyringStore(kr)
}

func makeKeyring() (keyring.Keyring, error) {
	return keyring.Open(keyring.Config{
		AllowedBackends:          keyring.AvailableBackends(),
		ServiceName:              serviceName,
		KeychainName:             keyChainName,
		KeychainTrustApplication: true,
		FileDir:                  filepath.Join(util.HomeDir(), "~/.config/pulsar", "credentials"),
		FilePasswordFunc:         keyringPrompt,
	})
}

func keyringPrompt(prompt string) (string, error) {
	return "", nil
}
