package redis

import "strconv"

func (c *Client) LIndex(key string, index int) (string, bool) {
	c.send("LINDEX", key, strconv.Itoa(index))
	return c.nextBulk()
}

// func (c *Client) LInsert(

func (c *Client) LLen(key string) int {
	c.send("LLEN", key)
	return c.nextInteger()
}

func (c *Client) LPop(key string) (string, bool) {
	c.send("LPOP", key)
	return c.nextBulk()
}

func (c *Client) LPush(key, value string) {
	c.send("LPUSH", key, value)
	return c.nextInteger()
}

func (c *Client) LSet(key string, index int, value string) {
	c.send("LSET", key, strconv.Itoa(index), value)
	c.nextOk()
}

func (c *Client) RPop(key string) (string, bool) {
	c.send("RPOP", key)
	return c.nextBulk()
}

func (c *Client) RPush(key, value string) int {
	c.send("RPUSH", key, value)
	return c.nextInteger()
}
