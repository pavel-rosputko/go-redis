package redis

import "testing"

var client Client

const (
	key = "key"
	value = "value"
)

func TestExist(t *testing.T) {
	client.Set(key, value)

	r := client.Exists(key)
	if !r { t.Fatal("") }
}
