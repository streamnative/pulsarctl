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

package algorithm

import (
	"github.com/pkg/errors"
)

// SignatureAlgorithm is a collection of all signature algorithm and it provides
// some basic method to use.
type SignatureAlgorithm interface {
	// GetName is used to getting the signature algorithm name. We provide HS256, HS384, HS512.
	GetName() string

	// IsHMAC is used to checking if the algorithm is belong to HMAC.
	IsHMAC() bool

	// GenerateSecret is used to generating a secret.
	GenerateSecret() []byte
}

func GetSignatureAlgorithm(algorithm string) (SignatureAlgorithm, error) {
	switch algorithm {
	case "HS256":
		return new(HS256), nil
	case "HS384":
		return new(HS384), nil
	case "HS512":
		return new(HS512), nil
	default:
		return nil, errors.Errorf("the signature algorithm '%s' is invalid. Valid options include: 'HS256', "+
			"'HS384', 'HS512'\n", algorithm)
	}
}
