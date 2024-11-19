package graph

import (
	"go-api-service/graph/model"
	grpc "go-api-service/grpc/clients"
	"go-api-service/grpc/pb"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	NavblueClient *grpc.NavblueClient
}

func mapDelays(delays []*pb.Delay) []*model.Delay {
	var mappedDelays []*model.Delay
	for _, d := range delays {
		mappedDelays = append(mappedDelays, &model.Delay{
			Code:    d.Code,
			Code2:   d.Code2,
			Minutes: d.Minutes,
		})
	}
	return mappedDelays
}
func mapFuels(fuels []*pb.Fuel) []*model.Fuel {
	var mappedFuels []*model.Fuel
	for _, f := range fuels {
		mappedFuels = append(mappedFuels, &model.Fuel{
			Type:         f.Type,
			Source:       f.Source,
			Quantity:     f.Quantity,
			QuantityUnit: f.QuantityUnit,
			Density:      f.Density,
			DensityUnit:  f.DensityUnit,
		})
	}
	return mappedFuels
}
