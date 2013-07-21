package main

import (
	"poll"
	"fmt"
	"net"
	"time"
)

var noDeadLine = time.Time{}
var listenerBacklog = 50

func main() {

	fmt.Println("begin")

	poll, err := poll.NewPoll()
	if err != nil {
		fmt.Println("new poll error")
		return
	}
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":9070")
	if err != nil {
		fmt.Println("tcp4 resolve tcp addr error")
		return
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("listen error")
		return
	}

	poll.AddFD(listener.fd, 'r', true)

	for {
		result, events, err := poll.FetchEvent(0)
		if err != nil {
			return
		}
		if result < 0 {
			return
		}
		if result == 0 {
			continue
		}
		for _, ev := range events {
			fmt.Println(ev.fd)
			fmt.Println(ev.data)
		}
	}
}
