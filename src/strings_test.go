package redis

import "testing"

var client Client

const (
	key = "key"
	value = "value"
)

func TestGet(t *testing.T) {
	client.Set(key, value)

	res := client.Get(key)

	if string(res) != value { t.Fatal("") }
}

/* 
func TestSetStrLen(t *testing.T) {
	client.Set(key, value)
	if client.StrLen(key) != len(value) { t.Fatal("") }
} */

func TestIncr(t *testing.T) {
	client.Set(key, "10")
	r := client.Incr(key)
	if r != 11 { t.Fatal("") }
}

func TestAppend(t *testing.T) {
	client.Set(key, value)
	r := client.Append(key, value)

	if r != len(value + value) { t.Fatal("") }

	res := client.Get(key)
	if string(res) != (value + value) { t.Fatal("") }
}

func BenchmarkStrings(bm *testing.B) {
	for i := 0; i < bm.N; i++ {
		client.Set("key", "value")
	}
}
