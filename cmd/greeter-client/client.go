package main

import (
	"context"
	"fmt"
	greetpb "github.com/rbicker/greeter/api/proto/v1"
	"io"
	"log"
	"time"

	"google.golang.org/grpc/codes"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
	// doServerStreaming(c)
	// doClientStreaming(c)
	// doBidirectionalStreaming(c)
	// doUnaryWithDeadline(c, 5*time.Second)
	// doUnaryWithDeadline(c, 1*time.Second)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an unary RPC")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Raphael",
			LastName:  "Bicker",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)

}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do greeter-server streaming RPC")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Raphael",
			LastName:  "Bicker",
		},
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v\n", msg)
	}
}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to da a Client Streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Raphael",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tanja",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Peter",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Marlon",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pius",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling LongGreet: %v", err)
	}

	for _, req := range requests {
		stream.Send(req)
		fmt.Printf("Sending req: %v\n", req)
		time.Sleep(1000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while closing stream: %v", err)
	}
	fmt.Printf("LongGreet Response: %v\n", res)
}

func doBidirectionalStreaming(c greetpb.GreetServiceClient) {
	requests := []*greetpb.GreetEveryoneRequest{
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Raphael",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Tanja",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Peter",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Marlon",
			},
		},
		&greetpb.GreetEveryoneRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Pius",
			},
		},
	}

	//  we create a stream by invoking greeter-client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}
	waitc := make(chan struct{})
	// we send a bunch of messages to the greeter-client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			stream.Send(req)
			fmt.Printf("Sending request: %v\n", req)
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the greeter-client (go routine)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("I guess I am finished...")
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error while receiving stream: %v", err)
			}
			fmt.Printf("Received response: %v\n", res)
		}
	}()
	// block until everything is done
	<-waitc
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, timeout time.Duration) {
	fmt.Println("Starting to do an unary RPC with deadline")
	req := &greetpb.GreetWithDeadlineRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Raphael",
			LastName:  "Bicker",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit! Deadline was exceeded")
			} else {
				fmt.Printf("Unexpected grpc error: %v \n", statusErr)
			}
			return
		} else {
			log.Fatalf("error while calling GreetWithDeadline RPC: %v", err)
		}
	}
	log.Printf("Response from GreetWithDeadline: %v", res.Result)

}
