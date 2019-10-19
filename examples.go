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

package main

import (
	"fmt"
	"net/http"

	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func Examples() {

	config := &pulsar.Config{
		WebServiceURL: "http://localhost:8080",
		HTTPClient:    http.DefaultClient,

		// If the server enable the TLSAuth
		// Auth: auth.NewAuthenticationTLS()

		// If the server enable the TokenAuth
		// TokenAuth: auth.NewAuthenticationToken()
	}

	// the default NewPulsarClient will use v2 APIs. If you need to request other version APIs,
	// you can specified the API version like this:
	// admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V2)
	admin, err := pulsar.New(config)
	if err != nil {
		// handle the err
		return
	}

	// more APIs, you can find them in the pkg/pulsar/admin.go
	// You can find all the method in the pkg/pulsar
	clusters, err := admin.Clusters().List()
	if err != nil {
		// handle the error
	}

	// handle the result
	fmt.Println(clusters)
}
