package cache

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
)

const defaultPrefix = "fnb_"

var once sync.Once
var client *Client

// Client struct act as function receiver for cache
type Client struct {
	prefix string
	rdc    *redis.Client
}

//Init function setup key prefix and init the redis client
func Init(prefix string) {
	once.Do(func() {
		rdc := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})

		if prefix == "" {
			prefix = defaultPrefix
		}

		client = &Client{
			prefix: prefix,
			rdc:    rdc,
		}

		_, err := client.rdc.Ping().Result()
		if err != nil {
			log.Fatalf("Could not connect to redis %v", err)
		}

	})
}

// Get singletoned client
func Get() *Client {
	return client
}

// Set value with given key
func (c *Client) Set(key string, value interface{}) error {
	return c.rdc.Set(c.prefix+key, value, 0).Err()
}

// Get value with give key
func (c *Client) Get(key string) (string, error) {
	return c.rdc.Get(c.prefix + key).Result()
}
