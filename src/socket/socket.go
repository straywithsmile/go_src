package main

import (
	"fmt"
	"net"
	"strings"
	"os"
	"code.google.com/p/mahonia"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println(Show("input param is error"))
		os.Exit(0)
	}
	param := strings.ToLower(args[1])
	if param == "server" {
		Server()
	} else if param == "client" {
		Client()
	} else {
		fmt.Println("error param");
		os.Exit(0)
	}
}

func Show(s string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(s)
}

func Server() {
	exit := make(chan bool);
	ip := net.ParseIP("127.0.0.1")
	addr := net.TCPAddr{ip, 8888, ""}
	go func() {
		listen, err := net.ListenTCP("tcp", &addr)
		if err != nil {
			fmt.Println(Show("initial error"), Show(err.Error()))
			exit <- true
			return
		}
		fmt.Println(Show("listening..."))
		for {
			client, err := listen.AcceptTCP()
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(Show("client connect"), Show(client.RemoteAddr().String()))
				data := make([]byte, 1024)
				c, err := client.Read(data)
				if err != nil {
					fmt.Println(Show(err.Error()))
				}
			fmt.Println(Show(string(data[0:c])))
			client.Write([]byte("hello client!\r\n"))
			client.Close()
		}
	}();
	<- exit;
	fmt.Println(Show("server closed"))
}

func Client() {
	client, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println(Show("server connect fail"), Show(err.Error()))
		return
	}
	defer client.Close()
	client.Write([]byte("hello, server!"))
	buf := make([]byte, 1024)
	c, err := client.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(Show(string(buf[0:c])))
}
