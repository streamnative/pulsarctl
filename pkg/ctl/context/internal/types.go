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

package internal

type Config struct {
	AuthInfos      map[string]*AuthInfo `yaml:"auth-info"`
	Contexts       map[string]*Context  `yaml:"contexts"`
	CurrentContext string               `yaml:"current-context"`
}

type AuthInfo struct {
	// LocationOfOrigin indicates where this object came from.  It is used for round tripping
	// config post-merge, but never serialized.
	LocationOfOrigin      string
	ClientCertificate     string `yaml:"client-certificate,omitempty"`
	ClientCertificateData []byte `yaml:"client-certificate-data,omitempty"`
	ClientKey             string `yaml:"client-key,omitempty"`
	ClientKeyData         []byte `yaml:"client-key-data,omitempty"`
	Token                 string `yaml:"token,omitempty"`
	TokenFile             string `yaml:"tokenFile,omitempty"`
}

type Context struct {
	AuthInfo         string `yaml:"user"`
	BrokerServiceURL string `yaml:"admin-service-url"`
	BookieServiceURL string `yaml:"bookie-service-url"`
}

type ConfigOverrides struct {
	AuthInfo       AuthInfo
	Context        Context
	CurrentContext string
	Timeout        string
}

func NewConfig() *Config {
	return &Config{
		AuthInfos: make(map[string]*AuthInfo),
		Contexts:  make(map[string]*Context),
	}
}
