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

import "strconv"

type Namespaces interface {
	// Get the list of all the namespaces for a certain tenant
	GetNamespaces(tenant string) ([]string, error)

	// Get the list of all the topics under a certain namespace
	GetTopics(namespace string) ([]string, error)

	// Get the dump all the policies specified for a namespace
	GetPolicies(namespace string) (*Policies, error)

	// Creates a new empty namespace with no policies attached
	CreateNamespace(namespace string) error

	// Creates a new empty namespace with no policies attached
	CreateNsWithNumBundles(namespace string, numBundles int) error

	// Creates a new namespace with the specified policies
	CreateNsWithPolices(namespace string, polices Policies) error

	// Creates a new empty namespace with no policies attached
	CreateNsWithBundlesData(namespace string, bundleData *BundlesData) error

	// Delete an existing namespace
	DeleteNamespace(namespace string) error

	// Delete an existing bundle in a namespace
	DeleteNamespaceBundle(namespace string, bundleRange string) error

	// Set the messages Time to Live for all the topics within a namespace
	SetNamespaceMessageTTL(namespace string, ttlInSeconds int) error

	// Get the message TTL for a namespace
	GetNamespaceMessageTTL(namespace string) (int, error)

	// Get the retention configuration for a namespace
	GetRetention(namespace string) (*RetentionPolicies, error)

	// Set the retention configuration for all the topics on a namespace
	SetRetention(namespace string, policy RetentionPolicies) error

	// Get backlog quota map on a namespace
	GetBacklogQuotaMap(namespace string) (map[BacklogQuotaType]BacklogQuota, error)

	// Set a backlog quota for all the topics on a namespace
	SetBacklogQuota(namespace string, backlogQuota BacklogQuota) error

	// Remove a backlog quota policy from a namespace
	RemoveBacklogQuota(namespace string) error

	// Set maxConsumersPerSubscription for a namespace.
	SetMaxConsumersPerSubscription(namespace NameSpaceName, max int) error

	// Get the maxConsumersPerSubscription for a namespace.
	GetMaxConsumersPerSubscription(namespace NameSpaceName) (int, error)

	// Set maxConsumersPerTopic for a namespace.
	SetMaxConsumersPerTopic(namespace NameSpaceName, max int) error

	// Get the maxProducersPerTopic for a namespace.
	GetMaxConsumersPerTopic(namespace NameSpaceName) (int, error)

	// Set maxProducersPerTopic for a namespace.
	SetMaxProducersPerTopic(namespace NameSpaceName, max int) error

	// Get the maxProducersPerTopic for a namespace.
	GetMaxProducersPerTopic(namespace NameSpaceName) (int, error)
}

type namespaces struct {
	client   *client
	basePath string
}

func (c *client) Namespaces() Namespaces {
	return &namespaces{
		client:   c,
		basePath: "/namespaces",
	}
}

func (n *namespaces) GetNamespaces(tenant string) ([]string, error) {
	var namespaces []string
	endpoint := n.client.endpoint(n.basePath, tenant)
	err := n.client.get(endpoint, &namespaces)
	return namespaces, err
}

func (n *namespaces) GetTopics(namespace string) ([]string, error) {
	var topics []string
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return nil, err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String(), "topics")
	err = n.client.get(endpoint, &topics)
	return topics, err
}

func (n *namespaces) GetPolicies(namespace string) (*Policies, error) {
	var police Policies
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return nil, err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String())
	err = n.client.get(endpoint, &police)
	return &police, err
}

func (n *namespaces) CreateNsWithNumBundles(namespace string, numBundles int) error {
	return n.CreateNsWithBundlesData(namespace, NewBundlesDataWithNumBundles(numBundles))
}

func (n *namespaces) CreateNsWithPolices(namespace string, policies Policies) error {
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String())
	return n.client.put(endpoint, &policies, nil)
}

func (n *namespaces) CreateNsWithBundlesData(namespace string, bundleData *BundlesData) error {
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String())
	polices := new(Policies)
	polices.Bundles = bundleData

	return n.client.put(endpoint, &polices, nil)
}

func (n *namespaces) CreateNamespace(namespace string) error {
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String())
	return n.client.put(endpoint, nil, nil)
}

func (n *namespaces) DeleteNamespace(namespace string) error {
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String())
	return n.client.delete(endpoint, nil)
}

func (n *namespaces) DeleteNamespaceBundle(namespace string, bundleRange string) error {
	ns, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, ns.String(), bundleRange)
	return n.client.delete(endpoint, nil)
}

func (n *namespaces) GetNamespaceMessageTTL(namespace string) (int, error) {
	var ttl int
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return 0, err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "messageTTL")
	err = n.client.get(endpoint, &ttl)
	return ttl, err
}

func (n *namespaces) SetNamespaceMessageTTL(namespace string, ttlInSeconds int) error {
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}

	endpoint := n.client.endpoint(n.basePath, nsName.String(), "messageTTL")
	return n.client.post(endpoint, &ttlInSeconds, nil)
}

func (n *namespaces) SetRetention(namespace string, policy RetentionPolicies) error {
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "retention")
	return n.client.post(endpoint, &policy, nil)
}

func (n *namespaces) GetRetention(namespace string) (*RetentionPolicies, error) {
	var policy RetentionPolicies
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return nil, err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "retention")
	err = n.client.get(endpoint, &policy)
	return &policy, err
}

func (n *namespaces) GetBacklogQuotaMap(namespace string) (map[BacklogQuotaType]BacklogQuota, error) {
	var backlogQuotaMap map[BacklogQuotaType]BacklogQuota
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return nil, err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "backlogQuotaMap")
	err = n.client.get(endpoint, &backlogQuotaMap)
	return backlogQuotaMap, err
}

func (n *namespaces) SetBacklogQuota(namespace string, backlogQuota BacklogQuota) error {
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "backlogQuota")
	return n.client.post(endpoint, &backlogQuota, nil)
}

func (n *namespaces) RemoveBacklogQuota(namespace string) error {
	nsName, err := GetNamespaceName(namespace)
	if err != nil {
		return err
	}
	endpoint := n.client.endpoint(n.basePath, nsName.String(), "backlogQuota")
	params := map[string]string{
		"backlogQuotaType": string(DestinationStorage),
	}
	return n.client.deleteWithQueryParams(endpoint, nil, params)
}

func (n *namespaces) SetMaxConsumersPerSubscription(namespace NameSpaceName, max int) error {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxConsumersPerSubscription")
	return n.client.post(endpoint, max, nil)
}

func (n *namespaces) GetMaxConsumersPerSubscription(namespace NameSpaceName) (int, error) {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxConsumersPerSubscription")
	b, err := n.client.getAndDecode(endpoint, nil, false)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(b))
}

func (n *namespaces) SetMaxConsumersPerTopic(namespace NameSpaceName, max int) error {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxConsumersPerTopic")
	return n.client.post(endpoint, max, nil)
}

func (n *namespaces) GetMaxConsumersPerTopic(namespace NameSpaceName) (int, error) {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxConsumersPerTopic")
	b, err := n.client.getAndDecode(endpoint, nil, false)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(b))
}

func (n *namespaces) SetMaxProducersPerTopic(namespace NameSpaceName, max int) error {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxProducersPerTopic")
	return n.client.post(endpoint, max, nil)
}

func (n *namespaces) GetMaxProducersPerTopic(namespace NameSpaceName) (int, error) {
	endpoint := n.client.endpoint(n.basePath, namespace.String(), "maxProducersPerTopic")
	b, err := n.client.getAndDecode(endpoint, nil, false)
	if err != nil {
		return -1, err
	}
	return strconv.Atoi(string(b))
}
