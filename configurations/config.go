package configurations

import (
	"os"
)

type Config struct {
	GrpcNavblueAddr string
}

func LoadConfiguration() *Config {
	// For example, loading from environment variables or default values
	grpcNavblueAddr := os.Getenv("GRPC_NAVBLUE_ADDR")
	if grpcNavblueAddr == "" {
		grpcNavblueAddr = "127.0.0.1:50051" // Default address
	}

	//grpcManifestAddr := os.Getenv("GRPC_MANIFEST_ADDR")
	//if grpcManifestAddr == "" {
	//	grpcManifestAddr = "localhost:50052" // Default address
	//}

	return &Config{
		GrpcNavblueAddr: grpcNavblueAddr,
	}
}
