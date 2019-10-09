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
	"net/url"
	"strconv"
)

type Subscriptions interface {
	Create(TopicName, string, MessageID) error
	Delete(TopicName, string) error
	List(TopicName) ([]string, error)
	ResetCursorToMessageID(TopicName, string, MessageID) error
	ResetCursorToTimestamp(TopicName, string, int64) error
	ClearBacklog(TopicName, string) error
	SkipMessages(TopicName, string, int64) error
	ExpireMessages(TopicName, string, int64) error
	ExpireAllMessages(TopicName, int64) error
}

type subscriptions struct {
	client   *client
	basePath string
	SubPath  string
}

func (c *client) Subscriptions() Subscriptions {
	return &subscriptions{
		client:   c,
		basePath: "",
		SubPath:  "subscription",
	}
}

func (s *subscriptions) Create(topic TopicName, sName string, messageID MessageID) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName))
	return s.client.put(endpoint, messageID)
}

func (s *subscriptions) Delete(topic TopicName, sName string) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName))
	return s.client.delete(endpoint)
}

func (s *subscriptions) List(topic TopicName) ([]string, error) {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), "subscriptions")
	var list []string
	return list, s.client.get(endpoint, &list)
}

func (s *subscriptions) ResetCursorToMessageID(topic TopicName, sName string, id MessageID) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName), "resetcursor")
	return s.client.post(endpoint, id)
}

func (s *subscriptions) ResetCursorToTimestamp(topic TopicName, sName string, timestamp int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"resetcursor", strconv.FormatInt(timestamp, 10))
	return s.client.post(endpoint, "")
}

func (s *subscriptions) ClearBacklog(topic TopicName, sName string) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName), "skip_all")
	return s.client.post(endpoint, "")
}

func (s *subscriptions) SkipMessages(topic TopicName, sName string, n int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"skip", strconv.FormatInt(n, 10))
	return s.client.post(endpoint, "")
}

func (s *subscriptions) ExpireMessages(topic TopicName, sName string, expire int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"expireMessages", strconv.FormatInt(expire, 10))
	return s.client.post(endpoint, "")
}

func (s *subscriptions) ExpireAllMessages(topic TopicName, expire int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), "all_subscription",
		"expireMessages", strconv.FormatInt(expire, 10))
	return s.client.post(endpoint, "")
}
