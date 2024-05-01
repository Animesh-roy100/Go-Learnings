package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Animesh-roy100/chat-server-grpc/chatserver"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Enter Server IP:Port ::: ")
	reader := bufio.NewReader(os.Stdin)
	serverID, err := reader.ReadString('\n')

	if err != nil {
		log.Printf("Failed to read from console: %v", err)
	}
	serverID = strings.Trim(serverID, "\r\n")

	log.Println("Connecting: " + serverID)

	// connect to grpc server
	conn, err := grpc.Dial(serverID, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gPRC server: %v", err)
	}
	defer conn.Close()

	// call ChatService to create a stream
	client := chatserver.NewServicesClient(conn)

	stream, err := client.ChatService(context.Background())
	if err != nil {
		log.Fatalf("Failed to call ChatService: %v", err)
	}

	// implement communication with gPRC server
	ch := clientHandle{stream: stream}
	ch.clientConfig()
	go ch.sendMessage()
	go ch.receiveMessage()

	// blocker
	bl := make(chan bool)
	<-bl
}

// clientHandle
type clientHandle struct {
	stream     chatserver.Services_ChatServiceClient
	ClientName string
}

func (ch *clientHandle) clientConfig() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Your Name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from console: %v", err)
	}
	ch.ClientName = strings.Trim(name, "\r\n")
}

// send message
func (ch *clientHandle) sendMessage() {
	// create an infinite loop for scanning the console for new message
	for {
		reader := bufio.NewReader(os.Stdin)
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read from console: %v", err)
		}
		clientMessage = strings.Trim(clientMessage, "\r\n")

		clientMessageBox := &chatserver.FromClient{
			Name: ch.ClientName,
			Body: clientMessage,
		}

		err = ch.stream.Send(clientMessageBox)

		if err != nil {
			log.Printf("Error while sending message to server: %v", err)
		}
	}
}

// receive message
func (ch *clientHandle) receiveMessage() {
	for {
		mssg, err := ch.stream.Recv()
		if err != nil {
			log.Printf("Error in receiving message from server: %v", err)
		}

		// print message to console
		fmt.Printf("%s : %s \n", mssg.Name, mssg.Body)
	}
}
