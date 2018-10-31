package main

import (
	"context"
	"io"
	"log"

	"github.com/AAA-Intelligence/leuchtfeuer/common"
	"google.golang.org/grpc"
)

func main() {
	// dunno, mach erstmal weg // (y) hab mir so viel mühe gegeben :cry:

	// lass erstmal auskommentiert
	// path := "./log/"
	// dateStr := time.Now().Format("20060102")
	// f, err := os.OpenFile(path+dateStr+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()

	// log.SetOutput(f)

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//TODO error handling
	client := common.NewMessengerClient(conn)
	client.SetPublicKey(context.Background(), &common.PublicKey{Content: "Test"})
	stream, err := client.ReceiveMessages(context.Background(), &common.Empty{})
	if err != nil {
		panic(err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			log.Println("Stream closed")
			break
		}
		if err != nil {
			log.Fatalf("%v.ReceiveMessages(_) = _, %v", client, err)
		}
		log.Println(msg)
	}
}
