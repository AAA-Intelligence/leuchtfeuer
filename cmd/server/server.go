package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/AAA-Intelligence/leuchtfeuer/common"
	"google.golang.org/grpc"
)

var clientChannels map[string]chan string

type messengerServer struct {
}

func (s *messengerServer) SetPublicKey(ctx context.Context, pk *common.PublicKey) (*common.Empty, error) {
	fmt.Println("Received public key: ", pk.GetContent())
	return &common.Empty{}, nil
}

func (s *messengerServer) SendMessage(ctx context.Context, m *common.OutgoingMessage) (*common.Empty, error) {
	return &common.Empty{}, nil
}

func (s *messengerServer) ReceiveMessages(e *common.Empty, messageStream common.Messenger_ReceiveMessagesServer) error {
	fmt.Println("Opened message stream")
	//go func() {
	fmt.Println("Sending message...")
	message := common.IncomingMessage{
		Sender:  "bob",
		Content: "Hello world, I am Jeff from the Overwatch team. The rising tide of toxicity...",
	}
	if err := messageStream.Send(&message); err != nil {
		fmt.Println("Err:", err)
	}
	//}()

	return nil
}

func main() {
	fmt.Println("Starting Leuchtfeuer server...")
	fmt.Println("This is an early development build, expect some errors. See our github page to report issues.\n")
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	common.RegisterMessengerServer(grpcServer, &messengerServer{})
	grpcServer.Serve(lis)
}
