package redis

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
)

type Client struct {
	Addr	string
	conn	net.Conn
	reader	*bufio.Reader
}

func (c *Client) connect() {
	addr := c.Addr
	if addr == "" { addr = "localhost:6379" }
	conn, err := net.Dial("tcp", "", addr)
	if err != nil { panic(err) }
	c.conn = conn
	c.reader = bufio.NewReader(conn)
}

func (c *Client) readLine() string {
	line, err := c.reader.ReadString('\n')
	if err != nil { panic(err) }
	// println("readLine: ", string(line[0:len(line) - 2]))
	return line[0:len(line) - 2]
}

func (c *Client) readSingle(single string) {
	line := c.readLine()
	if line[0] != '+' || line[1:] != single { panic("not +" + single) }
}

func atoi(s string) int {
	r, e := strconv.Atoi(s)
	if e != nil { panic(e) }
	return r
}

func (c *Client) readInt() int {
	line := c.readLine()
	if line[0] != ':' { panic("not :") }
	return atoi(line[1:])
}

func (c *Client) readBulk() []byte {
	line := c.readLine()
	if line[0] != '$' { panic("not $") }

	size := atoi(line[1:])
	if size == -1 { return nil }

	lr := io.LimitReader(c.reader, int64(size + 2))

	bytes, err := ioutil.ReadAll(lr)
	if err != nil { panic(err) }
	return bytes[0:size]
}

func (c *Client) readMultiBulk() [][]byte {
	line := c.readLine()
	if line[0] != '*' { panic("not *") }

	count := atoi(line[1:])
	result := make([][]byte, count)
	for i := 0; i < count; i++ { result[i] = c.readBulk() }
	return result
}

func (c *Client) write(cmd string, args []string) {
	b := bytes.NewBufferString(fmt.Sprintf("*%d\r\n$%d\r\n%s\r\n", len(args) + 1, len(cmd), cmd))
	for _, arg := range args {
		b.WriteString(fmt.Sprintf("$%d\r\n%s\r\n", len(arg), arg))
	}

	// println("write:", string(c.buffer.Bytes()))

	// buffer.WriteTo(c.conn)

	if c.conn == nil { c.connect() }

	c.conn.Write(b.Bytes())
}

func (c *Client) writeRead(cmd string, args []string, read func()) {
	defer func() {
		e := recover()
		if e == os.EOF {
			c.connect()
			c.write(cmd, args)
			read()
		} else if e != nil {
			panic(e)
		}
	}()

	c.write(cmd, args)
	read()
}

func (c *Client) writeReadInt(cmd string, args ...string) (r int) {
	c.writeRead(cmd, args, func() { r = c.readInt() })
	return
}

func (c *Client) writeReadBulk(cmd string, args ...string) (r []byte) {
	c.writeRead(cmd, args, func() { r = c.readBulk() })
	return
}

func (c *Client) writeReadMultiBulk(cmd string, args ...string) (r [][]byte) {
	c.writeRead(cmd, args, func() { r = c.readMultiBulk() })
	return
}

func (c *Client) writeReadSingle(cmd string, single string, args ...string) {
	c.writeRead(cmd, args, func() { c.readSingle(single) })
}

