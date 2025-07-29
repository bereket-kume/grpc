// plugins/template/main.go - Template for creating new bank plugins

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/recitelabs/grpc/proto"
	"google.golang.org/grpc"
)

// BankTemplateServer implements the BankPlugin service
type BankTemplateServer struct {
	pb.UnimplementedBankPluginServer
}

// SendPayment handles payment requests
func (s *BankTemplateServer) SendPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// TODO: Add your bank-specific payment processing logic here
	// This is where you would integrate with your bank's actual API

	log.Printf("BankTemplate: Processing payment of %s %s to %s", req.Amount, req.Currency, req.RecipientName)

	// Example: Add your bank's API call here
	// response, err := yourBankAPI.ProcessPayment(req.Amount, req.Currency, req.AccountNumber, req.RecipientName)

	return &pb.PaymentResponse{
		Success:       true,
		TransactionId: "TEMPLATE123",
		Message:       "Payment processed successfully by BankTemplate",
	}, nil
}

// CheckStatus handles status check requests
func (s *BankTemplateServer) CheckStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	// TODO: Add your bank-specific status checking logic here
	// This is where you would query your bank's API for transaction status

	log.Printf("BankTemplate: Checking status for transaction %s", req.TransactionId)

	// Example: Add your bank's status API call here
	// status, err := yourBankAPI.GetTransactionStatus(req.TransactionId)

	return &pb.StatusResponse{Status: "completed"}, nil
}

func main() {
	// TODO: Change the port number to avoid conflicts with other plugins
	port := ":50053" // Change this to a unique port

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBankPluginServer(s, &BankTemplateServer{})

	log.Printf("BankTemplate Plugin started on %s", port)
	log.Println("Press Ctrl+C to stop the server")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
