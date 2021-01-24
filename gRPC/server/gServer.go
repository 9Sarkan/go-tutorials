package main

import (
	"fmt"
	"net"

	p "github.com/mactsouk/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// MessageService struct
type MessageService struct{}

var port = ":8080"

// SayIt grpc Function
func (MessageService) SayIt(ctx context.Context, r *p.Request) (*p.Response, error) {
	fmt.Println("Request Text: ", r.Text)
	fmt.Println("Request Subtext: ", r.Subtext)

	response := &p.Response{
		Text:    r.Text,
		Subtext: "Got it",
	}
	return response, nil
}
func main() {
	server := grpc.NewServer()
	var messageService MessageService
	p.RegisterMessageServiceServer(server, messageService)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Serving...")
	server.Serve(listen)
}
