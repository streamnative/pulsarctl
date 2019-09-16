package pulsar

import (
	"net/url"
	"strconv"
)

type Subscriptions interface {
	Create(TopicName, string, MessageId) error
	Delete(TopicName, string) error
	List(TopicName) ([]string, error)
	ResetCursorWithMessageId(TopicName, string, MessageId) error
	ResetCursorWithTimestamp(TopicName, string, int64) error
	ClearBacklog(TopicName, string) error
	SkipMessages(TopicName, string, int64) error
	PeekMessages(TopicName, string, int) error
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

func (s *subscriptions) Create(topic TopicName, sName string, messageId MessageId) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName))
	return s.client.put(endpoint, messageId, nil)
}

func (s *subscriptions) Delete(topic TopicName, sName string) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName))
	return s.client.delete(endpoint, nil)
}

func (s *subscriptions) List(topic TopicName) ([]string, error) {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), "subscriptions")
	var list []string
	return list, s.client.get(endpoint, &list)
}

func (s *subscriptions) ResetCursorWithMessageId(topic TopicName, sName string, id MessageId) error {
	endpoint := s.client.endpoint(s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName), "resetcursor")
	return s.client.post(endpoint, id, nil)
}

func (s *subscriptions) ResetCursorWithTimestamp(topic TopicName, sName string, timestamp int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"resetcursor", strconv.FormatInt(timestamp, 10))
	return s.client.post(endpoint, "", nil)
}

func (s *subscriptions) ClearBacklog(topic TopicName, sName string) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName), "skip_all")
	return s.client.post(endpoint, "", nil)
}

func (s *subscriptions) SkipMessages(topic TopicName, sName string, n int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"skip", strconv.FormatInt(n, 10))
	return s.client.post(endpoint, "", nil)
}

func (s *subscriptions) PeekMessages(topic TopicName, sName string, n int) error {
	//endpoint := s.client.endpoint(
	//	s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
	//	"position", strconv.Itoa(n))
	return nil
}

func (s *subscriptions) peekNthMessages(topic TopicName, sName string, pos int) error {
	//endpoint := s.client.endpoint(
	//	s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
	//	"position", strconv.Itoa(pos))
	return nil
}

func (s *subscriptions) ExpireMessages(topic TopicName, sName string, expire int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), s.SubPath, url.QueryEscape(sName),
		"expireMessages", strconv.FormatInt(expire, 10))
	return s.client.post(endpoint, "", nil)
}

func (s *subscriptions) ExpireAllMessages(topic TopicName, expire int64) error {
	endpoint := s.client.endpoint(
		s.basePath, topic.GetRestPath(), "all_subscription",
		"expireMessages", strconv.FormatInt(expire, 10))
	return s.client.post(endpoint, "", nil)
}
