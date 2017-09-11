package nsq

import (
	"encoding/base64"
	"log"

	"github.com/nsqio/go-nsq"
)

//分发器
type Dispatcher struct {
	topic  string
	handle func(string, interface{})
}

type ConsumerDispatcher struct {
	F *Dispatcher
	C *nsq.Consumer
}

//主题发现
type TopicDiscoverer struct {
	topics map[string]*ConsumerDispatcher
	cfg    *nsq.Config
}

func NewTopicDiscoverer(topics []string, maxInFlight int, lookupdHTTPAddrs []string, handle func(string, interface{})) *TopicDiscoverer {
	var discoverer TopicDiscoverer
	discoverer.topics = make(map[string]*ConsumerDispatcher)

	cfg := nsq.NewConfig()
	cfg.UserAgent = "router v0.1"
	cfg.MaxInFlight = maxInFlight

	discoverer.cfg = cfg
	for _, v := range topics {
		cd, err := NewConsumerDispatcher(v, cfg, lookupdHTTPAddrs, handle)
		if err != nil {
			log.Fatal(err)
		}
		discoverer.topics[v] = cd
	}
	return &discoverer
}

func NewConsumerDispatcher(topic string, cfg *nsq.Config, lookupdHTTPAddrs []string, handle func(string, interface{})) (*ConsumerDispatcher, error) {
	var err error
	r := NewDispatcher(topic, handle)
	if err != nil {
		return nil, err
	}

	consumer, err := nsq.NewConsumer(topic, "channel-to-router", cfg)
	if err != nil {
		return nil, err
	}
	consumer.AddHandler(r)
	err = consumer.ConnectToNSQLookupds(lookupdHTTPAddrs)
	if err != nil {
		log.Fatal(err)
	}
	return &ConsumerDispatcher{
		C: consumer,
		F: r,
	}, nil
}

func NewDispatcher(topic string, handle func(string, interface{})) *Dispatcher {
	var dispatcher Dispatcher
	dispatcher.topic = topic
	dispatcher.handle = handle
	return &dispatcher
}

func (p *Dispatcher) HandleMessage(msg *nsq.Message) error {
	body, err := base64.StdEncoding.DecodeString(string(msg.Body[:]))
	if err != nil {
		log.Println("err=", err, " topic=", p.topic, " msg=", string(msg.Body))
		p.handle(p.topic, msg.Body)
		return nil
	}
	p.handle(p.topic, body)
	return nil
}

type Producer struct {
	producer *nsq.Producer
}

func NewProducer(destNsqdTCPAddrs string) (*Producer, error) {
	var p Producer
	cfg := nsq.NewConfig()
	cfg.UserAgent = "router v0.1"
	producer, err := nsq.NewProducer(destNsqdTCPAddrs, cfg)
	if err != nil {
		panic(err)
		return &p, err
	}
	p.producer = producer
	return &p, err
}

func (p *Producer) Publish(topic string, msg []byte) error {
	m := base64.StdEncoding.EncodeToString(msg)
	return p.producer.Publish(topic, []byte(m))
}
