package redis

import "strconv"

func (c *Client) HDel(key, field string) bool {
	return c.writeReadInt("HDEL", key, field) == 1
}

func (c *Client) HExists(key, field string) bool {
	return c.writeReadInt("HEXISTS", key, field) == 1
}

func (c *Client) HGet(key, field string) []byte {
	return c.writeReadBulk("HGET", key, field)
}

func (c *Client) HGetAll(key string) [][]byte {
	return c.writeReadMultiBulk("HGETALL", key)
}

func (c *Client) HIncrBy(key, field string, incr int) int {
	return c.writeReadInt("HINCRBY", key, field, strconv.Itoa(incr))
}

func (c *Client) HKeys(key string) [][]byte {
	return c.writeReadMultiBulk("HKEYS", key)
}

func (c *Client) HLen(key string) int {
	return c.writeReadInt("HLEN", key)
}

// NOTE key string, fields ...string
func (c *Client) HMGet(args ...string) [][]byte {
	return c.writeReadMultiBulk("HMGET", args...)
}

// NOTE key string, args ...string
func (c *Client) HMSet(args ...string) [][]byte {
	return c.writeReadMultiBulk("HMSET", args...)
}

func (c *Client) HSet(key, field, value string) bool {
	return c.writeReadInt("HSET", key, field, value) == 1
}

func (c *Client) HSetNX(key, field, value string) bool {
	return c.writeReadInt("HSETNX", key, field, value) == 1
}

func (c *Client) HVals(key string) [][]byte {
	return c.writeReadMultiBulk("HVALS", key)
}
