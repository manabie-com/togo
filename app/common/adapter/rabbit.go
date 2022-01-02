package adapter

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

// Rabbits ..
type Rabbits map[string]*Rabbit

// Get : Lấy Rabbit theo name đọc từ file config
func (adapters Rabbits) Get(name string) (result *Rabbit) {
	if adapter, ok := adapters[name]; ok {
		result = adapter
		result.Err = make(chan error)
		result.mutex = sync.RWMutex{}
	} else {
		panic("Không tìm thấy config Rabbit " + name)
	}
	return
}

// Rabbit ..
type Rabbit struct {
	// Rabbit instance info, đọc từ config vào
	Name  string `mapstructure:"name"`
	Host  string `mapstructure:"host"`
	Port  string `mapstructure:"port"`
	Vhost string `mapstructure:"vhost"`
	User  string `mapstructure:"user"`
	Pass  string `mapstructure:"pass"`
	// Rabbit consumer info, đọc từ config
	Consumer RbConsumerV2 `mapstructure:"consumer"`
	// Init thêm
	Err     chan error
	channel *amqp.Channel
	conn    *amqp.Connection
	mutex   sync.RWMutex
}

// RbConsumerV2 ..
type RbConsumerV2 struct {
	Queue          string `mapstructure:"queue"`
	Exchange       string `mapstructure:"exchange"`
	RoutingKey     string `mapstructure:"routing_key"`
	Prefetch       int    `mapstructure:"prefetch"`
	DeadExchange   string `mapstructure:"dead_exchange"`
	DeadRoutingKey string `mapstructure:"dead_routing_key"`
	TTL            int    `mapstructure:"ttl"`
}

// Init :
// . Tạo mới connection
// . Tạo mới channel
func (config *Rabbit) Init() {
	config.mutex.Lock()

	log.Printf("[%s][%s] RabbitMQ [connecting]\n", config.Name, config.Host)

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Vhost)

	var err error
	// TODO: Get connection
	config.conn, err = amqp.Dial(connString)
	if err != nil {
		log.Printf("[%s][%s] RabbitMQ [error]: %s", config.Name, config.Host, err)
		time.Sleep(1 * time.Second)
		config.mutex.Unlock()
		config.Init()
		return
	}
	// TODO : Get Channel
	config.channel, err = config.conn.Channel()
	if err != nil {
		log.Printf("[%s][%s] RabbitMQ get channel [error]: %s", config.Name, config.Host, err)
		time.Sleep(1 * time.Second)
		config.mutex.Unlock()
		config.Init()
		return
	}

	// Listen to NotifyClose Connection and reconnect
	go func() {
		<-config.conn.NotifyClose(make(chan *amqp.Error))
		config.Err <- errors.New("Connection Closed. Try to re-connect")
	}()

	// Listen to NotifyClose Channel and reconnect
	go func() {
		<-config.channel.NotifyClose(make(chan *amqp.Error))
		config.Err <- errors.New("Channel Closed. Try to re-connect")
	}()

	// TODO: Get channel
	config.channel, err = config.conn.Channel()
	if err != nil {
		log.Printf("[%s][%s] RabbitMQ Channel [error]: %s", config.Name, config.Host, err)
		time.Sleep(1 * time.Second)
		config.mutex.Unlock()
		config.Init()
		return
	}

	log.Printf("[%s][%s] RabbitMQ [connected]\n", config.Name, config.Host)
	return
}

// GetConnection ..
func (config *Rabbit) GetConnection() *amqp.Connection {
	return config.conn
}

// GetChannel ..
func (config *Rabbit) GetChannel() *amqp.Channel {
	return config.channel
}

// ReInit : reconnects the connection when connection is closed
func (config *Rabbit) ReInit() (err error) {
	config.Init()

	// re prepare with new connection
	err = config.Prepare()
	if err != nil {
		return
	}

	return
}

// Prepare :
// 1. Declare Exchange, Queue
// 2. Bind queue to Exchange
func (config *Rabbit) Prepare() (err error) {
	// Validate
	{
		if config.conn == nil {
			err = fmt.Errorf("[%s][%s] Chưa init Rabbit", config.Name, config.Host)
			return
		}

		if config.channel == nil {
			err = fmt.Errorf("[%s][%s] Chưa init channel Rabbit", config.Name, config.Host)
			return
		}
	}

	// Prepare prefetch
	{
		// Set default prefetch bằng 1
		if config.Consumer.Prefetch <= 0 {
			config.Consumer.Prefetch = 1
		}
		errPrefetch := config.channel.Qos(
			config.Consumer.Prefetch, // prefetch count
			0,                        // prefetch size
			false,                    // global
		)
		if errPrefetch != nil {
			err = errPrefetch
			return
		}

	}

	// Declare Exchange
	{
		if config.Consumer.Exchange != "" {
			errExchangeDeclare := config.channel.ExchangeDeclare(
				config.Consumer.Exchange, // name
				"topic",                  // type
				true,                     // durable
				false,                    // auto-deleted
				false,                    // internal
				false,                    // no-wait
				nil,                      // arguments
			)
			if errExchangeDeclare != nil {
				err = errExchangeDeclare
				return
			}
		}
	}

	// Declare Queue
	{
		if config.Consumer.Queue != "" {
			_, errQueueDeclare := config.channel.QueueDeclare(
				config.Consumer.Queue,     // name
				true,                      // durable
				false,                     // delete when unused
				false,                     // exclusive
				false,                     // no-wait
				config.Consumer.getAgrs(), // arguments
			)
			if errQueueDeclare != nil {
				err = errQueueDeclare
				return
			}
		}
	}

	// Binding Queue
	{
		if config.Consumer.Exchange != "" && config.Consumer.Queue != "" {
			errBinding := config.channel.QueueBind(
				config.Consumer.Queue,      // queue name
				config.Consumer.RoutingKey, // routing key
				config.Consumer.Exchange,   // exchange
				false,
				nil,
			)
			if errBinding != nil {
				err = errBinding
				return
			}
		}
	}

	return
}

// PrepareByConsumer : Declare Exchange, Queue theo consumer input
// 1. Declare Exchange, Queue
// 2. Bind queue to Exchange
func (config *Rabbit) PrepareByConsumer(consumer RbConsumerV2) (err error) {
	// Validate
	{
		if config.conn == nil {
			err = fmt.Errorf("[%s][%s] Chưa init Rabbit", config.Name, config.Host)
			return
		}

		if config.channel == nil {
			err = fmt.Errorf("[%s][%s] Chưa init channel Rabbit", config.Name, config.Host)
			return
		}
	}

	// Prepare prefetch
	{
		// Set default prefetch bằng 1
		if consumer.Prefetch <= 0 {
			consumer.Prefetch = 1
		}
		errPrefetch := config.channel.Qos(
			consumer.Prefetch, // prefetch count
			0,                 // prefetch size
			false,             // global
		)
		if errPrefetch != nil {
			err = errPrefetch
			return
		}

	}

	// Declare Exchange
	{
		if consumer.Exchange != "" {
			errExchangeDeclare := config.channel.ExchangeDeclare(
				consumer.Exchange, // name
				"topic",           // type
				true,              // durable
				false,             // auto-deleted
				false,             // internal
				false,             // no-wait
				nil,               // arguments
			)
			if errExchangeDeclare != nil {
				err = errExchangeDeclare
				return
			}
		}
	}

	// Declare Queue
	{
		if consumer.Queue != "" {
			_, errQueueDeclare := config.channel.QueueDeclare(
				consumer.Queue,     // name
				true,               // durable
				false,              // delete when unused
				false,              // exclusive
				false,              // no-wait
				consumer.getAgrs(), // arguments
			)
			if errQueueDeclare != nil {
				err = errQueueDeclare
				return
			}
		}
	}

	// Binding Queue
	{
		if consumer.Exchange != "" && consumer.Queue != "" {
			errBinding := config.channel.QueueBind(
				consumer.Queue,      // queue name
				consumer.RoutingKey, // routing key
				consumer.Exchange,   // exchange
				false,
				nil,
			)
			if errBinding != nil {
				err = errBinding
				return
			}
		}
	}

	return
}

// Consume : consumes the messages from the queue
func (config *Rabbit) Consume() (msgs <-chan amqp.Delivery, err error) {
	return config.channel.Consume(
		config.Consumer.Queue, // queue
		"",                    // consumer
		false,                 // auto-ack
		false,                 // exclusive
		false,                 // no-local
		false,                 // no-wait
		nil,                   // args
	)
}

// Publish : publishes a request to the amqp queue
func (config *Rabbit) Publish(consumer *RbConsumerV2, data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// CẬp nhật đích đến, nếu muốn bắn data đến 1 consumer khác
	if consumer != nil {
		config.Consumer.Exchange = consumer.Exchange
		config.Consumer.RoutingKey = consumer.RoutingKey
	}

	return config.channel.Publish(
		config.Consumer.Exchange,   // exchange
		config.Consumer.RoutingKey, // routing key
		false,                      // mandatory
		false,                      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
}

// getAgrs : Lấy Arguments config cho consumer
func (config RbConsumerV2) getAgrs() amqp.Table {

	args := amqp.Table{}

	if config.DeadExchange != "" {
		args["x-dead-letter-exchange"] = config.DeadExchange
	}

	if config.DeadRoutingKey != "" {
		args["x-dead-letter-routing-key"] = config.DeadRoutingKey
	}

	if config.TTL > 0 {
		args["x-message-ttl"] = config.TTL
	}

	return args
}
