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

package autorecovery

import (
	"context"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/test/bookkeeper"

	"github.com/stretchr/testify/assert"
)

func TestWhoIsAuditor(t *testing.T) {
	// todo: fix the container issue
	return
	// prepare the bookkeeper cluster environment
	ctx := context.Background()
	bk, err := bookkeeper.NewBookieCluster(&bookkeeper.ClusterSpec{
		ClusterName: "test-who-is-auditor",
		NumBookies:  1,
		BookieEnv: map[string]string{
			"BK_autoRecoveryDaemonEnabled": "true",
		},
	})
	// nolint
	defer bk.Close(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = bk.Start(ctx)
	// nolint
	defer bk.Stop(ctx)
	if err != nil {
		t.Fatal(err)
	}

	allBK := bk.GetAllBookieContainerID()

	httpAddr, err := bk.GetHTTPServiceURL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	args := []string{"--bookie-service-url", httpAddr, "who-is-auditor"}
	out, execErr, nameErr, err := testAutoRecoveryCommands(whoIsAuditorCmd, args)
	if err != nil {
		t.Fatal(err)
	}
	if nameErr != nil {
		t.Fatal(nameErr)
	}
	if execErr != nil {
		t.Fatal(execErr)
	}

	assert.Contains(t, out.String(), allBK[0][:12])
}
