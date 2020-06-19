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
	"encoding/json"
	"fmt"
	"testing"

	"github.com/streamnative/pulsarctl/pkg/auth"
	"github.com/streamnative/pulsarctl/pkg/pulsar/common"
	util "github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

func Test_Pulsar_test(t *testing.T) {
	admin, err := New(&common.Config{
		WebServiceURL:    "http://34.83.175.243:8080",
		AuthPlugin: auth.TokenPluginName,
		AuthParams: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6Ik5UTTROMEUwTlRSQk5FWTJRemMwTkRrME9FUTVRakV5TnpBek1rSTBNak00TVVSRFJESTVPUSJ9.eyJodHRwczovL3N0cmVhbW5hdGl2ZS5pby91c2VybmFtZSI6InB1bHNhci1jbGllbnQtc2FAZTJlLXRlc3Quc3RyZWFtbmF0aXZlLXRlc3QuYXV0aDAuY29tIiwiaXNzIjoiaHR0cHM6Ly9zdHJlYW1uYXRpdmUtdGVzdC5hdXRoMC5jb20vIiwic3ViIjoiZHpRd2hoVDFlZUNBSDJUdWRkeDRyQzRUWFpvWGNhTUhAY2xpZW50cyIsImF1ZCI6InVybjpzbjpwdWxzYXI6ZTJlLXRlc3Q6cHVsc2FyaW5zdGFuY2UiLCJpYXQiOjE1ODk3NzEwNTAsImV4cCI6MTU5MDM3NTg1MCwiYXpwIjoiZHpRd2hoVDFlZUNBSDJUdWRkeDRyQzRUWFpvWGNhTUgiLCJzY29wZSI6ImFkbWluIiwiZ3R5IjoiY2xpZW50LWNyZWRlbnRpYWxzIiwicGVybWlzc2lvbnMiOlsiYWRtaW4iXX0.AdwHc4ThCu5ryob8UEOsgaxDKJey1qfb2BVGJUSyIyk7ll6nTDf5stDUn-UbqTOx1Zlkynss25y7G84MpnkhDhVARbEJUNLHpw499jQ531EJL1QC6fjNMO5iujWrO2yt80Hooexpg5xwASWY2hWkwcpv1Cf-mlNJVmq6TtEd8f9Kuyw8OMblhXLF1MUv8FLKwupV6Q2cS9dkMzW1rTP440ip7MLJjlsUV2qOTynuaB6b0_l0cKqohpKvHbYPsg9XuOTd8erZCrGwZ_IZ6bNl1X7GT0ZL4K36MbVeB9ub8828D-mMyyI_Kyt7t376Th51SbT4w_r2l6ve3mOxR3knUg",
		PulsarAPIVersion: common.V2,
	})
	if err != nil {
		t.Fatal(err)
	}

	topicname, err := util.GetTopicName("test-stats-test-partition-1")
	if err != nil {
		t.Fatal(err)
	}

	stas, err := admin.Topics().GetStats(*topicname)
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(&stas)
	fmt.Println(string(b))
}

