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

package functionsworker

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
)

func monitoringMetrics(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Dump metrics for Monitoring"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []cmdutils.Example
	stats := cmdutils.Example{
		Desc:    "Dump metrics for Monitoring",
		Command: "pulsarctl functions-worker monitoring-metrics",
	}
	examples = append(examples, stats)
	desc.CommandExamples = examples

	var out []cmdutils.Output
	successOut := cmdutils.Output{
		Desc: "normal output",
		Out: "[\n" +
			"  {\n" +
			"    \"metrics\": {\n" +
			"      \"fun_default_pool_allocated\": 402653184,\n" +
			"      \"fun_default_pool_used\": 4734976,\n" +
			"      \"jvm_direct_memory_used\": 2550137118,\n" +
			"      \"jvm_gc_old_count\": 0,\n" +
			"      \"jvm_gc_old_pause\": 0,\n" +
			"      \"jvm_gc_young_count\": 0,\n" +
			"      \"jvm_gc_young_pause\": 0,\n" +
			"      \"jvm_heap_used\": 305348512,\n" +
			"      \"jvm_max_direct_memory\": 4294967296,\n" +
			"      \"jvm_max_memory\": 2147483648,\n" +
			"      \"jvm_thread_cnt\": 446,\n" +
			"      \"jvm_total_memory\": 2147483648\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"      \"metric\": \"jvm_metrics\"\n" +
			"    }\n" +
			"  }\n" +
			"]",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"monitoring-metrics",
		"Dump metrics for Monitoring",
		desc.ToString(),
		desc.ExampleToString(),
		"monitoring-metrics",
	)

	// set the run function
	vc.SetRunFunc(func() error {
		return doMonitoringMetrics(vc)
	})
}

func doMonitoringMetrics(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	metrics, err := admin.FunctionsWorker().GetMetrics()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), metrics)
	}

	return err
}
