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

package utils

import "github.com/pkg/errors"

type OffloadedReadPriority string

const (
	BookkeeperFirst    OffloadedReadPriority = "bookkeeper-first"
	TieredStorageFirst OffloadedReadPriority = "tiered-storage-first"
)

func (o OffloadedReadPriority) String() string {
	return string(o)
}

func ParseOffloadedReadPriority(s string) (OffloadedReadPriority, error) {
	switch s {
	case BookkeeperFirst.String():
		return BookkeeperFirst, nil
	case TieredStorageFirst.String():
		return TieredStorageFirst, nil
	default:
		return "", errors.New("unknown OffloadedReadPriority type")
	}
}

type OffloadPolicies struct {
	// common config
	OffloadersDirectory                     string                `json:"offloadersDirectory"`
	ManagedLedgerOffloadDriver              string                `json:"managedLedgerOffloadDriver"`
	ManagedLedgerOffloadMaxThreads          int64                 `json:"managedLedgerOffloadMaxThreads"`
	ManagedLedgerOffloadPrefetchRounds      int64                 `json:"managedLedgerOffloadPrefetchRounds"`
	ManagedLedgerOffloadThresholdInBytes    int64                 `json:"managedLedgerOffloadThresholdInBytes"`
	ManagedLedgerOffloadDeletionLagInMillis int64                 `json:"managedLedgerOffloadDeletionLagInMillis"`
	ManagedLedgerOffloadedReadPriority      OffloadedReadPriority `json:"managedLedgerOffloadedReadPriority"`

	// they are universal configurations and could be used to `aws-s3`, `google-cloud-storage` or `azureblob`.
	ManagedLedgerOffloadBucket                string `json:"managedLedgerOffloadBucket"`
	ManagedLedgerOffloadRegion                string `json:"managedLedgerOffloadRegion"`
	ManagedLedgerOffloadServiceEndpoint       string `json:"managedLedgerOffloadServiceEndpoint"`
	ManagedLedgerOffloadMaxBlockSizeInBytes   int64  `json:"managedLedgerOffloadMaxBlockSizeInBytes"`
	ManagedLedgerOffloadReadBufferSizeInBytes int64  `json:"managedLedgerOffloadReadBufferSizeInBytes"`

	// s3 config, set by service configuration
	S3ManagedLedgerOffloadRegion                string `json:"s3ManagedLedgerOffloadRegion"`
	S3ManagedLedgerOffloadBucket                string `json:"s3ManagedLedgerOffloadBucket"`
	S3ManagedLedgerOffloadServiceEndpoint       string `json:"s3ManagedLedgerOffloadServiceEndpoint"`
	S3ManagedLedgerOffloadMaxBlockSizeInBytes   int64  `json:"s3ManagedLedgerOffloadMaxBlockSizeInBytes"`
	S3ManagedLedgerOffloadReadBufferSizeInBytes int64  `json:"s3ManagedLedgerOffloadReadBufferSizeInBytes"`
	S3ManagedLedgerOffloadCredentialID          string `json:"s3ManagedLedgerOffloadCredentialId"`
	S3ManagedLedgerOffloadCredentialSecret      string `json:"s3ManagedLedgerOffloadCredentialSecret"`
	S3ManagedLedgerOffloadRole                  string `json:"s3ManagedLedgerOffloadRole"`
	S3ManagedLedgerOffloadRoleSessionName       string `json:"s3ManagedLedgerOffloadRoleSessionName"`

	// gcs config
	GcsManagedLedgerOffloadRegion                string `json:"gcsManagedLedgerOffloadRegion"`
	GcsManagedLedgerOffloadBucket                string `json:"gcsManagedLedgerOffloadBucket"`
	GcsManagedLedgerOffloadMaxBlockSizeInBytes   int64  `json:"gcsManagedLedgerOffloadMaxBlockSizeInBytes"`
	GcsManagedLedgerOffloadReadBufferSizeInBytes int64  `json:"gcsManagedLedgerOffloadReadBufferSizeInBytes"`
	GcsManagedLedgerOffloadServiceAccountKeyFile string `json:"gcsManagedLedgerOffloadServiceAccountKeyFile"`

	// file system config, set by service configuration
	FileSystemProfilePath string `json:"fileSystemProfilePath"`
	FileSystemURI         string `json:"fileSystemURI"`
}
