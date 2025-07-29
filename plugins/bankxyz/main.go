// plugins/bankxyz/main.go

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/recitelabs/grpc/proto"
	"google.golang.org/grpc"
)

type BankXYZServer struct {
	pb.UnimplementedBankPluginServer
}

func (s *BankXYZServer) SendPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Simulate calling BankXYZ's actual API here
	// You can add your specific BankXYZ integration logic here
	log.Printf("BankXYZ: Processing payment of %s %s to %s", req.Amount, req.Currency, req.RecipientName)

	return &pb.PaymentResponse{
		Success:       true,
		TransactionId: "XYZ789012",
		Message:       "Payment processed successfully by BankXYZ",
	}, nil
}

func (s *BankXYZServer) CheckStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	// Simulate checking transaction status with BankXYZ
	log.Printf("BankXYZ: Checking status for transaction %s", req.TransactionId)

	// You can add your specific status checking logic here
	return &pb.StatusResponse{Status: "completed"}, nil
}

func main() {
	// BankXYZ plugin runs on a different port to avoid conflicts
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBankPluginServer(s, &BankXYZServer{})

	log.Println("BankXYZ Plugin started on :50052")
	log.Println("Press Ctrl+C to stop the server")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
