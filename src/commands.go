package redis

func (c *Client) Ping() {
	c.writeReadSingle("PING", "PONG")
}
