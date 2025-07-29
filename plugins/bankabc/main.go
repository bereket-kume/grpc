// plugins/bankabc/main.go

package main

import (
	"context"
	"log"
	"net"

	pb "github.com/recitelabs/grpc/github.com/recitelabs/grpc/plugin"
	"google.golang.org/grpc"
)

type BankABCServer struct {
	pb.UnimplementedBankPluginServer
}

func (s *BankABCServer) SendPayment(ctx context.Context, req *pb.PaymentRequest) (*pb.PaymentResponse, error) {
	// Simulate calling BankABC's actual API here
	return &pb.PaymentResponse{
		Success:       true,
		TransactionId: "ABC123456",
		Message:       "Payment processed by BankABC",
	}, nil
}

func (s *BankABCServer) CheckStatus(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	// Simulate checking transaction status
	return &pb.StatusResponse{Status: "completed"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051") // or use Unix socket
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterBankPluginServer(s, &BankABCServer{})
	log.Println("BankABC Plugin started on :50051")
	s.Serve(lis)
}
