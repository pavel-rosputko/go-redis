package redis

import (
	"testing"
	"time"
)

var client Client

func Test1(t *testing.T) {
	client.Set("key", "value")
	println(string(client.Get("key")))

	time.Sleep(10 * 1000000000)

	client.Set("key", "value")
	println(string(client.Get("key")))
}
