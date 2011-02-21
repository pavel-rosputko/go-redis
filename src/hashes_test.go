package redis

import "fmt"
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

const keysCount = 10000
const fieldsCount = 10

func BenchmarkHash(bm *testing.B) {
	for i := 0; i < bm.N; i++ {
		client.Del(key)

		for j := 0; j < fieldsCount; j++ {
			client.HSet(key, fmt.Sprintf("field%d", j), fmt.Sprintf("value%d", j))
		}

		result := client.HGetAll(key)
		if len(result) != 2 * fieldsCount { println("fatal") }
	}
}
