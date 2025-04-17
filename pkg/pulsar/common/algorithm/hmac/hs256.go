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

package hmac

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"

	"github.com/streamnative/pulsarctl/pkg/pulsar/common/algorithm/keypair"

	"github.com/pkg/errors"
)

type HS256 struct{}

func (h *HS256) GenerateSecret() ([]byte, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	s := hmac.New(sha256.New, bytes)
	return s.Sum(nil), nil
}

func (h *HS256) GenerateKeyPair() (*keypair.KeyPair, error) {
	return nil, errors.New("unsupported operation")
}
