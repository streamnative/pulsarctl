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

package bookkeeper

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultCluster(t *testing.T) {
	ctx := context.Background()
	bkCluster, err := DefaultCluster()
	// nolint
	defer bkCluster.Close(ctx)
	if err != nil {
		t.Fatal(err)
	}
	err = bkCluster.Start(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer bkCluster.Stop(ctx)

	path, err := bkCluster.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(path + "/api/v1/bookie/list_bookie_info")
	// nolint
	defer resp.Body.Close()
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
}
