// main.go

package main

import (
	"context"
	"log"
	"time"

	pb "github.com/recitelabs/grpc/github.com/recitelabs/grpc/plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// You can change the bank by modifying this address
	// Available banks:
	// - BankABC: localhost:50051
	// - BankXYZ: localhost:50052
	// - Template: localhost:50053
	bankAddress := "localhost:50051" // Change this to connect to different banks

	conn, err := grpc.Dial(bankAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBankPluginClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.SendPayment(ctx, &pb.PaymentRequest{
		Amount:        "1000",
		Currency:      "ETB",
		AccountNumber: "123456789",
		RecipientName: "Bereket Kume",
	})
	if err != nil {
		log.Fatalf("SendPayment error: %v", err)
	}

	log.Printf("Payment Result: %+v", resp)
}
