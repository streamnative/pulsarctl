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

package oauth2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const (
	KeyFileTypeServiceAccount = "sn_service_account"
)

type KeyFileProvider struct {
	KeyFile string
}

type KeyFile struct {
	Type         string `json:"type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	ClientEmail  string `json:"client_email"`
}

func NewClientCredentialsProviderFromKeyFile(keyFile string) *KeyFileProvider {
	return &KeyFileProvider{
		KeyFile: keyFile,
	}
}

var _ ClientCredentialsProvider = &KeyFileProvider{}

func (k *KeyFileProvider) GetClientCredentials() (*KeyFile, error) {
	keyFile, err := ioutil.ReadFile(k.KeyFile)
	if err != nil {
		return nil, err
	}

	var v KeyFile
	err = json.Unmarshal(keyFile, &v)
	if err != nil {
		return nil, err
	}
	if v.Type != KeyFileTypeServiceAccount {
		return nil, fmt.Errorf("open %s: unsupported format", k.KeyFile)
	}

	return &v, nil
}
