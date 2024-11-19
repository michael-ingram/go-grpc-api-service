package grpc

import (
	"context"
	"go-api-service/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type NavitaireClient struct {
	Client pb.NavitaireServiceClient
	conn   *grpc.ClientConn
}

// New Navitaire Client initializes a new gRPC client for the Navitaire service.
func NewNavitaireClient(address string) *NavitaireClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to Navitaire gRpc Server: %v", err)
	}

	client := pb.NewNavitaireServiceClient(conn)
	return &NavitaireClient{
		Client: client,
		conn:   conn,
	}
}

func (c *NavitaireClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetStations retrieves a list of stations from the Navitaire service.
func (c *NavitaireClient) GetStations(ctx context.Context) (*pb.GetStationsResponse, error) {
	// Send an empty request as per the proto definition
	return c.Client.GetStations(ctx, &emptypb.Empty{})
}
