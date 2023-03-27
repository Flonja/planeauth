package main

import (
	"github.com/flonja/planeauth/auth"
	"github.com/flonja/sinkingchat/chat"
	"log"
	"time"
)

const LinusTechTips = "/live/5c13f3c006f1be15e08e05c0"

func main() {
	token, err := auth.Token()
	if err != nil {
		log.Fatalf("unable to obtain token: %v", err)
	}

	socket, err := chat.NewFloatplaneChatSocket(LinusTechTips, token)
	if err != nil {
		panic(err)
	}
	defer func(socket *chat.FloatplaneChatSocket) {
		err := socket.Close()
		if err != nil {
			panic(err)
		}
	}(socket)

	time.Sleep(time.Second) // timeout, since it would otherwise recognise the next message as spam
	if err = socket.SendMessage("another pointless message from golang! :soontm:"); err != nil {
		panic(err)
	}
}
