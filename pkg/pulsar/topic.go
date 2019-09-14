package pulsar

import (
	"strconv"
)

type Topics interface {
	Create(TopicName, int) error
	Delete(TopicName, bool, bool) error
	Update(TopicName, int) error
	GetMetadata(TopicName) (PartitionedTopicMetadata, error)
	List(NameSpaceName) ([]string, []string, error)
	GetStats(TopicName) (TopicStats, error)
	GetInternalStats(TopicName) (PersistentTopicInternalStats, error)
	GetPartitionedStats(TopicName, bool) (PartitionedTopicStats, error)
	Compact(TopicName) error
	CompactStatus(TopicName) (LongRunningProcessStatus, error)
	Unload(TopicName) error
	Offload(TopicName, MessageId) error
	OffloadStatus(TopicName) (OffloadProcessStatus, error)
	Terminate(TopicName) (MessageId, error)
}

type topics struct {
	client            *client
	basePath          string
	persistentPath    string
	nonPersistentPath string
}

func (c *client) Topics() Topics {
	return &topics{
		client:            c,
		basePath:          "",
		persistentPath:    "/persistent",
		nonPersistentPath: "/non-persistent",
	}
}

func (t *topics) Create(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	if partitions == 0 {
		endpoint = t.client.endpoint(t.basePath, topic.GetRestPath())
	}
	return t.client.put(endpoint, partitions, nil)
}

func (t *topics) Delete(topic TopicName, force bool, nonPartitioned bool) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	if nonPartitioned {
		endpoint = t.client.endpoint(t.basePath, topic.GetRestPath())
	}
	params := map[string]string{
		"force": strconv.FormatBool(force),
	}
	return t.client.deleteWithQueryParams(endpoint, nil, params)
}

func (t *topics) Update(topic TopicName, partitions int) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	return t.client.post(endpoint, partitions, nil)
}

func (t *topics) GetMetadata(topic TopicName) (PartitionedTopicMetadata, error) {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitions")
	var partitionedMeta PartitionedTopicMetadata
	err := t.client.get(endpoint, &partitionedMeta)
	return partitionedMeta, err
}

func (t *topics) List(namespace NameSpaceName) ([]string, []string, error) {
	var partitionedTopics, nonPartitionedTopics []string
	partitionedTopicsChan := make(chan []string)
	nonPartitionedTopicsChan := make(chan []string)
	errChan := make(chan error)

	pp := t.client.endpoint(t.persistentPath, namespace.String(), "partitioned")
	np := t.client.endpoint(t.nonPersistentPath, namespace.String(), "partitioned")
	p := t.client.endpoint(t.persistentPath, namespace.String())
	n := t.client.endpoint(t.nonPersistentPath, namespace.String())

	go t.getTopics(pp, partitionedTopicsChan, errChan)
	go t.getTopics(np, partitionedTopicsChan, errChan)
	go t.getTopics(p, nonPartitionedTopicsChan, errChan)
	go t.getTopics(n, nonPartitionedTopicsChan, errChan)

	requestCount := 4
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return nil, nil, err
			}
			continue
		case pTopic := <-partitionedTopicsChan:
			requestCount--
			partitionedTopics = append(partitionedTopics, pTopic...)
		case npTopic := <-nonPartitionedTopicsChan:
			requestCount--
			nonPartitionedTopics = append(nonPartitionedTopics, npTopic...)
		}
		if requestCount == 0 {
			break
		}
	}
	return partitionedTopics, nonPartitionedTopics, nil
}

func (t *topics) getTopics(endpoint string, out chan<- []string, err chan<- error) {
	var topics []string
	err <- t.client.get(endpoint, &topics)
	out <- topics
}

func (t *topics) GetStats(topic TopicName) (TopicStats, error) {
	var stats TopicStats
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "stats")
	err := t.client.get(endpoint, &stats)
	return stats, err
}

func (t *topics) GetInternalStats(topic TopicName) (PersistentTopicInternalStats, error) {
	var stats PersistentTopicInternalStats
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "internalStats")
	err := t.client.get(endpoint, &stats)
	return stats, err
}

func (t *topics) GetPartitionedStats(topic TopicName, perPartition bool) (PartitionedTopicStats, error) {
	var stats PartitionedTopicStats
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "partitioned-stats")
	params := map[string]string{
		"perPartition": strconv.FormatBool(perPartition),
	}
	err := t.client.getWithQueryParams(endpoint, &stats, params)
	return stats, err
}

func (t *topics) Compact(topic TopicName) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "compaction")
	return t.client.put(endpoint, "", nil)
}

func (t *topics) CompactStatus(topic TopicName) (LongRunningProcessStatus, error) {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "compaction")
	var status LongRunningProcessStatus
	err := t.client.get(endpoint, &status)
	return status, err
}

func (t *topics) Unload(topic TopicName) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "unload")
	return t.client.put(endpoint, "", nil)
}

func (t *topics) Offload(topic TopicName, messageId MessageId) error {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "offload")
	return t.client.put(endpoint, messageId, nil)
}

func (t *topics) OffloadStatus(topic TopicName) (OffloadProcessStatus, error) {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "offload")
	var status OffloadProcessStatus
	err := t.client.get(endpoint, &status)
	return status, err
}

func (t *topics) Terminate(topic TopicName) (MessageId, error) {
	endpoint := t.client.endpoint(t.basePath, topic.GetRestPath(), "terminate")
	var messageId MessageId
	err := t.client.post(endpoint, "", &messageId)
	return messageId, err
}
