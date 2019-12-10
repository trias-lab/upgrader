package main

import (
	"fmt"
	"time"
)
import zmq "github.com/pebbe/zmq4"

func old_main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	// socket.Connect("tcp://13.229.105.23:5672")
	// socket.Connect("tcp://127.0.0.1:6000")
	socket.Connect("tcp://192.168.1.141:5672")

	//  msg := "{'action':'publish', 'ranking':[[1,'192.168.1.178'],[2,'192.168.1.206'],[3,'192.168.1.207'],[4,'192.168.1.208'],[5,'192.168.1.209']]}"
	//  msg := "{'action':'attack', 'ranking':[[110, '172.31.15.250'], [105, '172.31.6.243'], [91, '172.31.6.160'], [86, '172.31.15.38'], [61, '172.31.12.132'], [49, '172.31.6.241'], [46, '172.31.29.237']]}"
	//msg := "{'action':'publish', 'ranking':[[12,'192.168.1.141'],[13,'192.168.1.143'],[14,'192.168.1.145'],[14,'192.168.1.147']]}"
	msg:="{'action':'publish', 'ranking':[[1,'192.168.1.141'],[1,'192.168.1.143'],[1,'192.168.1.143']]}"
	_, err := socket.Send(msg, 0)
	if err != nil {
		fmt.Printf("Send err: %v\n", err)
		return
	}
	fmt.Println("Sending", msg)
	socket.Close()

}

func main() {
	context, _ := zmq.NewContext()
	socket, _ := context.NewSocket(zmq.REQ)
	// socket.Connect("tcp://13.229.105.23:5672")
	// socket.Connect("tcp://127.0.0.1:6000")
	socket.Connect("tcp://192.168.1.141:5672")

	//  msg := "{'action':'publish', 'ranking':[[1,'192.168.1.178'],[2,'192.168.1.206'],[3,'192.168.1.207'],[4,'192.168.1.208'],[5,'192.168.1.209']]}"
	//  msg := "{'action':'attack', 'ranking':[[110, '172.31.15.250'], [105, '172.31.6.243'], [91, '172.31.6.160'], [86, '172.31.15.38'], [61, '172.31.12.132'], [49, '172.31.6.241'], [46, '172.31.29.237']]}"
	msg := "{'action':'publish', 'ranking':[[1,'192.168.1.141'],[1,'192.168.1.143'],[1,'192.168.1.145']]}"
	_, err := socket.Send(msg, 0)
	if err != nil {
		fmt.Printf("Send err: %v\n", err)
		return
	}
	fmt.Println("Sending", msg)

	var resp string
	c := make(chan error)
	go func() {
		resp, err = socket.Recv(0)
		c <-err
	}()

	select {
	case err := <- c:
		if err != nil {
			fmt.Printf("Get response err: %v\n", err)
			return
		}
		fmt.Printf("Get response mesg: %s\n", resp)
	case <- time.After(5 * time.Second):
		fmt.Printf("Wait server response timeout\n")
		return
	}

	time.Sleep(10 * time.Second)
}