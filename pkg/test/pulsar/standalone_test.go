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

package pulsar

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStandalone(t *testing.T) {
	ctx := context.Background()
	standalone := DefaultStandalone()
	err := standalone.Start(ctx)
	// nolint
	defer standalone.Stop(ctx)
	if err != nil {
		t.Fatal(err)
	}

	port, err := standalone.MappedPort(ctx, "8080")
	if err != nil {
		t.Fatal(err)
	}
	path := "http://localhost:" + port.Port() + "/admin/v2/tenants"

	resp, err := http.Get(path)
	// nolint
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)

}
