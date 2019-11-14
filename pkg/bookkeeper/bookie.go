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
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/bookkeeper/bkdata"
)

type Bookie interface {
	// List all the available bookies
	List(bkdata.BookieType, bool) (map[string]string, error)

	// Get the bookies disk usage info of a cluster
	Info() (map[string]string, error)

	// Get the last log marker
	LastLogMark() (map[string]string, error)

	// Get all the files on disk of the current bookie
	ListDiskFile(bkdata.FileType) (map[string]string, error)

	// Expand storage for a bookie
	ExpandStorage() error

	// Trigger GC for a bookie
	GC() error

	// Check the GC status
	GCStatus() (map[string]string, error)

	// Details of the Garbage Collection
	GCDetails() ([]bkdata.GCStatus, error)

	// State of a bookie
	State() (*bkdata.State, error)
}

type bookie struct {
	bk       *bookieClient
	basePath string
	params   map[string]string
}

func (c *bookieClient) Bookie() Bookie {
	return &bookie{
		bk:       c,
		basePath: "/bookie",
		params:   make(map[string]string),
	}
}

func (b *bookie) List(t bkdata.BookieType, show bool) (map[string]string, error) {
	endpoint := b.bk.endpoint(b.basePath, "/list_bookies")
	b.params["type"] = t.String()
	b.params["print_hostnames"] = strconv.FormatBool(show)
	bookies := make(map[string]string)
	_, err := b.bk.Client.GetWithQueryParams(endpoint, &bookies, b.params, true)
	return bookies, err
}

func (b *bookie) Info() (map[string]string, error) {
	endpoint := b.bk.endpoint(b.basePath, "/list_bookie_info")
	info := make(map[string]string)
	err := b.bk.Client.Get(endpoint, &info)
	return info, err
}

func (b *bookie) LastLogMark() (map[string]string, error) {
	endpoint := b.bk.endpoint(b.basePath, "/last_log_mark")
	marker := make(map[string]string)
	err := b.bk.Client.Get(endpoint, &marker)
	return marker, err
}

func (b *bookie) ListDiskFile(fileType bkdata.FileType) (map[string]string, error) {
	endpoint := b.bk.endpoint(b.basePath, "/list_disk_file")
	b.params["file_type"] = fileType.String()
	files := make(map[string]string)
	_, err := b.bk.Client.GetWithQueryParams(endpoint, &files, b.params, true)
	return files, err
}

func (b *bookie) ExpandStorage() error {
	endpoint := b.bk.endpoint(b.basePath, "/expand_storage")
	return b.bk.Client.Put(endpoint, nil)
}

func (b *bookie) GC() error {
	endpoint := b.bk.endpoint(b.basePath, "/gc")
	return b.bk.Client.Put(endpoint, nil)
}

func (b *bookie) GCStatus() (map[string]string, error) {
	endpoint := b.bk.endpoint(b.basePath, "/gc")
	status := make(map[string]string)
	err := b.bk.Client.Get(endpoint, &status)
	return status, err
}

func (b *bookie) GCDetails() ([]bkdata.GCStatus, error) {
	endpoint := b.bk.endpoint(b.basePath, "/gc_details")
	details := make([]bkdata.GCStatus, 0)
	err := b.bk.Client.Get(endpoint, &details)
	return details, err
}

func (b *bookie) State() (*bkdata.State, error) {
	endpoint := b.bk.endpoint(b.basePath, "/state")
	var state bkdata.State
	err := b.bk.Client.Get(endpoint, &state)
	return &state, err
}
