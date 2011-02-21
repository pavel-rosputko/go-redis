package redis

func (c *Client) Exists(key string) bool {
	return c.writeReadInt("EXISTS", key) == 1
}

func (c *Client) Del(keys ...string) int {
	return c.writeReadInt("DEL", keys...)
}
