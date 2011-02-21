package redis

import "testing"

var client Client

const (
	key = "key"
	field = "field"
	value = "value"
)

func TestHGet(t *testing.T) {
	client.Del(key)

	client.HSet(key, field, value)

	result := client.HGet(key, field)

	if string(result) != value { t.Fatal("") }
}

func TestHGetAll(t *testing.T) {
	client.Del(key)
	client.HSet(key, field, value)

	result := client.HGetAll(key)

	if len(result) != 2 { t.Fatal("") }
	if string(result[0]) != field { t.Fatal("") }
	if string(result[1]) != value { t.Fatal("") }
}
