package main

import (
	"fmt"
	"time"

	zmq "github.com/pebbe/zmq4"
)

const (
	internalCommunication = "inproc://internal-communication"
	timeout               = 30 * time.Second
)

func main() {
	sender, err := newSenderSocket()
	if nil != err {
		fmt.Println("new source socket with error: ", err)
		return
	}
	defer sender.Close()

	receiver, err := newReceiveSocket()
	if nil != err {
		fmt.Println("new receiver socket with error: ", err)
		return
	}
	defer receiver.Close()

	go senderLoop(sender)
	go receiverLoop(receiver)

	<-time.After(timeout)
}

func newSenderSocket() (*zmq.Socket, error) {
	src, err := zmq.NewSocket(zmq.PAIR)
	if nil != err {
		return nil, err
	}

	src.Bind(internalCommunication)

	return src, nil
}

func newReceiveSocket() (*zmq.Socket, error) {
	receiver, err := zmq.NewSocket(zmq.PAIR)
	if nil != err {
		return nil, err
	}

	err = receiver.Connect(internalCommunication)
	if nil != err {
		return nil, err
	}

	return receiver, nil
}

func senderLoop(soc *zmq.Socket) {
	interval := 5 * time.Second
	timer := time.NewTimer(interval)

	for i := 0; i < 5; i++ {
		select {
		case <-timer.C:
			msg := fmt.Sprintf("%d message", i)
			_, err := soc.Send(msg, 0)
			if nil != err {
				fmt.Println("send message with error: ", err)
				continue
			}
			timer.Reset(interval)
		}
	}
}

func receiverLoop(soc *zmq.Socket) {
	for {
		msg, err := soc.Recv(0)
		if nil != err {
			fmt.Println("receive message with error: ", err)
			continue
		}
		fmt.Println("received msg: ", msg)
	}
}
