package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kkothule/grpc/chat"
	"google.golang.org/grpc"
)

func main() {

	retryPolicy := `{
		"methodConfig":[{
		  "name":[{"service":"chat.Ping"}],
		  "waitForReady":false,
		  "retryPolicy": {
			  "MaxAttempts":4,
			  "InitialBackoff":"50s",
			  "MaxBackoff":"50s",
			  "BackoffMultiplier":2,
			  "RetryableStatusCodes":["ABORTED"]
		  }
		}]}`

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithDefaultServiceConfig(retryPolicy))
	//	opts = append(opts, grpc.WithDefaultServiceConfig(retryThrottling))
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(":9000", opts...)
	if err != nil {
		log.Fatal("errorr in dial to server:%v", err)
	}
	defer conn.Close()
	c := chat.NewPingClient(conn)
	message := chat.PingMessage{
		Greeting: "test hello from client",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	response, err := c.SayHello(ctx, &message)

	if err != nil {
		log.Fatal("error while sending message to server")
	}
	fmt.Print("Recived response from server : %s", response.Greeting)

	time.Sleep(10 * time.Second)

	fmt.Print("sending second request : ", time.Now())
	response, err = c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatal("error while sending message to server: %v", err)
	}
	fmt.Print("222222222222222Recived response from server :", response.Greeting, "herer:", time.Now())
}
