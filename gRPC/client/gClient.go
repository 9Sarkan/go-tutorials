package main

import (
	"fmt"

	p "github.com/mactsouk/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = ":8080"

func aboutSayIt(ctx context.Context, m p.MessageServiceClient, text string) (*p.Response, error) {
	request := &p.Request{
		Text:    text,
		Subtext: "new message!",
	}
	r, err := m.SayIt(ctx, request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return r, nil
}
func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("Dial: \t%v\n", err)
		return
	}
	client := p.NewMessageServiceClient(conn)
	r, err := aboutSayIt(context.Background(), client, "Test message!")
	if err != nil {
		fmt.Printf("Say it error: \t%v\n", err)
		return
	}
	fmt.Printf("Response Text: \t%s\nResponse subtext: \t%s\n", r.Text, r.Subtext)
}
