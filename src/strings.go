package redis

func (c *Client) Set(key, value string) {
	c.writeReadOk("SET", key, value)
}

func (c *Client) Get(key string) []byte {
	return c.writeReadBulk("GET", key)
}

func (c *Client) StrLen(key string) int {
	return c.writeReadInt("STRLEN", key)
}

func (c *Client) Incr(key string) int {
	return c.writeReadInt("INCR", key)
}

func (c *Client) Append(key string, value string) int {
	return c.writeReadInt("APPEND", key, value)
}

