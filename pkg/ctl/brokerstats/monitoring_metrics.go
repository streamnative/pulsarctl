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

package brokerstats

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"
)

func dumpMonitoringMetrics(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "Dumps the metrics for Monitoring"
	desc.CommandPermission = "This command requires super-user permissions."

	var examples []pulsar.Example
	get := pulsar.Example{
		Desc:    "Dumps the metrics for Monitoring",
		Command: "pulsarctl broker-stats monitoring-metrics",
	}
	examples = append(examples, get)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out: "{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_ml_cache_evictions\": 0,\n" +
			"        \"brk_ml_cache_hits_rate\": 0.0,\n" +
			"        \"brk_ml_cache_hits_throughput\": 0.0,\n" +
			"        \"brk_ml_cache_misses_rate\": 0.0,\n" +
			"        \"brk_ml_cache_misses_throughput\": 0.0,\n" +
			"        \"brk_ml_cache_pool_active_allocations\": 0,\n" +
			"        \"brk_ml_cache_pool_active_allocations_huge\": 0,\n" +
			"        \"brk_ml_cache_pool_active_allocations_normal\": 0,\n" +
			"        \"brk_ml_cache_pool_active_allocations_small\": 0,\n" +
			"        \"brk_ml_cache_pool_active_allocations_tiny\": 0,\n" +
			"        \"brk_ml_cache_pool_allocated\": 0,\n" +
			"        \"brk_ml_cache_pool_used\": 0,\n" +
			"        \"brk_ml_cache_used_size\": 0,\n" +
			"        \"brk_ml_count\": 3\n" +
			"    },\n" +
			"    \"dimensions\": {}\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_AddEntryLatencyBuckets_0.0_0.5\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_0.5_1.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_1.0_5.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_10.0_20.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_100.0_200.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_20.0_50.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_200.0_1000.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_5.0_10.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_50.0_100.0\": 0,\n" +
			"        \"brk_AddEntryLatencyBuckets_OVERFLOW\": 0,\n" +
			"        \"brk_in_rate\": 0.0,\n" +
			"        \"brk_in_tp_rate\": 0.0,\n" +
			"        \"brk_max_replication_delay_second\": 0.0,\n" +
			"        \"brk_msg_backlog\": 0.0,\n" +
			"        \"brk_no_of_consumers\": 3,\n" +
			"        \"brk_no_of_producers\": 2,\n" +
			"        \"brk_no_of_replicators\": 0,\n" +
			"        \"brk_no_of_subscriptions\": 3,\n" +
			"        \"brk_out_rate\": 0.0,\n" +
			"        \"brk_out_tp_rate\": 0.0,\n" +
			"        \"brk_replication_backlog\": 0.0,\n" +
			"        \"brk_storage_size\": 0.0\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"namespace\": \"public/functions\"\n" +
			"    }\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_zk_read_rate_s\": 0.0,\n" +
			"        \"brk_zk_read_time_75percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_95percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_99_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_99_9_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_mean_ms\": \"NaN\",\n" +
			"        \"brk_zk_read_time_median_ms\": \"NaN\"\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"broker\": \"127.0.0.1\",\n" +
			"        \"cluster\": \"standalone\",\n" +
			"        \"metric\": \"zk_read_latency\"\n" +
			"    }\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_default_pool_allocated\": 402653184,\n" +
			"        \"brk_default_pool_used\": 4521984,\n" +
			"        \"jvm_direct_memory_used\": 2550137118,\n" +
			"        \"jvm_gc_old_count\": 0,\n" +
			"        \"jvm_gc_old_pause\": 0,\n" +
			"        \"jvm_gc_young_count\": 0,\n" +
			"        \"jvm_gc_young_pause\": 0,\n" +
			"        \"jvm_heap_used\": 350432256,\n" +
			"        \"jvm_max_direct_memory\": 4294967296,\n" +
			"        \"jvm_max_memory\": 2147483648,\n" +
			"        \"jvm_thread_cnt\": 434,\n" +
			"        \"jvm_total_memory\": 2147483648\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"metric\": \"jvm_metrics\"\n" +
			"    }\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_topic_load_rate_s\": 0.0,\n" +
			"        \"brk_topic_load_time_75percentile_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_95percentile_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_99_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_99_9_percentile_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_mean_ms\": \"NaN\",\n" +
			"        \"brk_topic_load_time_median_ms\": \"NaN\"\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"broker\": \"127.0.0.1\",\n" +
			"        \"cluster\": \"standalone\",\n" +
			"        \"metric\": \"topic_load_times\"\n" +
			"    }\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_ml_AddEntryBytesRate\": 0.0,\n" +
			"        \"brk_ml_AddEntryErrors\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_0.0_0.5\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_0.5_1.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_1.0_5.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_10.0_20.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_100.0_200.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_20.0_50.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_200.0_1000.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_5.0_10.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_50.0_100.0\": 0.0,\n" +
			"        \"brk_ml_AddEntryLatencyBuckets_OVERFLOW\": 0.0,\n" +
			"        \"brk_ml_AddEntryMessagesRate\": 0.0,\n" +
			"        \"brk_ml_AddEntrySucceed\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_0.0_128.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_1024.0_2084.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_102400.0_1232896.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_128.0_512.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_16384.0_102400.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_2084.0_4096.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_4096.0_16384.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_512.0_1024.0\": 0.0,\n" +
			"        \"brk_ml_EntrySizeBuckets_OVERFLOW\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_0.0_0.5\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_0.5_1.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_1.0_5.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_10.0_20.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_100.0_200.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_20.0_50.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_200.0_1000.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_5.0_10.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_50.0_100.0\": 0.0,\n" +
			"        \"brk_ml_LedgerSwitchLatencyBuckets_OVERFLOW\": 0.0,\n" +
			"        \"brk_ml_MarkDeleteRate\": 0.0,\n" +
			"        \"brk_ml_NumberOfMessagesInBacklog\": 0.0,\n" +
			"        \"brk_ml_ReadEntriesBytesRate\": 0.0,\n" +
			"        \"brk_ml_ReadEntriesErrors\": 0.0,\n" +
			"        \"brk_ml_ReadEntriesRate\": 0.0,\n" +
			"        \"brk_ml_ReadEntriesSucceeded\": 0.0,\n" +
			"        \"brk_ml_StoredMessagesSize\": 0.0\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"namespace\": \"public/functions/persistent\"\n" +
			"    }\n" +
			"}\n" +
			"\n" +
			"{\n" +
			"    \"metrics\": {\n" +
			"        \"brk_zk_write_rate_s\": 0.0,\n" +
			"        \"brk_zk_write_time_75percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_95percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_99_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_99_9_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_99_percentile_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_mean_ms\": \"NaN\",\n" +
			"        \"brk_zk_write_time_median_ms\": \"NaN\"\n" +
			"    },\n" +
			"    \"dimensions\": {\n" +
			"        \"broker\": \"127.0.0.1\",\n" +
			"        \"cluster\": \"standalone\",\n" +
			"        \"metric\": \"zk_write_latency\"\n" +
			"    }\n" +
			"}",
	}
	out = append(out, successOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"monitoring-metrics",
		"Dumps the metrics for Monitoring",
		desc.ToString(),
		desc.ExampleToString(),
		"monitoring-metrics")

	vc.SetRunFunc(func() error {
		return doDumpMonitoringMetrics(vc)
	})
}

func doDumpMonitoringMetrics(vc *cmdutils.VerbCmd) error {
	admin := cmdutils.NewPulsarClient()
	metrics, err := admin.BrokerStats().GetMetrics()
	if err != nil {
		cmdutils.PrintError(vc.Command.OutOrStderr(), err)
	} else {
		cmdutils.PrintJSON(vc.Command.OutOrStdout(), metrics)
	}
	return err
}
