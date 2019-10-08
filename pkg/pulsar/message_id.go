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
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type MessageID struct {
	LedgerID         int64 `json:"ledgerId"`
	EntryID          int64 `json:"entryId"`
	PartitionedIndex int   `json:"partitionedIndex"`
}

var Latest = MessageID{0x7fffffffffffffff, 0x7fffffffffffffff, -1}
var Earliest = MessageID{-1, -1, -1}

func ParseMessageId(str string) (*MessageID, error) {
	s := strings.Split(str, ":")
	if len(s) != 2 {
		return nil, errors.Errorf("Invalid message id string. %s", str)
	}

	ledgerId, err := strconv.ParseInt(s[0], 10, 64)
	if err != nil {
		return nil, errors.Errorf("Invalid ledger id string. %s", str)
	}

	entryId, err := strconv.ParseInt(s[1], 10, 64)
	if err != nil {
		return nil, errors.Errorf("Invalid entry id string. %s", str)
	}

	return &MessageID{LedgerID: ledgerId, EntryID: entryId, PartitionIndex: -1}, nil
}
