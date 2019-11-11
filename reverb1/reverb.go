package reverb

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {

		echo(c, input.Text(), 1*time.Second)    // reverb1
		go echo(c, input.Text(), 1*time.Second) // reverb2

	}
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}
