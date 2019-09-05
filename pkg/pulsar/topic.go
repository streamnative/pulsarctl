package pulsar

import (
	. "github.com/streamnative/pulsarctl/pkg/pulsar/common"
	"strconv"
)

type Topics interface {
	CreatePartitionedTopic(TopicName, int) error
	DeletePartitionedTopic(TopicName, bool) error
	UpdatePartitionedTopic(TopicName, int) error
	GetPartitionedTopicMeta(TopicName) (PartitionedTopicMetadata, error)
	ListPartitionedTopic(NameSpaceName) ([]string, error)
}

type topics struct {
	client   *client
	basePath string
}

func (c *client) Topics() Topics {
	return &topics{
		client:   c,
		basePath: "",
	}
}

func (t *topics) CreatePartitionedTopic(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	return t.client.put(endpoint, partitions, nil)
}

func (t *topics) DeletePartitionedTopic(topic TopicName, force bool) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	params := map[string]string{
		"force": strconv.FormatBool(force),
	}
	return t.client.deleteWithQueryParams(endpoint, nil, params)
}

func (t *topics) UpdatePartitionedTopic(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	return t.client.post(endpoint, &PartitionedTopicMetadata{partitions}, nil)
}

func (t *topics) GetPartitionedTopicMeta(topic TopicName) (PartitionedTopicMetadata, error) {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	var partitionedMeta PartitionedTopicMetadata
	err := t.client.get(endpoint, &partitionedMeta)
	return partitionedMeta, err
}

func (t *topics) ListPartitionedTopic(namespace NameSpaceName) ([]string, error) {
	var persistentTopics []string
	persistent := t.client.endpoint(t.basePath, "persistent", namespace.String(), "partitioned")
	err := t.client.get(persistent, &persistentTopics)
	if err != nil {
		return nil, err
	}

	var nonPersistentTopics []string
	nonPersistent := t.client.endpoint(t.basePath, "non-persistent", namespace.String(), "partitioned")
	err = t.client.get(nonPersistent, &nonPersistentTopics)
	if err != nil {
		return nil, err
	}

	return append(persistentTopics, nonPersistentTopics...), nil
}
