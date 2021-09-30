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

package topic

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

type setOffloadPoliciesFlags struct {
	offloadersDirectory        string
	driver                     string
	region                     string
	bucket                     string
	endpoint                   string
	awsID                      string
	awsSecret                  string
	s3Role                     string
	s3RoleSessionName          string
	maxThreads                 int64
	prefetchRounds             int64
	maxBlockSizeInBytes        int64
	readBufferSizeInBytes      int64
	offloadThresholdInBytes    int64
	offloadDeletionLagInMillis int64
	offloadedReadPriority      string
	fileSystemProfilePath      string
	fileSystemURI              string
}

const (
	DefaultMaxThreads            = 2
	DefaultOffloaderDirectory    = "./offloaders"
	DefaultPrefetchRounds        = 1
	DefaultMaxBlockSizeInBytes   = 64 * 1024 * 1024 // 64MB
	DefaultReadBufferSizeInBytes = 1 * 1024 * 1024  // 1MB
	DefaultS3RoleSessionName     = "pulsar-s3-offload"
	DefaultFileSystemURI         = "hdfs://127.0.0.1:9000"
)

func SetOffloadPoliciesCmd(vc *cmdutils.VerbCmd) {
	desc := cmdutils.LongDescription{}
	desc.CommandUsedFor = "Set the offload policies for a topic"
	desc.CommandPermission = "This command requires tenant admin permissions."

	var examples []cmdutils.Example
	examples = append(examples, cmdutils.Example{
		Desc: desc.CommandUsedFor,
		Command: "pulsarctl topics set-offload-policies persistent://myprop/clust/ns1/ds1 " +
			"--driver s3 " +
			"--region region " +
			"--bucket bucket " +
			"--endpoint endpoint " +
			"--max-block-size 8 " +
			"--read-buffer-size 9 " +
			"--threshold 10 " +
			"--offloaded-read-priority tiered-storage-first",
	})
	desc.CommandExamples = examples

	vc.SetDescription(
		"set-offload-policies",
		desc.CommandUsedFor,
		desc.ToString(),
		desc.ExampleToString(),
	)

	flags := setOffloadPoliciesFlags{}
	vc.FlagSetGroup.InFlagSet("Set offload policies", func(flagSet *pflag.FlagSet) {
		flagSet.StringVarP(
			&flags.offloadersDirectory,
			"offloaders-directory",
			"",
			DefaultOffloaderDirectory,
			"Offloaders directory is used to find the offload drivers")
		flagSet.StringVarP(
			&flags.driver,
			"driver",
			"",
			"",
			"ManagedLedger offload driver, "+
				"supports S3, aws-s3, google-cloud-storage, filesystem, azureblob and aliyun-oss")
		flagSet.StringVarP(
			&flags.region,
			"region",
			"",
			"",
			"ManagedLedger offload region, s3 and google-cloud-storage requires this parameter")
		flagSet.StringVarP(
			&flags.bucket,
			"bucket",
			"",
			"",
			"ManagedLedger offload bucket, s3 and google-cloud-storage requires this parameter")
		flagSet.StringVarP(
			&flags.endpoint,
			"endpoint",
			"",
			"",
			"ManagedLedger offload service endpoint, only s3 requires this parameter")
		flagSet.StringVarP(
			&flags.awsID,
			"aws-id",
			"",
			"",
			"AWS Credential Id to use when using driver S3 or aws-s3")
		flagSet.StringVarP(
			&flags.awsSecret,
			"aws-secret",
			"",
			"",
			"AWS Credential Secret to use when using driver S3 or aws-s3")
		flagSet.StringVarP(
			&flags.s3Role,
			"s3-role",
			"",
			"",
			"S3 Role used for STSAssumeRoleSessionCredentialsProvider")
		flagSet.StringVarP(
			&flags.s3RoleSessionName,
			"s3-role-session-name",
			"",
			"",
			"S3 role session name used for STSAssumeRoleSessionCredentialsProvider")
		flagSet.StringVarP(
			&flags.fileSystemProfilePath,
			"file-system-profile-path",
			"",
			"",
			"file system profile path to use when using driver filesystem")
		flagSet.StringVarP(
			&flags.fileSystemURI,
			"file-system-uri",
			"",
			DefaultFileSystemURI,
			"file system url to use when using driver filesystem")
		flagSet.Int64VarP(
			&flags.maxThreads,
			"max-threads",
			"",
			DefaultMaxThreads,
			"ManagedLedger offload max threads")
		flagSet.Int64VarP(
			&flags.maxBlockSizeInBytes,
			"max-block-size",
			"",
			DefaultMaxBlockSizeInBytes,
			"ManagedLedger offload max block Size in bytes, "+
				"s3 and google-cloud-storage requires this parameter")
		flagSet.Int64VarP(
			&flags.readBufferSizeInBytes,
			"read-buffer-size",
			"",
			DefaultReadBufferSizeInBytes,
			"ManagedLedger offload read buffer size in bytes, "+
				"s3 and google-cloud-storage requires this parameter")
		flagSet.Int64VarP(
			&flags.offloadThresholdInBytes,
			"threshold",
			"",
			0,
			"ManagedLedger offload threshold in bytes")
		flagSet.Int64VarP(
			&flags.offloadDeletionLagInMillis,
			"deletion-lag",
			"",
			0,
			"ManagedLedger offload deletion lag in millisecond")
		flagSet.Int64VarP(
			&flags.prefetchRounds,
			"prefetch-rounds",
			"",
			DefaultPrefetchRounds,
			"ManagedLedger offload prefetch rounds")
		flagSet.StringVarP(
			&flags.offloadedReadPriority,
			"offloaded-read-priority",
			"",
			utils.TieredStorageFirst.String(),
			"Read priority for offloaded messages. By default, "+
				"once messages are offloaded to long-term storage, "+
				"brokers read messages from long-term storage, "+
				"but messages can still exist in BookKeeper for a period depends on your configuration. "+
				"For messages that exist in both long-term storage and BookKeeper, "+
				"you can set where to read messages from "+
				"with the option `tiered-storage-first` or `bookkeeper-first`")

		_ = cobra.MarkFlagRequired(flagSet, "driver")
		_ = cobra.MarkFlagRequired(flagSet, "threshold")
	})
	vc.EnableOutputFlagSet()

	vc.SetRunFuncWithNameArg(func() error {
		return doSetOffloadPolicies(vc, flags)
	}, "the topic name is not specified or the topic name is specified more than one")

}

func doSetOffloadPolicies(vc *cmdutils.VerbCmd, flags setOffloadPoliciesFlags) error {
	topic := vc.NameArg
	admin := cmdutils.NewPulsarClient()
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return err
	}

	data := utils.OffloadPolicies{
		ManagedLedgerOffloadMaxThreads:            flags.maxThreads,
		OffloadersDirectory:                       flags.offloadersDirectory,
		ManagedLedgerOffloadPrefetchRounds:        flags.prefetchRounds,
		ManagedLedgerOffloadDriver:                flags.driver,
		ManagedLedgerOffloadThresholdInBytes:      flags.offloadThresholdInBytes,
		ManagedLedgerOffloadDeletionLagInMillis:   flags.offloadDeletionLagInMillis,
		ManagedLedgerOffloadBucket:                flags.bucket,
		ManagedLedgerOffloadRegion:                flags.region,
		ManagedLedgerOffloadServiceEndpoint:       flags.endpoint,
		ManagedLedgerOffloadMaxBlockSizeInBytes:   flags.maxBlockSizeInBytes,
		ManagedLedgerOffloadReadBufferSizeInBytes: flags.readBufferSizeInBytes,
	}

	offloadedReadPriority, err := utils.ParseOffloadedReadPriority(flags.offloadedReadPriority)
	if err != nil {
		return err
	}
	data.ManagedLedgerOffloadedReadPriority = offloadedReadPriority

	switch {
	case strings.EqualFold(flags.driver, "s3") || strings.EqualFold(flags.driver, "aws-s3"):
		if flags.s3RoleSessionName == "" {
			data.S3ManagedLedgerOffloadRoleSessionName = DefaultS3RoleSessionName
		} else {
			data.S3ManagedLedgerOffloadRoleSessionName = flags.s3RoleSessionName
		}
		if flags.s3Role != "" {
			data.S3ManagedLedgerOffloadRole = flags.s3Role
		}
		if flags.s3RoleSessionName != "" {
			data.S3ManagedLedgerOffloadRoleSessionName = flags.s3RoleSessionName
		}
		if flags.awsID != "" {
			data.S3ManagedLedgerOffloadCredentialID = flags.awsID
		}
		if flags.awsSecret != "" {
			data.S3ManagedLedgerOffloadCredentialSecret = flags.awsSecret
		}
		data.S3ManagedLedgerOffloadRegion = flags.region
		data.S3ManagedLedgerOffloadBucket = flags.bucket
		data.S3ManagedLedgerOffloadServiceEndpoint = flags.endpoint
		data.S3ManagedLedgerOffloadServiceEndpoint = flags.endpoint
	case strings.EqualFold(flags.driver, "google-cloud-storage"):
		data.GcsManagedLedgerOffloadRegion = flags.region
		data.GcsManagedLedgerOffloadBucket = flags.bucket
		data.GcsManagedLedgerOffloadMaxBlockSizeInBytes = flags.maxBlockSizeInBytes
		data.GcsManagedLedgerOffloadReadBufferSizeInBytes = flags.readBufferSizeInBytes
	case strings.EqualFold(flags.driver, "filesystem"):
		data.FileSystemURI = flags.fileSystemURI
		data.FileSystemProfilePath = flags.fileSystemProfilePath
	}

	err = admin.Topics().SetOffloadPolicies(*topicName, data)
	if err == nil {
		vc.Command.Printf("Set the offload policies successfully for [%s]", topicName.String())
	}

	return err
}
