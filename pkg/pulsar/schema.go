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
	"fmt"
	"strconv"

	"github.com/streamnative/pulsarctl/pkg/pulsar/utils"
)

// Schema is admin interface for schema management
type Schema interface {
	// GetSchemaInfo retrieves the latest schema of a topic
	GetSchemaInfo(topic string) (*utils.SchemaInfo, error)

	// GetSchemaInfoWithVersion retrieves the latest schema with version of a topic
	GetSchemaInfoWithVersion(topic string) (*utils.SchemaInfoWithVersion, error)

	// GetSchemaInfoByVersion retrieves the schema of a topic at a given <tt>version</tt>
	GetSchemaInfoByVersion(topic string, version int64) (*utils.SchemaInfo, error)

	// DeleteSchema deletes the schema associated with a given <tt>topic</tt>
	DeleteSchema(topic string) error

	// CreateSchemaByPayload creates a schema for a given <tt>topic</tt>
	CreateSchemaByPayload(topic string, schemaPayload utils.PostSchemaPayload) error
}

type schemas struct {
	client   *client
	basePath string
}

// Schemas is used to access the schemas endpoints
func (c *client) Schemas() Schema {
	return &schemas{
		client:   c,
		basePath: "/schemas",
	}
}

func (s *schemas) GetSchemaInfo(topic string) (*utils.SchemaInfo, error) {
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return nil, err
	}
	var response utils.GetSchemaResponse
	endpoint := s.client.endpoint(s.basePath, topicName.GetTenant(), topicName.GetNamespace(),
		topicName.GetEncodedTopic(), "schema")

	err = s.client.get(endpoint, &response)
	if err != nil {
		return nil, err
	}

	info := utils.ConvertGetSchemaResponseToSchemaInfo(topicName, response)
	return info, nil
}

func (s *schemas) GetSchemaInfoWithVersion(topic string) (*utils.SchemaInfoWithVersion, error) {
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return nil, err
	}
	var response utils.GetSchemaResponse
	endpoint := s.client.endpoint(s.basePath, topicName.GetTenant(), topicName.GetNamespace(),
		topicName.GetEncodedTopic(), "schema")

	err = s.client.get(endpoint, &response)
	if err != nil {
		fmt.Println("err:", err.Error())
		return nil, err
	}

	info := utils.ConvertGetSchemaResponseToSchemaInfoWithVersion(topicName, response)
	return info, nil
}

func (s *schemas) GetSchemaInfoByVersion(topic string, version int64) (*utils.SchemaInfo, error) {
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return nil, err
	}

	var response utils.GetSchemaResponse
	endpoint := s.client.endpoint(s.basePath, topicName.GetTenant(), topicName.GetNamespace(), topicName.GetEncodedTopic(),
		"schema", strconv.FormatInt(version, 10))

	err = s.client.get(endpoint, &response)
	if err != nil {
		return nil, err
	}

	info := utils.ConvertGetSchemaResponseToSchemaInfo(topicName, response)
	return info, nil
}

func (s *schemas) DeleteSchema(topic string) error {
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return err
	}

	endpoint := s.client.endpoint(s.basePath, topicName.GetTenant(), topicName.GetNamespace(),
		topicName.GetEncodedTopic(), "schema")

	fmt.Println(endpoint)

	return s.client.delete(endpoint)
}

func (s *schemas) CreateSchemaByPayload(topic string, schemaPayload utils.PostSchemaPayload) error {
	topicName, err := utils.GetTopicName(topic)
	if err != nil {
		return err
	}

	endpoint := s.client.endpoint(s.basePath, topicName.GetTenant(), topicName.GetNamespace(),
		topicName.GetEncodedTopic(), "schema")

	return s.client.post(endpoint, &schemaPayload)
}
