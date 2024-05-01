package chatserver

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// Define a messageUnit struct for handling message in server
type messageUnit struct {
	ClientName        string
	MessageBody       string
	MessageUniqueCode int
	ClientUniqueCode  int
}

// messageHandle struct will hold slice of messageUnits
type messageHandle struct {
	MQue []messageUnit
	mu   sync.Mutex // Add a mutex variable to handle asynchronous read write operation
}

var messageHandleObject = messageHandle{}

// Create a structure ChatServer that will implement ChatService method
type ChatServer struct {
}

func (cs *ChatServer) mustEmbedUnimplementedServicesServer() {}

// Define the ChatService method which takes an arg of type Service_ChatServiceServer
func (is *ChatServer) ChatService(csi Services_ChatServiceServer) error {
	clientUniqueCode := rand.Intn(1e6)
	errch := make(chan error)

	// receive messages - init a go routine
	go receiveFromStream(csi, clientUniqueCode, errch)

	// send messages - init a go routine
	go sendToStream(csi, clientUniqueCode, errch)

	return <-errch
}

// receive messages
func receiveFromStream(csi_ Services_ChatServiceServer, clientUniqueCode_ int, errch_ chan error) {
	// implement a loop
	for {
		mssg, err := csi_.Recv()
		if err != nil {
			log.Printf("Error in receiving message from client: %v", err)
			errch_ <- err
		} else {
			// after receiving a message from client, the same is added to message-Que(MQue)
			messageHandleObject.mu.Lock()

			messageHandleObject.MQue = append(messageHandleObject.MQue, messageUnit{
				ClientName:        mssg.Name,
				MessageBody:       mssg.Body,
				MessageUniqueCode: rand.Intn(1e8),
				ClientUniqueCode:  clientUniqueCode_,
			})

			messageHandleObject.mu.Unlock()

			log.Printf("%v", messageHandleObject.MQue[len(messageHandleObject.MQue)-1])
		}
	}
}

// send message
func sendToStream(csi_ Services_ChatServiceServer, clientUniqueCode_ int, errch_ chan error) {
	// implement a loop
	for {
		// loop through messages in MQue
		for {
			time.Sleep(500 * time.Microsecond)

			messageHandleObject.mu.Lock()

			if len(messageHandleObject.MQue) == 0 {
				messageHandleObject.mu.Unlock()
				break
			}

			senderUniqueCode := messageHandleObject.MQue[0].ClientUniqueCode
			senderNameForClient := messageHandleObject.MQue[0].ClientName
			messageForClient := messageHandleObject.MQue[0].MessageBody

			messageHandleObject.mu.Unlock()

			// send message to designated client
			if senderUniqueCode != clientUniqueCode_ {
				err := csi_.Send(&FromServer{Name: senderNameForClient, Body: messageForClient})

				if err != nil {
					errch_ <- err
				}

				messageHandleObject.mu.Lock()

				// Delete message from MQue once sent to the designated client
				if len(messageHandleObject.MQue) > 1 {
					messageHandleObject.MQue = messageHandleObject.MQue[1:] // delete the message at index 0 after sending to the receiver
				} else {
					messageHandleObject.MQue = []messageUnit{}
				}

				messageHandleObject.mu.Unlock()
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
}
