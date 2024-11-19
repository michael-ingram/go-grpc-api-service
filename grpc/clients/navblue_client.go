package grpc

import (
	"context"
	"go-api-service/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	NavblueServiceAddress = "localhost:50051" // Update with actual service address
)

// Navblue Client wraps the gRPC client connection
type NavblueClient struct {
	Client pb.NavblueServiceClient
	conn   *grpc.ClientConn
}

// New Navblue Client creates a new instance of NavblueClient
func NewNavblueClient(address string) *NavblueClient {
	// Create a connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect to gRPC server at %s: %v", address, err)
	}

	// Create a new NavblueService client
	client := pb.NewNavblueServiceClient(conn)

	return &NavblueClient{
		Client: client,
		conn:   conn,
	}
}

// Close closes the gRPC connection
func (c *NavblueClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetFlightById fetches flight details by its unique ID
func (c *NavblueClient) GetFlightById(ctx context.Context, uniqueID string) (*pb.FlightByIdResponse, error) {
	// Prepare the request
	req := &pb.FlightByIdRequest{
		UniqueId: uniqueID,
	}

	// Set a timeout for the request
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Make the gRPC call
	resp, err := c.Client.FlightById(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
