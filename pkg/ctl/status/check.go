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

package status

import (
	"net/http"

	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/auth"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/admin/config"
	"github.com/apache/pulsar-client-go/pulsaradmin/pkg/rest"
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func checkStatusCmd(vc *cmdutils.VerbCmd) {
	var desc cmdutils.LongDescription
	desc.CommandUsedFor = "Check service(broker or proxy) status"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	check := cmdutils.Example{
		Desc:    "Check service(broker or proxy) status",
		Command: "pulsarctl status check",
	}
	examples = append(examples, check)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out:  "Ok",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"check",
		"Check service(broker or proxy) status",
		desc.ToString(),
		desc.ExampleToString(),
		"check")

	vc.SetRunFunc(func() error {
		return doCheckStatus(vc)
	})
	vc.EnableOutputFlagSet()
}
func doCheckStatus(vc *cmdutils.VerbCmd) error {
	cfg := cmdutils.PulsarCtlConfig
	if len(cfg.WebServiceURL) == 0 {
		cfg.WebServiceURL = admin.DefaultWebServiceURL
	}
	authProvider, err := auth.GetAuthProvider((*config.Config)(cfg))
	if err != nil {
		return err
	}
	client := &rest.Client{
		ServiceURL:  cmdutils.PulsarCtlConfig.WebServiceURL,
		VersionInfo: admin.ReleaseVersion,
		HTTPClient: &http.Client{
			Timeout:   admin.DefaultHTTPTimeOutDuration,
			Transport: authProvider,
		},
	}
	data, err := client.GetWithQueryParams("/status.html", nil, nil, false)
	if err != nil {
		return err
	}
	vc.Command.Printf(string(data))
	return nil
}
