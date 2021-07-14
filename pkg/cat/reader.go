package cat

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/apache/pulsar-client-go/pulsar"
	"io"
	"time"
	"unicode/utf8"
)

type ReaderFormat int

const (
	PayloadOnlyAsString ReaderFormat = iota
	PayloadOnlyAsHex
	PayloadOnlyAsBase64
	MetadataAsJson
)

var ReaderFormatIds = map[ReaderFormat][]string{
	PayloadOnlyAsString: {"payload-string"},
	PayloadOnlyAsHex:    {"payload-hex"},
	PayloadOnlyAsBase64: {"payload-base64"},
	MetadataAsJson:      {"full-message"},
}

type Reader struct {
	client     pulsar.Client
	reader     pulsar.Reader
	opts       ReaderOpts
	readMsgs   int
	partialMsg []byte
	finished   bool
	closed     bool
}

type ReaderMessage struct {
	Topic           string            `json:"topic"`
	ProducerName    string            `json:"producer_name"`
	Properties      map[string]string `json:"properties,omitempty"`
	PayloadString   string            `json:"payload,omitempty"`
	PayloadBytes    []byte            `json:"payloadRaw,omitempty"`
	ID              string            `json:"id"`
	PublishTime     time.Time         `json:"publish_time"`
	EventTime       time.Time         `json:"event_time,omitempty"`
	Key             string            `json:"key,omitempty"`
	OrderingKey     string            `json:"ordering_key,omitempty"`
	RedeliveryCount uint32            `json:"redelivery_count,omitempty"`
	IsReplicated    bool              `json:"is_replicated,omitempty"`
	ReplicatedFrom  string            `json:"replicated_from,omitempty"`
}

func FromPulsarMessage(msg pulsar.Message) *ReaderMessage {
	payloadString := ""
	payloadBytes := make([]byte, 0)
	if utf8.Valid(msg.Payload()) {
		payloadString = string(msg.Payload())
	} else {
		payloadBytes = msg.Payload()
	}
	return &ReaderMessage{
		Topic:           msg.Topic(),
		ProducerName:    msg.ProducerName(),
		Properties:      msg.Properties(),
		PayloadString:   payloadString,
		PayloadBytes:    payloadBytes,
		ID:              msg.ID().String(),
		PublishTime:     msg.PublishTime(),
		EventTime:       msg.EventTime(),
		Key:             msg.Key(),
		OrderingKey:     msg.OrderingKey(),
		RedeliveryCount: msg.RedeliveryCount(),
		IsReplicated:    msg.IsReplicated(),
		ReplicatedFrom:  msg.GetReplicatedFrom(),
	}
}

type ReaderOpts struct {
	Topic     string
	StartAt   pulsar.MessageID
	Format    ReaderFormat
	Compacted bool
	MsgCount  int
	Tailing   bool
}

func NewReader(client pulsar.Client, opts ReaderOpts) (*Reader, error) {
	pulsarReader, err := client.CreateReader(pulsar.ReaderOptions{
		Topic:          opts.Topic,
		StartMessageID: opts.StartAt,
		ReadCompacted: opts.Compacted,
		StartMessageIDInclusive: true,
		ReceiverQueueSize: 10000,
	})
	if err != nil {
		return nil, err
	}
	reader := &Reader{client: client, reader: pulsarReader, opts: opts, readMsgs: 0, partialMsg: make([]byte, 0), finished: false, closed: false}
	return reader, nil
}

func (r *Reader) Read(p []byte) (n int, err error) {
	if r.closed {
		return 0, io.EOF
	}
	if r.finished && !r.opts.Tailing {
		return 0, io.EOF
	}
	toRead := len(p)
	readSoFar := 0
	if len(r.partialMsg) > 0 {
		copied := copy(p, r.partialMsg)
		if copied < len(r.partialMsg) {
			r.partialMsg = r.partialMsg[copied:]
		} else {
			r.partialMsg = make([]byte, 0)
		}
		toRead -= copied
		readSoFar += copied
	}
	reachedMsgCount := false
	if r.opts.MsgCount != -1 && r.readMsgs >= r.opts.MsgCount {
		reachedMsgCount = true
	}
	if !reachedMsgCount {
		for toRead > 0 && !r.closed {
			ctx, _ := context.WithTimeout(context.Background(), time.Millisecond*100)
			msg, err := r.reader.Next(ctx)
			if msg == nil {
				continue
			}
			if err != nil {
				return len(p) - toRead, err
			}
			r.readMsgs += 1
			buff, err := r.formatMsg(msg)
			if err != nil {
				return len(p) - toRead, err
			}
			var copied int
			if toRead < len(buff) {
				copied = copy(p[readSoFar:], buff[:toRead])
				r.partialMsg = buff[toRead:]
			} else {
				copied = copy(p[readSoFar:], buff)
			}
			readSoFar += copied
			toRead -= copied
		}
	} else {
		r.finished = true
	}
	return readSoFar, nil
}

func (r *Reader) formatMsg(msg pulsar.Message) ([]byte, error) {
	switch r.opts.Format {
	case PayloadOnlyAsString:
		if utf8.Valid(msg.Payload()) {
			return append(msg.Payload(), []byte("\n")...), nil
		} else {
			return []byte{}, errors.New("invalid utf8 string")
		}
	case PayloadOnlyAsBase64:
		return []byte(base64.StdEncoding.EncodeToString(msg.Payload())), nil
	case PayloadOnlyAsHex:
		return []byte(hex.EncodeToString(msg.Payload())), nil
	case MetadataAsJson:
		meta := FromPulsarMessage(msg)
		msg, err := json.Marshal(meta)
		if err != nil {
			return []byte{}, err
		}
		return append(msg, []byte("\n")...), err
	default:
		return nil, errors.New("unrecognized format")
	}
}

func (r *Reader) ReadStats() string {
	// TODO improve the stats
	return fmt.Sprintf("Topic: %v, messagesRead %v", r.opts.Topic, r.readMsgs)
}

func (r *Reader) Close() error {
	r.closed = true
	r.reader.Close()
	return nil
}
